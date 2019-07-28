#!/bin/bash

cd `dirname $0`

rm -rf ./dist

mkdir ./dist
mkdir ./dist/static
mkdir ./dist/templates
mkdir ./dist/testdatas

cp -f ./setting.yml ./dist/

dep ensure
if [ $? -ne 0 ]; then
  echo "`dep ensure` filed"
  exit 1
fi

go build -o dist/server main.go
if [ $? -ne 0 ]; then
  echo "`go build -o dist/server main.go` filed"
  exit 2
fi

echo server/build.sh is successful

