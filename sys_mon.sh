#!/bin/bash


echo "===== Sbor metrik =====" >> server_report.log
date >> server_report.log
if ping -c 1 8.8.8.8 &> /dev/null
then
	echo "Working"
else
	echo "Error"
fi
echo "Active: " >> server_report.log
pgrep -c bash >> server_report.log
echo "================================" >> server_report.log
echo "Done."


