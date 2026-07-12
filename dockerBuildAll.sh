#!/bin/bash
set -e

for app in kptn-main-frontend kptn-app1-frontend kptn-app1-backend; do
  cd ${app}
  ./dockerBuild.sh
  cd ..
done   

