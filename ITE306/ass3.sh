#!/bin/bash
#Requirement 7
echo “Enter a welcome message: “ 
read  $welcome 
sudo echo "if '${USER}' == khaosking; then
	echo “$welcome”
fi" >> /etc/profile
