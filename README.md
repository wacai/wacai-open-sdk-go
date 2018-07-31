# wacai-open-sdk-go
The client sdk(now, it's at demo code phase) of Golang for wacai open platform

## API网关
### 接口协议
- 使用HTTPS协议作为目前的交互协议
- 客户端统一使用POST方式向网关入口提交数据
- 请求报文、响应报文格式都是JSON，content_type为application/json
- 交互的编码格式统一为UTF-8
- HTTP正常响应的http code都是200，非正常返回400

### 使用配置
- 测试环境
	- 申请app_key/app_secret,向挖财开放平台申请访问app_key/app_secret,申请后 进行app_key/app_secret替换,替换为步骤1申请的(在/config/AppConfig.go中修改)
- 生产环境
	- 申请app_key/app_secret,向挖财开放平台申请访问app_key/app_secret,申请后 进行app_key/app_secret替换,替换为步骤1申请的(在AppConfig.go中修改)
	- 修改地址(生产环境),系统上线时，需要修改网关地址(在/config/AppConfig.go中修改)

### 使用实例
```go
package main

import(
	"fmt"
	"api"
)

func main() {	
	// api name
	var apiName = "api.test.post.fixed";
	
	// api version
	var apiVersion = "1.0";

	var apiClient = &api.ApiClient{ApiName: apiName,ApiVersion: apiVersion}
	
	// request data(json)
	var jsonData = "{\"uid\":123,\"name\":\"zy\"}"
	
	// start invoke...
	var result = apiClient.HttpPostJson(jsonData)
	
	// Print response
	fmt.Println(result)
}
```