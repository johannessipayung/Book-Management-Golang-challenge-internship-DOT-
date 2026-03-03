pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }
    
    environment {
        SCRIPT = "book-management"
    }

    stages {
        stage('Build') {
            steps {
                sh '/usr/local/go/bin/go mod download'
                sh '/usr/local/go/bin/go build -o book-management .'
            }
        }

        stage('Run') {
            steps {
                sh '/usr/local/go/bin/${SCRIPT}'
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