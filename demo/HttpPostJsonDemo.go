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


