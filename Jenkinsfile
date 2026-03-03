pipeline {
    agent {
        node {
            label 'golang && linux'
        }
    }
    
    stages {
        stage('Build') {
            steps {
                echo ("Hello Build 1")
                sleep(5)
                echo ("Hello Build 2")
            }
        }

        stage('Test') {
            steps {
                echo ("Hello Test 1")
                sleep(5)
                echo ("Hello Test 2")
            }
        }

        stage('Deploy') {
            steps {
                echo ("Hello Deploy 1")
                sleep(5)
                echo ("Hello Deploy 2")
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