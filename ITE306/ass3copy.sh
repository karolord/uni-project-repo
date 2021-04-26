# !/bin/bash
# Kosar N. Aziz kn18011@auis.edu.krd
# Karo K. Rasool kk19046@auis.edu.krd
# Citations:
# https://www.howtoforge.com/linux-chpasswd-command/
# https://www.cyberciti.biz/faq/protect-command-by-configuring-linux-unix-group-permissions/
# https://unix.stackexchange.com/questions/67503/move-all-files-with-a-certain-extension-from-multiple-subdirectories-into-one-di

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
mkdir tmp/auis
find /tmp/auis -name '*.config' -exec cp {} /home/$username1 \;
find /tmp/auis -name '*.config' -exec cp {} /home/$username2 \;
find /tmp/auis -name '*.config' -exec cp {} /home/$username3 \;
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

read -p “Enter a welcome message: “ $welcome >> .bashrc

echo Enter a welcome message
read $welcome >> .bashrc


echo Enter a welcome message
read $welcome
echo $welcome >> .bashrc
