# package http

此包覆盖了标准库中的http实现，关于http的所有操作需要在此包进行

## server.go

负责监听tcp端口，并打开web服务器，并使用`handler`处理和回应请求

## handler.go

负责处理`http`请求，将请求转发至相关业务handler

## xxx_serice.go

处理具体业务逻辑请求处理
