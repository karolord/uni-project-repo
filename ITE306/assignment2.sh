#!/bin/bash
# Karo Rasool kk19046@auis.edu.krd
# Kosar Aziz kn18011@auis.edu.krd
#Requirment 1
mkdir -p /tmp/auis
man renice > /tmp/auis/help-exported.txt


#requirement2
alias ite306="yes AUIS"


#requirement3
echo "Karo Rasool and Kosar Aziz" > /home/ite306-$HOSTNAME.txt

#requirement4
sudo find /tmp -type f | wc -l
# citation find manual

#requirement 5
echo $1 $2 $3 $4 $5 $6 $7 $8 $9
var=$(echo $1 $2 $3 $4 $5 $6 $7 $8 $9 | wc -w)
if [[ $var -lt 3 ]]
then
echo Error
exit
fi
#requirement6
v1=$(ps aux | awk 'FNR == 2 {print $9}')
V2=$(date -d "$v1 today + 360 minutes" +'%H:%M')
shutdown $V2