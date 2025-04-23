pipeline {
    agent any

    environment {
        GO_VERSION = '1.19'
        GOPATH = "${env.WORKSPACE}\\go"
        GOROOT = "C:\\Program Files\\Go"
        PATH = "${env.PATH};${GOROOT}"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        // stage('Install Go') {
        //     steps {
        //         bat '''
        //             @echo off
        //             where choco >nul 2>&1
        //             if %ERRORLEVEL% NEQ 0 (
        //                 powershell -NoProfile -ExecutionPolicy Bypass -Command "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
        //             )
        //             choco install golang --version=1.19 -y
        //         '''
        //     }
        // }

        stage('Install Dependencies') {
            steps {
                bat """
                    set "GOROOT=C:\\Program Files\\Go"
                    set "GOPATH=%WORKSPACE%\\go"
                    set "PATH=%GOROOT%\\bin;%PATH%"\
                    go version
                    go mod tidy
                """
            }
        }

        stage('Run Migrations') {
            steps {
                bat """
                    set PATH=%PATH%;C:\\Go\\bin
                    set GOPATH=%WORKSPACE%\\go
                    set GOROOT=C:\\Go
                    go run cmd\\migrate\\main.go
                """
            }
        }

        stage('Run Tests') {
            steps {
                bat """
                    set PATH=%PATH%;C:\\Go\\bin
                    set GOPATH=%WORKSPACE%\\go
                    set GOROOT=C:\\Go
                    go test ./...
                """
            }
        }

        stage('Run Application') {
            steps {
                bat """
                    set PATH=%PATH%;C:\\Go\\bin
                    set GOPATH=%WORKSPACE%\\go
                    set GOROOT=C:\\Go
                    go run main.go
                """
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