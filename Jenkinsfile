
node {
    def root = tool name: 'Go', type: 'go'
    env.GOROOT = root
    env.PATH = "${root}/bin:${env.PATH}"

    stage('Checkout') {
        git url: 'https://github.com/ZefanyaPrasetiyo/be-encore.git'
    }

    stage('PreTest') {
        sh 'go version'
        sh 'go env GOMOD'
        sh 'go vet ./...'
    }

    stage('Test') {
        sh 'go test -cover ./...'
    }

    stage('Build') {
        sh 'go build ./...'
    }

    stage('Deploy') {
        echo 'Deployment belum dikonfigurasi'
    }
}
