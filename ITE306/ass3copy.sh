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
chage $username1 -M 14
chage $username2 -M 14
chage $username3 -M 14

#Requirement 4
ps -u auis_test

#Requirement 5
cd /tmp/auis
mv *.config /home/kn18-00011/Desktop

#Requirement 6
chmod -wx firefox

#Requirement 7
