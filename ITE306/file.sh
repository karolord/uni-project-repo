#!/bin/bash
v1=$(ps aux | awk 'FNR == 2 {print $9}')
V2=$(date -d "$v1 today + 360 minutes" +'%H:%M')
shutdown $V2