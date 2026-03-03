pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }

    stages {
        stage('Build') {
            steps {
                sh 'go mod download'
                sh 'go build -o book-management .'
            }
        }

        stage('Run') {
            steps {
                sh './book-management'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
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