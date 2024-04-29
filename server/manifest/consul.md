# consul

请进入/opt/consul目录执行 docker-compose up -d 启动consul
consul的管理员token是: 3563a19f-756e-47e4-9ae6-3694bd7bead8

docker run -d -p 8500:8500 --restart=always 
    --name=consul consul:1.9.5 agent 
    -server 
    -bootstrap 
    -ui 
    -node=node1 
    -client='0.0.0.0'