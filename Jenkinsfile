pipeline {
    agent any
      environment {
        GO_VERSION = '1.22.0'  // Specify the Go version
        GO_HOME = "/usr/local/go"
        PATH = "${GO_HOME}/bin:${env.PATH}"
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOPATH = "${env.JENKINS_HOME}/jobs/${env.JOB_NAME}/builds/${env.BUILD_ID}"
    }
    stages {
        stage (" Clean up "){
        steps {
            deleteDir()
        }
    }
    stage('Checkout') {
            steps {
                // Checkout code from the repository
                git branch: 'main', url: 'https://github.com/rahul-chaube/NTTRepository.git'
            }
    }
    // stage (" clone repository "){
    //     steps {
    //         sh "git clone https://github.com/rahul-chaube/DockerStudy.git"
    //     }
    // }
    stage (" Check Version "){
        steps {
            sh "go version"
        }
    }
    stage (" Build "){
        steps{
                sh("pwd")
                sh ("ls")
                sh ("go build main.go")
            
        }
    }
    }
     
}