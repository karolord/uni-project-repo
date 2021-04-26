#!/bin/bash
#Requirement 7
echo “Enter a welcome message: “ 
read welcome
    .bashrc
if who -u | grep -q “^$username1”; then
	echo “$welcome”
fi
