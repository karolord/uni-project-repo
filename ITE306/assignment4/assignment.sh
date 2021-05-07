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



