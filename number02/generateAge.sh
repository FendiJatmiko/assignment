#!/bin/bash


for i in {1..10000} 
do
  echo ${RANDOM:0:2} >> age.txt

done
