#!/bin/bash

TARGET=$1
TMP=/tmp/backup.txt

echo "Backing up $TARGET"

cp $TARGET $TMP
tar -czf backup.tar.gz $TMP

echo "Backup created: backup.tar.gz"
