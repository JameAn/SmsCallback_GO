# SmsCallback_Go

#### 项目介绍
用于sp运营商对用户上行短信的callback回调，将数据分解subcode分发到php Api

#### 软件架构
软件架构说明
利用ngx_lua项目做前端网关，配置大汉三通callback服务器ip白名单，将流量代理到本项目服务器，单台服务器本项目配置默认配置两个server做平滑重启更新


#### 安装教程

1. 搭建golang基础环境，配置GOPATH，PATH等环境变量
2. 在GOPATH/src目录下clone本项目，配置项目根目录conf下的各个参数
3. 执行go install 安装文件

#### 使用说明

1. 执行根目录下的serviceStart.sh，默认开启开启两个server，占用8001，8002端口
2. 新代码更新后重新执行go install，并重新执行serviceStart.sh。自动平滑关闭服务进程并启动新进程
3. 执行serviceStop结束服务进程


