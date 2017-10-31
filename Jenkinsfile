pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('Build') {
            steps {
                sh 'find . && go build main.go'
            }
        }

        stage('Test') {
            parallel {
                stage('Test 1') {
                    steps {
                        sh "go test -v"
                    }
                }
                stage('Test 2') {
                    steps {
                        sh "go test -v"
                    }
                }
                stage('Test 3') {
                    steps {
                        sh "go test -v"
                    }
                }
            }
        }

        stage('Release') {
            steps {
                sh './main'
            }
        }
    }
}
