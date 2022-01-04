# Trace Mocker

`TM` = `TraceMocker` 的缩写

提供链路测试的能力，通过指定的 URL 模拟请求。

## 路由列表
### SimpleMock - 基础模拟

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

## 任务管理

### TraceMocker 提供了任务管理能力

可以通过 TraceMocker 创建任务并写入到一个Object-Mocker中，实现任务的自动调度。
此过程中，TraceMocker 会自动从 Object-Mocker 中同步任务，并加载对应 Holder 与 配置中 NodeId 的相关数据。

一个任务可以包含多个请求。