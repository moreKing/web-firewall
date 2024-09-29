#!/bin/bash

rm -rf ~/firewalld-web
cp -r /mnt/hgfs/firewalld/web ~/firewalld-web
cd ~/firewalld-web
pnpm install
pnpm run build
cd -

cd server
go mod tidy
CGO_ENABLED=0 go build -o server
cd ..
rm -rf build
mkdir -p build/web-firewalld

cp script/* build/
cp -r server/manifest build/web-firewalld/
cp -r server/resource build/web-firewalld/
cp -r server/server build/web-firewalld/

rm -rf build/web-firewalld/resource/public/html
cp -r ~/firewalld-web/dist build/web-firewalld/resource/public/html

rm -rf build/web-firewalld/resource/*.sqlite3
mv build/web-firewalld/resource/db.sqlite3.bak build/web-firewalld/resource/db.sqlite3
