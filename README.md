# tfjl-h5
 塔防精灵H5版本（本仓库仅供学习交流使用，不得用于非法用途，否则后果自负）

## 项目介绍
塔防精灵H5版本，前端使用JavaScript，后端使用golang，数据库使用MongoDB，使用Websocket通信。前端源码来自于官方，后端源码由本人开发。

## 项目演示
[塔防精灵H5](http://xiaeer.top/tfjlh5/)
账号：tfjl1
密码：tfjl666
账号：tfjl2
密码：tfjl666

## 本地搭建
1. 安装MongoDB，创建数据库`tfjl`，配置好数据库的账号密码，并使用mongorestore相关命令导入`dump`文件夹下的游戏数据（若没有mongorestore命令，可以直接使用根目录下的`json_import.py`导入，需要安装python环境，运行`python json_import.py`命令），然后在`db文件夹`下的`dbconnection.go`文件中的InitDatabase方法中配置好数据库账号密码
2. 将本仓库下载到本地，打开仓库所在目录的命令行，执行`go mod tidy`下载相关依赖，执行`go run main.go`，启动后端服务
3. 访问`http://localhost:8080/tfjlh5/`，输入账号密码即可进入游戏
4. 双开小号开房间，需要搭建对战服务器，参考[tfjl-h5-fight](https://github.com/Xiaeer/tfjl-h5-fight)
