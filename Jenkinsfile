pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }

    environment {
        GO_BIN = "/usr/local/go/bin/go"
        IMAGE_NAME = "johannes/book-management"
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

        // =========================
        // CI SECTION (DOCKER BUILD)
        // =========================
        stage('Docker Build (CI)') {
            steps {
                sh 'docker build -t ${IMAGE_NAME}:latest .'
            }
        }

        stage('Push Image (CI)') {
            when {
                branch 'main'
            }
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    docker tag ${IMAGE_NAME}:latest $DOCKER_USER/book-management:latest
                    docker push $DOCKER_USER/book-management:latest
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

        // =========================
        // CD SECTION (DEPLOY)
        // =========================
        stage('Deploy (CD)') {
            when {
                branch 'main'
            }
            steps {
                sh '''
                docker pull ${IMAGE_NAME}:latest
                docker stop book-management || true
                docker rm book-management || true
                docker run -d -p 8080:8080 --name book-management ${IMAGE_NAME}:latest
                '''
            }
        }

        stage('Debug') {
            steps {
                sh 'docker --version'
                sh 'echo $PATH'
                sh 'which go'
                sh 'which golangci-lint'
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