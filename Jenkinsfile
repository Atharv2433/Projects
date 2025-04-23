pipeline {
    agent any

    environment {
        GO_VERSION = '1.19' // Specify your Go version
        GO_PATH = '/usr/local/go'
        GOPATH = "${env.WORKSPACE}/go"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install Go') {
            steps {
                script {
                    if (!isUnix()) {
                        bat "choco install golang --version=${GO_VERSION}"
                    } else {
                        sh "curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz"
                        sh "sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz"
                    }
                }
            }
        }

        stage('Set Go Environment') {
            steps {
                script {
                    if (isUnix()) {
                        sh """
                            export PATH=\$PATH:/usr/local/go/bin
                            export GOPATH=\$WORKSPACE/go
                            export GOROOT=/usr/local/go
                        """
                    } else {
                        bat """
                            set PATH=%PATH%;C:\\Go\\bin
                            set GOPATH=%WORKSPACE%\\go
                            set GOROOT=C:\\Go
                        """
                    }
                }
            }
        }

        stage('Install Dependencies') {
            steps {
                script {
                    sh 'go mod tidy'
                }
            }
        }

        stage('Run Migrations') {
            steps {
                script {
                    // Replace with your migration command
                    sh 'go run cmd/migrate/main.go'
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    sh 'go test ./...'
                }
            }
        }

        stage('Run Application') {
            steps {
                script {
                    sh 'go run main.go'
                }
            }
        }
    }

    post {
        always {
            echo 'Cleaning up...'
        }

        success {
            echo 'Build and tests passed!'
        }

        failure {
            echo 'Build or tests failed.'
        }
    }
}
