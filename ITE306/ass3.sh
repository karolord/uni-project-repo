#!/bin/bash
username1=“user-$RANDOM”
username2=“user-$RANDOM”
username3=“user-$RANDOM”
groupadd firefox
useradd -m -g firefox “$username1”
useradd -m -g firefox “$username2”
useradd -m -g firefox “$username3”

chgrp firefox /usr/bin/firefox
chmod 750 /usr/bin/firefox