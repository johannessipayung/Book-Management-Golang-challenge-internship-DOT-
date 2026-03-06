pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }

    environment {
        GO_BIN = "/usr/local/go/bin/go"
        DOCKERHUB_REPO = "peenesss/book-management"
        CONTAINER_NAME = "book-management"
        APP_PORT = "9090"
        PATH = "/opt/homebrew/bin:/usr/local/bin:/usr/local/go/bin:${env.PATH}"
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Quality Checks') {
            parallel {

                stage('Vet') {
                    steps {
                        sh '${GO_BIN} vet ./...'
                    }
                }

                stage('Lint') {
                    steps {
                        sh '''
                        export PATH=$PATH:/usr/local/go/bin:/opt/homebrew/bin
                        golangci-lint run
                        '''
                    }
                }

                stage('Test') {
                    steps {
                        sh '${GO_BIN} test -coverprofile=coverage.out ./...'
                    }
                }
            }
        }

        stage('Build Binary') {
            steps {
                sh '${GO_BIN} mod download'
                sh '${GO_BIN} build -o book-management .'
            }
        }

        stage('Docker Build (CI)') {
            steps {
                sh 'docker build -t ${DOCKERHUB_REPO}:latest .'
            }
        }

        stage('Push Image (CI)') {
            when {
                expression { env.BRANCH_NAME == 'main' || env.GIT_BRANCH == 'origin/main' }
            }
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    docker push ${DOCKERHUB_REPO}:latest
                    '''
                }
            }
        }

        stage('Generate Coverage Report') {
            steps {
                sh '${GO_BIN} tool cover -html=coverage.out -o coverage.html'
            }
        }

        stage('Archive Artifacts') {
            steps {
                archiveArtifacts artifacts: 'coverage.html', fingerprint: true
            }
        }

        stage('Deploy (CD)') {
            steps {
                sshagent(['vps-ssh']) {
                    sh '''
                    echo "Connecting to VPS and deploying..."

                    ssh -o StrictHostKeyChecking=no root@103.149.177.39 << 'ENDSSH'

                    echo "Stopping old container if exists..."
                    docker stop book-management || true
                    docker rm book-management || true

                    echo "Pulling latest image from Docker Hub..."
                    docker pull peenesss/book-management:latest

                    echo "Running new container on port 9090..."
                    docker run -d \
                    --name book-management \
                    -p 9090:8080 \
                    --restart unless-stopped \
                    peenesss/book-management:latest

                    ENDSSH
                    '''
                }
            }
        }   

        stage('Debug') {
            steps {
                sh 'docker --version'
                sh 'echo $PATH'
                sh 'which go'
                sh 'which golangci-lint'
                sh 'docker ps -a'
            }
        }
    }

    post {
        always {
            echo "Pipeline finished."
        }
        success {
            echo "Build SUCCESS"
        }
        failure {
            echo "Build FAILED"
        }
    }
}