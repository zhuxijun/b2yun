# package mssql

## client.go

数据库客户端，负责打开关闭数据库，并返回数据库连接对象

## session.go

数据库连接对象, 负责数据库单次连接，身份认证

## xxx_service.go

具体接口实现，负责`package root`中的接口实现

## xxx_model.go

数据层模型定义