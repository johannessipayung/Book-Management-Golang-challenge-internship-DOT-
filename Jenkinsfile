pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }

    stages {

        stage('Build') {
            steps {
                sh '/usr/local/go/bin/go version'
                sh '/usr/local/go/bin/go mod download'
                sh '/usr/local/go/bin/go build -o book-management .'
                sh 'ls -la'
            }
        }
        stage('Run (Short)') {
            steps {
                sh 'timeout 5 ./book-management || true'
            }
        }

        stage('Test') {
            steps {
                sh '/usr/local/go/bin/go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                echo "Simulating deploy..."
            }
        }
    }

    post {
        always {
            echo "I will always say Hello again!"
        }
        success {
            echo "Yay, success"
        }
        failure {
            echo "Oh no, failure"
        }
        cleanup {
            echo "Don't care success or error"
        }
    }
}