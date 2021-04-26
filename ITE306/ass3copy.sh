#!/bin/bash
#Requirement 1
username1=“user-$RANDOM”
username2=“user-$RANDOM”
username3=“user-$RANDOM”
useradd -m “$username1”
useradd -m “$username2”
useradd -m “$username3”

#Requirement 2
groupadd ite306 
useradd -m -g ite306 auis_test

#Requirement 3
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
chown .profile root
alias firefox=‘/dev/null’

#Requirement 7
