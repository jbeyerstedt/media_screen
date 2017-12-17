#!/bin/sh

cd /home/pi/media_screen

while inotifywait -q -e modify,create,delete,move /home/pi/video/ .
do
  go run media_screen.go
done

