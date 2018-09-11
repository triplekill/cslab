#!/bin/bash
for k in $( seq 1 10000 )
do
   curl "localhost:8000/hi"
   echo $k
done
