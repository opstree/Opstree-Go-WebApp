#!/bin/bash

set -ex
sleep 10
service filebeat start

./ot-go-webapp
