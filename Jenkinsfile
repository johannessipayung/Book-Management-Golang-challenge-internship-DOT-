pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }

    environment {
        GO_BIN = "/usr/local/go/bin/go"
        PATH = "/opt/homebrew/bin:${env.PATH}"
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
                        sh 'golangci-lint run || true'
                    }
                }

                stage('Test') {
                    steps {
                        sh '${GO_BIN} test -coverprofile=coverage.out ./...'
                    }
                }
            }
        }

        stage('Build') {
            steps {
                sh '${GO_BIN} mod download'
                sh '${GO_BIN} build -o book-management .'
            }
        }

        stage('Generate Coverage Report') {
            steps {
                sh '${GO_BIN} tool cover -html=coverage.out -o coverage.html'
            }
        }

        stage('Archive Artifacts') {
            steps {
                archiveArtifacts artifacts: 'book-management, coverage.html', fingerprint: true
            }
        }

        stage('Deploy') {
            when {
                branch 'main'
            }
            steps {
                echo "Simulating deploy to environment..."
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
