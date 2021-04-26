#!/bin/bash
username1=“$RANDOM”
username2=“$RANDOM”
username3=“$RANDOM”
useradd -m user-“$username1”
useradd -m user-“$username2”
useradd -m user-“$username3”


useradd -m -g ite306 auis_test

echo “$x” | passwd “$username1”
echo “$x” | passwd “$username2”
echo “$x” | passwd “$username3”


top -U auis_test
ps -u auis_test


cd /tmp/auis
mv *.config /home/kn18-00011/Desktop


chmod -wx firefox
