#!/bin/bash

USAGE=$(df -h / | awk 'NR==2 {print $5}' | tr -d '%')

THRESHOLD=80

echo "tekushaya zagruzga diska: $USAGE%"

if [ "$USAGE" -gt "$THRESHOLD" ]; then
	echo "Vnimanie!Malo mesto!"
else
	echo "Vse Otlichno"
fi

