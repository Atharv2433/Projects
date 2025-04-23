pipeline {
    agent any

    environment {
        GO_VERSION = '1.19'
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
                    if (isUnix()) {
                        sh """
                            curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                            sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                        """
                    } else {
                        bat '''
                            @echo off
                            where choco >nul 2>&1
                            if %ERRORLEVEL% NEQ 0 (
                                powershell -NoProfile -ExecutionPolicy Bypass -Command "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
                            )
                            choco install golang --version=%GO_VERSION% -y
                        '''
                    }
                }
            }
        }

        stage('Install Dependencies') {
            steps {
                script {
                    if (isUnix()) {
                        sh """
                            export PATH=\$PATH:/usr/local/go/bin
                            export GOPATH=\$WORKSPACE/go
                            export GOROOT=/usr/local/go
                            go mod tidy
                        """
                    } else {
                        bat """
                            set PATH=%PATH%;C:\\Go\\bin
                            set GOPATH=%WORKSPACE%\\go
                            set GOROOT=C:\\Go
                            go mod tidy
                        """
                    }
                }
            }
        }

        stage('Run Migrations') {
            steps {
                script {
                    if (isUnix()) {
                        sh """
                            export PATH=\$PATH:/usr/local/go/bin
                            export GOPATH=\$WORKSPACE/go
                            export GOROOT=/usr/local/go
                            go run cmd/migrate/main.go
                        """
                    } else {
                        bat """
                            set PATH=%PATH%;C:\\Go\\bin
                            set GOPATH=%WORKSPACE%\\go
                            set GOROOT=C:\\Go
                            go run cmd\\migrate\\main.go
                        """
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    if (isUnix()) {
                        sh """
                            export PATH=\$PATH:/usr/local/go/bin
                            export GOPATH=\$WORKSPACE/go
                            export GOROOT=/usr/local/go
                            go test ./...
                        """
                    } else {
                        bat """
                            set PATH=%PATH%;C:\\Go\\bin
                            set GOPATH=%WORKSPACE%\\go
                            set GOROOT=C:\\Go
                            go test ./...
                        """
                    }
                }
            }
        }

        stage('Run Application') {
            steps {
                script {
                    if (isUnix()) {
                        sh """
                            export PATH=\$PATH:/usr/local/go/bin
                            export GOPATH=\$WORKSPACE/go
                            export GOROOT=/usr/local/go
                            go run main.go
                        """
                    } else {
                        bat """
                            set PATH=%PATH%;C:\\Go\\bin
                            set GOPATH=%WORKSPACE%\\go
                            set GOROOT=C:\\Go
                            go run main.go
                        """
                    }
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