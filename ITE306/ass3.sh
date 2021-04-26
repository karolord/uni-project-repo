REQUIRMENT 1
username1=“$RANDOM”
username2=“$RANDOM”
username3=“$RANDOM”
useradd -m user-“$username1”
useradd -m user-“$username2”
useradd -m user-“$username3”

REQUIRMENT 2
useradd -m -g ite306 auis_test

REQUIRMENT 3
echo “$x” | passwd “$username1”
echo “$x” | passwd “$username2”
echo “$x” | passwd “$username3”

REQUIRMENT 4
top -U auis_test
ps -u auis_test

REQUIRMENT 5
cd /tmp/auis
mv *.config /home/kn18-00011/Desktop

REQUIRMENT 6
chmod -wx firefox

REQUIRMENT 7