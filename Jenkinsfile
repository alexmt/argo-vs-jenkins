podTemplate(label: 'ci-workflow', containers: [
    containerTemplate(name: 'golang', image: 'golang', ttyEnabled: true, command: 'cat')
]) {
    node('kube') {
        stage('Build') {
            container('golang') {
                git 'https://github.com/alexmt/argo-vs-jenkins.git'
                sh 'go build main.go'
            }
        }

        stage('Release') {
            parallel {
                stage('Test 1') {
                    container('golang') {
                        sh "go test -v"
                    }
                }
                stage('Test 2') {
                    container('golang') {
                        sh "go test -v"
                    }
                }
                stage('Test 3') {
                    container('golang') {
                        sh "go test -v"
                    }
                }
            }
        }

        stage('Release') {
            container('golang') {
                sh './main'
            }
        }
    }
}
