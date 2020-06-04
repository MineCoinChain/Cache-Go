# 						缓存系统实现-GO

### 1.HTTP传输测试结果

采用redis-benchmark作为对比，单客户端发送100000次请求，数据长度为1000B。redis测试结果如下，RPS大概在12700左右：

![1591269642452](./src/doc/redis-http.png)


