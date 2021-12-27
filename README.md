# Trace Mocker

`TM` = `TraceMocker` 的缩写

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