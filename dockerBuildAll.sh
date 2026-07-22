#!/bin/bash
set -e

for app in main-frontend app1-frontend app1-backend; do
  cd ${app}
  ./dockerBuild.sh
  cd ..
done   

