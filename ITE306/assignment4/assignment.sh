#!/bin/bash
#Requirement 1
lxc launch image:ubuntu/21.04 Webserver
#Requirement 2
lxc exec Webserver apt install apache2
lxc exec Webserver -- sudo systemctl enable apache2
#Requirement 3
lxc exec Webserver pidof apache2
lxc exec Webserver -- sudo renice -n 19 $(lxc exec Webserver pidof apache2)
#requirement 4 
lxc exec Webserver apt install maillutils
string1='#!/bin/bash
sudo renice -n 19 $(pidof apache2)
echo "hello" | mail -s "notification" kk19046@auis.edu.krd
echo "hello" | mail -s "notification" kn18011@auis.edu.krd'
string2='[Unit]
Description="this will change the priority of apache2 and notify the users"
[Service]
Execstart=/etc/systemd/system/script.sh'
echo "$string1" >> script.sh
echo "$string2" >> email-apache.service
lxc file push script.sh Webserver/etc/systemd/system/
lxc file push email-apache.service Webserver/etc/systemd/system/
lxc exec Webserver -- sudo chmod +x /etc/systemd/system/script.sh
lxc exec Webserver -- sudo systemctl daemon-reload
lxc exec Webserver -- sudo systemctl start email-apache.service
#requirement 5
string3="[Unit]
Description=starts the email-apache service
Requires=email-apache.service
[Timer]
Unit=email-apache.service
OnUnitActiveSec=259200s
Persistant=true
[Install]
WantedBy=timers.target"
echo "$string3" >> email-apache.timer
lxc file push email-apache.timer Webserver/etc/systemd/system/
lxc exec Webserver -- sudo systemctl daemon-reload
lxc exec Webserver -- sudo systemctl start email-apache.timer
lxc exec Webserver -- sudo systemctl enable email-apache.timer
#Requirement 6
echo "pass123" | scp lastscript.sh ite306@10.10.10.2
