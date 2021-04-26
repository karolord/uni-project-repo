#!/bin/bash
#Requirement 1
username1=“user-$RANDOM”
username2=“user-$RANDOM”
username3=“user-$RANDOM”
groupadd firefox
useradd -m -g firefox “$username1”
useradd -m -g firefox “$username2”
useradd -m -g firefox “$username3”

#Requirement 2
groupadd ite306 
useradd -m -g ite306 auis_test

#Requirement 3

echo $username1':'$1 | sudo chpasswd
echo $username2':'$1 | sudo chpasswd
echo $username3':'$1 | sudo chpasswd
echo test_auis':'$1”| sudo chpasswd
passwd -x "$username1" 14
passwd -x "$username2" 14
passwd -x "$username3" 14
passwd -x "test_auis" 14

#Requirement 4
ps -u auis_test

#Requirement 5
mkdir /tmp
cd /tmp
mkdir /auis
cd /auis
cp *.config /home/$username1
cp *.config /home/$username2
cp *.config /home/$username3
rm *.config

#Requirement 6
chgrp firefox /usr/bin/firefox
chmod 750 /usr/bin/firefox

#Requirement 7
echo “Enter a welcome message: “ 
read welcome
nano .bashrc
if who -u | grep -q “^$username1”; then
	echo “$welcome”
fi
