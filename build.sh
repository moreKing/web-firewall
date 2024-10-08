#!/bin/bash

current_dir=$(pwd)

echo $current_dir

rm -rf ~/firewalld-web
mkdir -p ~/firewalld-web

cd $current_dir/web
for item in $(ls -a); do
    if [ "$item" == "node_modules" ]; then
        echo 'skip node_modules'
    elif [ "$item" == "." ]; then
    
        echo 'skip .'
    elif [ "$item" == ".." ]; then
        echo 'skip ..'
    else
        cp -rf "$item" ~/firewalld-web/
        echo "$item"
    fi
done

cd ~/firewalld-web
pnpm install
pnpm run build

echo 'web build success'
cd $current_dir

cd $current_dir/server
go mod tidy
echo 'server start  build'
CGO_ENABLED=0 go build -o server
echo 'server build success'

cd $current_dir
rm -rf build
mkdir -p build/web-firewalld

cp $current_dir/script/* $current_dir/build/
cp -r $current_dir/server/manifest $current_dir/build/web-firewalld/
cp -r $current_dir/server/resource $current_dir/build/web-firewalld/
cp -r $current_dir/server/server $current_dir/build/web-firewalld/

rm -rf $current_dir/build/web-firewalld/resource/public/html
cp -r ~/firewalld-web/dist $current_dir/build/web-firewalld/resource/public/html

rm -rf $current_dir/build/web-firewalld/resource/*.sqlite3
mv $current_dir/build/web-firewalld/resource/db.sqlite3.bak $current_dir/build/web-firewalld/resource/db.sqlite3
rm -rf $current_dir/build/web-firewalld/manifest/config/config.yaml
mv $current_dir/build/web-firewalld/manifest/config/config.yaml.prod $current_dir/build/web-firewalld/manifest/config/config.yaml

rm -rf ~/firewalld-web