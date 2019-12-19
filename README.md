# golang_code
Golang示例代码集合

![Image of Yaktocat](https://pic4.zhimg.com/v2-c4b4c4cac2beb43f81ec436291d9b203_1200x500.jpg)

## gRPC ##
`calculator` 目录提供三种gRPC API类型示例,分别是:
- unary
- client streaming
- bi directional streaming

![Image of Yaktocat](https://github.com/b43646/golang_code/blob/master/api_type.png)


另外该示例默认开启服务端的GRPC反射服务。

```
[root@demo go]# evans -p 50001 -r

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


calculator.CalculatorService@127.0.0.1:50001> show package
+------------+
|  PACKAGE   |
+------------+
| calculator |
+------------+

calculator.CalculatorService@127.0.0.1:50001> show service
+-------------------+---------+----------------+-----------------+
|      SERVICE      |   RPC   |  REQUEST TYPE  |  RESPONSE TYPE  |
+-------------------+---------+----------------+-----------------+
| CalculatorService | Sum     | SumRequest     | SumResponse     |
| CalculatorService | FindMax | FindMaxRequest | FindMaxResponse |
| CalculatorService | GetMax  | FindMaxRequest | FindMaxResponse |
+-------------------+---------+----------------+-----------------+

calculator.CalculatorService@127.0.0.1:50001> show message
+-----------------+
|     MESSAGE     |
+-----------------+
| FindMaxRequest  |
| FindMaxResponse |
| SumRequest      |
| SumResponse     |
+-----------------+

calculator.CalculatorService@127.0.0.1:50001> call Sum
first_num (TYPE_INT32) => 123
second_num (TYPE_INT32) => 123
{
  "sumResult": 246
}

calculator.CalculatorService@127.0.0.1:50001> call GetMax
num (TYPE_INT32) => 1
num (TYPE_INT32) => {
  "maxNum": 1
}
num (TYPE_INT32) => 8
num (TYPE_INT32) => {
  "maxNum": 8
}
num (TYPE_INT32) => 3
num (TYPE_INT32) => 9
num (TYPE_INT32) => {
  "maxNum": 9
}
num (TYPE_INT32) => 12
num (TYPE_INT32) => {
  "maxNum": 12
}
num (TYPE_INT32) => 3
num (TYPE_INT32) => 14
num (TYPE_INT32) => {
  "maxNum": 14
}
num (TYPE_INT32) => 12
num (TYPE_INT32) => 99
num (TYPE_INT32) => {
  "maxNum": 99
}
num (TYPE_INT32) => 100
num (TYPE_INT32) => {
  "maxNum": 100
}

```

`helloworld` 目录提供了gRPC SSL配置示例。

## CGO ##
提供cgo示例代码

