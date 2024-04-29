#!/usr/bin/env bash

cd ..
make build

mkdir -p ~/app/chunyu

mkdir -p ~/app/chunyu/admin/service/bin/

mkdir -p ~/app/chunyu/admin/service/configs/

mv -f ./app/admin/service/bin/server ~/app/chunyu/admin/service/bin/server

cp -rf ./app/admin/service/configs/*.yaml ~/app/chunyu/admin/service/configs/

cd ~/app/chunyu/admin/service/bin/
pm2 start --namespace chunyu --name admin server -- -conf ../configs/

pm2 save

pm2 restart chunyu