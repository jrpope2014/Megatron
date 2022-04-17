#!/bin/bash

yum update -y

yum install -y python3 \
    git


python3 -m pip install ansible
