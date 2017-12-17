#!/bin/sh

echo "getting dependencies"
apt install fbi golang inotify-tools

echo "installing deamon"
cp media_screen.service /lib/systemd/system/
systemctl enable media_screen.service

