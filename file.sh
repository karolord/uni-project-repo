#!/bin/bash
#for loop
ls -ali | grep | awk '{print $10}' | xargs -n 1 -I % ./script %

