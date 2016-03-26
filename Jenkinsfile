node('executor') {
    git 'git@github.com:conjurinc/summon.git'

    stage 'Test - Unit'
    sh './test.sh'
    step([$class: 'JUnitResultArchiver', testResults: 'junit.xml'])

    stage 'Build'
    sh '''
      ./build.sh
    '''

    stage 'Test - Acceptance'
    sh '''
      cp pkg/linux-amd64/summon .
      cd acceptance && make
    '''

    stage 'Package'
    sh './package.sh'
    archive 'pkg/**/*'
}

