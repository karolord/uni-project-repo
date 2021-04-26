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
echo “$x” | passwd -x 14 “$username1”
echo “$x” | passwd -x 14 “$username2”
echo “$x” | passwd -x 14 “$username3” 
passwd -x 14 $username1
passwd -x 14 $username2
passwd -x 14 $username3
echo “$x” | passwd “$username1”
echo “$x” | passwd “$username2”
echo “$x” | passwd “$username3”
passwd -x "$username1" 14
passwd -x "$username2" 14
passwd -x "$username3" 14

#Requirement 4
ps -u auis_test

#Requirement 5
cd /tmp/auis
mv *.config /home/kn18-00011/Desktop

#Requirement 6
chgrp firefox /usr/bin/firefox
chmod 750 /usr/bin/firefox

#Requirement 7
