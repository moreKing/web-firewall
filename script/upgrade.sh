output_file="upgrade.log"
exec > >(tee -a "$output_file") 2>&1

cp -rf web-firewalld/resource/public /usr/local/web-firewalld/resource/public
cp -f web-firewalld/server /usr/local/web-firewalld/web-firewalld
chmod +x  /usr/local/web-firewalld/web-firewalld

cp -f web-firewalld.service /etc/systemd/system/

systemctl daemon-reload
systemctl restart web-firewalld


echo "Web Firewalld update successfully."