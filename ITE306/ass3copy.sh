#!/bin/bash
#Requirement 1
username1=“$RANDOM”
username2=“$RANDOM”
username3=“$RANDOM”
useradd -m user-“$username1”
useradd -m user-“$username2”
useradd -m user-“$username3”

#Requirement 2
useradd -m -g ite306 auis_test

#Requirement 3
echo “$x” | passwd “$username1”
echo “$x” | passwd “$username2”
echo “$x” | passwd “$username3”
$username1 -M 14
$username2 -M 14
$username3 -M 14

#Requirement 4
top -U auis_test
ps -u auis_test

#Requirement 5
cd /tmp/auis
mv *.config /home/kn18-00011/Desktop

#Requirement 6
chmod -wx firefox

#Requirement 7
