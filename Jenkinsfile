pipeline {
    agent any
    triggers{
     githubPullRequest {
             
            cron('* * * * *') // Check for new PRs every minute
            allowMembersOfWhitelistedOrgsAsAdmin(true) // Optional: allow members of whitelisted orgs as admins
        }
    }
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
                git branch: 'main', url: 'https://github.com/rahul-chaube/NTTRepository.git'
            }
    }
    stage (" Check Version "){
        steps {
            sh "go version"
        }
    }
    stage (" Build "){
        steps{
                sh ("go build main.go")
            
        }
    }

    stage ("Final stage"){
        steps{
                echo ("Completed")
            
        }
    }
    }
     
}
