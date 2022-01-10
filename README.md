# Trace Mocker

`TM` = `TraceMocker` 的缩写

提供链路测试的能力，通过指定的 URL 模拟请求。

## 路由列表

### /mock/http/simple-request - 基础模拟

基础模拟只会使用`GET`方法请求指定地址，不会有其他行为。

- 请求方法：`ANY`
- 请求体：json

请求体结构
```golang
// SimpleRequestBody only save the next-route.
type SimpleRequestBody struct {
	NextBody  interface{} `json:"next_body"`
	NextRoute string      `json:"next_route"`
}
```

请求将直接返回下一层返回体。

### /mock/http/trace-mock - 链路模拟

链路模拟通过构造指定结构的请求体，向指定的服务发起指定方法的请求。

- 请求方法：`ANY`
- 请求体：json
请求体构造
```go
type TraceRequestBody struct {
	NextService string `json:"next_service"`
	NextMethod  string `json:"next_method"`
}

[] TraceRequestBody
```
样例：
```json
[
    {
        "next_service": "127.0.0.1:3000",
        "next_method": "POST"
    }
]
```

next_method: 默认为 POST，会将小写字符自动转换为大写字符

## 任务管理

### TraceMocker 提供了任务管理能力

可以通过 TraceMocker 创建任务并写入到一个Object-Mocker中，实现任务的自动调度。

#### NodeId

此过程中，TraceMocker 会自动从 Object-Mocker 中同步任务，并加载对应 Holder 与 配置中 node_id 的相关数据。

一个任务可以包含多个请求。

创建任务时，如果未指定 Holder ，则自动调度到创建任务的节点。多个节点如果 node_id 相同，则会同时执行任务。

NodeId 可以通过配置文件指定:
```yaml
application:
  node_id: <Your node id>
```
或者通过环境变量: `NODE_ID` 指定。

### 创建任务

