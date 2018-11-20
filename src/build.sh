#!/bin/bash

case "$(uname)" in
*MINGW* | *WIN32* | *CYGWIN*)
       echo 'ERROR: Do not use make.bash to build on Windows.'
       echo 'Use make.bat instead.'
       echo
       exit 1
       ;;
Darwin)
    echo "begin build in Darwin"
    GOOS=darwin
    GOARCH=amd64
    ;;
*)
    echo "begin build in linux"
    GOOS=linux
    GOARCH=amd64
    ;;
esac

go build -gcflags '-N -l'

