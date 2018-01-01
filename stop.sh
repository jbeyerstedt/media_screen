#!/bin/sh

cd /home/pi/media_screen

sudo killall fbi

sudo supervisorctl stop video_looper && cd /etc/supervisor/conf.d && sudo mv video_looper.conf video_looper.conf.disabled

