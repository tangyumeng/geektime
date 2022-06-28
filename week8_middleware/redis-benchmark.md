
###  Go进阶训练营第8周作业

1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

   >命令行命令： 
   >
   >```shell l
   >redis-benchmark -d 10 -t get,set
   >redis-benchmark -d 20 -t get,set
   >redis-benchmark -d 50 -t get,set
   >redis-benchmark -d 100 -t get,set
   >redis-benchmark -d 200 -t get,set
   >redis-benchmark -d 1000 -t get,set
   >redis-benchmark -d 5000 -t get,set
   >
   >```

### 结果

#### set 

|      | 执行次数和耗时                            | 每秒请求次数                  |
| ---- | ----------------------------------------- | ----------------------------- |
| 10   | 100000 requests completed in 0.99 seconds | 101112.23 requests per second |
| 20   | 100000 requests completed in 1.19 seconds | 83752.09 requests per second  |
| 50   | 100000 requests completed in 1.16 seconds | 85984.52 requests per second  |
| 100  | 100000 requests completed in 1.18 seconds | 84388.19 requests per second  |
| 200  | 100000 requests completed in 1.22 seconds | 81833.06 requests per second  |
| 1000 | 100000 requests completed in 1.00 seconds | 100100.10 requests per second |
| 5000 | 100000 requests completed in 1.04 seconds | 95969.28 requests per second  |

### get

|      | 执行次数和耗时                            | 每秒请求次数                  |
| ---- | ----------------------------------------- | ----------------------------- |
| 10   | 100000 requests completed in 0.99 seconds | 101010.10 requests per second |
| 20   | 100000 requests completed in 1.22 seconds | 81900.09 requests per second  |
| 50   | 100000 requests completed in 1.17 seconds | 85397.09 requests per second  |
| 100  | 100000 requests completed in 1.12 seconds | 88888.89 requests per second  |
| 200  | 100000 requests completed in 1.20 seconds | 83056.48 requests per second  |
| 1000 | 100000 requests completed in 1.02 seconds | 97751.71 requests per second  |
| 5000 | 100000 requests completed in 1.04 seconds | 95785.44 requests per second  |

