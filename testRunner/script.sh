#!/bin/bash

git clone $1 -b $2
echo 100 > report.txt
echo "big text" >> report.txt
echo "biggg" >> report.txt
