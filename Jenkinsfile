pipeline {
    agent {
        docker { image 'golang' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'go build main.go'
            }
        }
    }
}
