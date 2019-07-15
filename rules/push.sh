#!/bin/bash
############
if [ $# == 0 ];then
    git add -A
    git commit -m 'chore: fix dp'
    git push
else
    git add -A
    git commit -m "$1"
    git push
fi
