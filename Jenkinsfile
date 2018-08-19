// Run on an agent where we want to use Go
node {
    // Install the desired Go version
    def root = tool name: 'Go 1.10.3', type: 'go'

    // Export environment variables pointing to the directory where Go was installed
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"
            env.CHECKOUTPATH="$GOPATH/src/github.com/ezzizzle/uaparser"

            dir("$CHECKOUTPATH") {
                stage('Checkout') {
                    checkout scm
                    // git(url: 'https://github.com/ezzizzle/uaparser.git')
                }

                stage('Pre Test') {
                    echo 'Pulling Dependencies'

                    sh 'go version'
                    sh 'go get -u github.com/jteeuwen/go-bindata/...'
                    sh 'cd useragent && go-bindata -pkg useragent reference/'
                }

                stage('Test') {
                    sh 'pwd'
                    sh 'ls'
                    sh 'go test ./...'
                }

                stage('Build') {
                    dir('uaparserserver') {
                        parallel mac: {
                            echo 'Building for Mac'
                            sh 'GOOS=darwin GOARCH=amd64 go build -o build/uaparserserver.mac'
                            archiveArtifacts artifacts: 'build/uaparserserver.*', fingerprint: true
                        },
                        linux: {
                            echo 'Building for Linux'
                            sh 'GOOS=linux GOARCH=amd64 go build -o build/uaparserserver.linux'
                            archiveArtifacts artifacts: 'build/uaparserserver.*', fingerprint: true
                        }
                    }
                }
            }
        }
    }
}