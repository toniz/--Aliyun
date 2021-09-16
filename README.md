# 阿里云链路追踪服务接入DEMO
演示如何使用github.com/toniz/otel/zipkin为Go程序接入到阿里云的链路追踪服务。

## 目录说明
* server HTTP服务编写举例
* client 客户端Go程序编写举例
* pack1 pack2 演示HTTP服务调用外部包如何传递调用关系。

## 使用
* 开通aliyun的链路追踪服务服务创建,获取接入的zikpin链接.
* 填写server和client的tracing_config.json文件。
* 编译和运行

## 结果演示：
调用结果在阿里云的SLS服务里面查看：
![result1](doc/调用结果1.png)  
--- 
![result2](doc/调用结果2.png)  
---  
![result3](doc/调用结果3.png)  
---


etc..

