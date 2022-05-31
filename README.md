![](https://img.shields.io/badge/-ceph-green)
# grpc usage (完整demo,可直接从第五步开始运行)
* **1.安装`protoc`**
```
protoc --version  #确保版本在3.0+
```
***
* 2.**创建项目,gomod引入对应包**
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
***
* 3.**编写user.proto文件,前往[proto学习](https://blog.csdn.net/xp178171640/article/details/102951328)**
```
protoc -I . users/user.proto --go_out=plugins=grpc:.  //会在users包下生成user.pb.go文件
```
***
* 4.**编写服务端代码,`server/main.go`,编写客户端代码(也可通过postman直接访问服务端,新版本已支持),`client/main.go`**
***
* 5.**先启动服务端，再启动客户端访问**
```
go run server/main.go
go run client/main.go
```
***
* **6.控制台日志**
```
//server log

2022/05/31 15:17:32 receive users index request:page 1 page_size 12
2022/05/31 15:17:32 receive users uid request:uid 1
2022/05/31 15:17:32 receive users uid request:name big_cat password:123456,age:29
2022/05/31 15:17:32 receive users uid request:uid 1

====================================================================

//client log

aaaa 28
bbbb 1
aaaa 28
2022/05/31 15:30:53 user index success: success
2022/05/31 15:30:53 user view success: success
2022/05/31 15:30:53 user post success: success
2022/05/31 15:30:53 user delete success: success
```




