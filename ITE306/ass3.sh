#!/bin/bash
#Requirement 7
echo “Enter a welcome message: “ 
read welcome
echo "if who -u | grep -q “^khaosking”; then
	echo “$welcome”
fi" >> ~/.bashrc

