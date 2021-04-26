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
echo “$x” | passwd -x 14 “$username1”
echo “$x” | passwd -x 14 “$username2”
echo “$x” | passwd -x 14 “$username3” 
passwd -x 14 $username1
passwd -x 14 $username2
passwd -x 14 $username3
#Requirement 4
ps -u auis_test

#Requirement 5
cd /tmp/auis
mv *.config /home/kn18-00011/Desktop

#Requirement 6
chmod -wx firefox

#Requirement 7
