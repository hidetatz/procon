#!/bin/bash

# sed solution
sed -n 10P file.txt

# awk solution
awk 'NR==10' file.txt
