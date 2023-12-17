#!/bin/bash
mkdir $1
cd $1
touch $1.go $1_test.go $1_data.txt $1_test_data.txt
echo "package main" >> $1.go
echo "package main" >> $1_test.go