#!/bin/bash
username1=“user-$RANDOM”
username2=“user-$RANDOM”
username3=“user-$RANDOM”

useradd -m $username1
useradd -m $username2
useradd -m $username3

echo -e "$x\n$x"| passwd  $username1
echo -e "$x\n$x"| passwd  $username2
echo -e "$x\n$x"| passwd  $username3
passwd -x 14 $username1
passwd -x 14 $username2
passwd -x 14 $username3