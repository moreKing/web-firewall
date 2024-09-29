#!/bin/bash

output_file="install.log"
exec > >(tee -a "$output_file") 2>&1

systemctl disable web-firewalld
systemctl stop web-firewalld

rm -rf /usr/local/web-firewalld
rm -rf  /etc/systemd/system/web-firewalld.service

systemctl daemon-reload

echo "Web Firewalld uninstall successfully."