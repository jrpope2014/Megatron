#!/bin/bash

yum update -y

amazon-linux-extras install -y python3.8


python3 -m pip install ansible
