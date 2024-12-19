# pitayaproject

//初始化服务器
go mod init game

//初始化pitaya
go get github.com/topfreegames/pitaya/v2/

go mod tidy

nohup ./server --type=gw > log/gw.log 2>&1 &
nohup ./server --type=db > log/db.log 2>&1 &
nohup ./server --type=lobby > log/lobby.log 2>&1 &
nohup ./server --type=login > log/login.log 2>&1 &
nohup ./server --type=web > log/web.log 2>&1 &


初始化 pitaya
go get github.com/topfreegames/pitaya/v2/

安装 etcd
详细看 etcd.md
