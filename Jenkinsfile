pipeline {
    agent any

    environment {
        DOCKERHUB_REPO = "peenesss/book-management"
        VPS_IP = "103.149.177.39"
        CONTAINER_NAME = "book-management"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                credentialsId: 'github-johannes',
                url: 'git@github.com:johannessipayung/Book-Management-Golang-challenge-internship-DOT-.git'
            }
        }

        stage('Quality Checks') {
            parallel {

                stage('Go Vet') {
                    steps {
                        sh 'go vet ./...'
                    }
                }

                stage('Lint') {
                    steps {
                        sh 'golangci-lint run'
                    }
                }

                stage('Test') {
                    steps {
                        sh 'go test ./... -coverprofile=coverage.out'
                    }
                }
            }
        }

        stage('Build Binary') {
            steps {
                sh '''
                go mod download
                CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book-management .
                '''
            }
        }

        stage('Docker Build') {
            steps {
                sh '''
                docker buildx build \
                --platform linux/amd64 \
                -t ${DOCKERHUB_REPO}:latest \
                .
                '''
            }
        }

        stage('Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub',
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

        stage('Deploy to VPS') {
            steps {
                sshagent(['root']) {
                    sh '''
                    ssh -o StrictHostKeyChecking=no root@${VPS_IP} '

                        echo "Pull latest image..."
                        docker pull ${DOCKERHUB_REPO}:latest

                        echo "Stop old container..."
                        docker stop ${CONTAINER_NAME} || true
                        docker rm ${CONTAINER_NAME} || true

                        echo "Run new container..."
                        docker run -d \
                        --name ${CONTAINER_NAME} \
                        -p 9090:8080 \
                        --restart unless-stopped \
                        -e DB_HOST=172.17.0.1 \
                        -e DB_USER=johannes \
                        -e DB_PASSWORD=mypassword123 \
                        -e DB_NAME=challengego \
                        -e DB_PORT=5432 \
                        -e DB_SSLMODE=disable \
                        ${DOCKERHUB_REPO}:latest

                        echo "Running containers:"
                        docker ps
                    '
                    '''
                }
            }
        }

        stage('Health Check') {
            steps {
                sh '''
                sleep 10
                curl -f http://103.149.177.39:9090 || exit 1
                '''
            }
        }

    }

    post {

        success {
            echo "Deployment SUCCESS"
        }

        failure {
            echo "Pipeline FAILED"
        }

        always {
            archiveArtifacts artifacts: 'coverage.out', fingerprint: true
        }
    }
}