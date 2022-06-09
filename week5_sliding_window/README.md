### 模仿Hystrix 实现滑动窗口计数



#### 思路

实现滑动窗口代码位置 ./pkg/hystrix

- 定义 Bucket 防止请求结果对象(./week5_sliding_window/hystrix/bucket.go)
- 定义滑动窗口对象 (./week5_sliding_window/hystrix/rolling-window.go)
  - FIFO bucket 切片

#### testcase 

提供 Client 方法， 用于在 cmd-> client-> main 中进行高并发调用，模拟业务请求，进而触发熔断

#### cmd 

client 模拟客户端请求

downstream 模拟下游请求

upstream 模拟上游请求