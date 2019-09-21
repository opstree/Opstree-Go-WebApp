#!/bin/bash

set -ex
sleep 25
service filebeat start

./ot-go-webapp
