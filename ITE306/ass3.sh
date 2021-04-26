#!/bin/bash
#Requirement 7
user="karo"
echo -e $user':'$1 | sudo chpasswd
passwd -x 14 $user

