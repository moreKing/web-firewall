#!/bin/bash


output_file="install.log"
exec > >(tee -a "$output_file") 2>&1

cp -rf web-firewalld /usr/local/web-firewalld
chmod +x  /usr/local/web-firewalld/server
cp -rf web-firewalld.service /etc/systemd/system/

systemctl daemon-reload
systemctl enable web-firewalld
systemctl start web-firewalld

mv /usr/local/web-firewalld/server /usr/local/web-firewalld/web-firewalld

echo "Web Firewalld installed successfully."
