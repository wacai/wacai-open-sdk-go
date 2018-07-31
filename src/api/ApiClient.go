package api

import(
	"strconv"
	"time"
	"bytes"
	"net/http"
	"io/ioutil"
	"tools"
	"config"
	"fmt"
)

type ApiClient struct{
	ApiName string
	ApiVersion string
}

func (apiClient ApiClient) HttpPostJson(jsonData string) string{
	// ApiName
	apiName := apiClient.ApiName
	// ApiVersion
	apiVersion := apiClient.ApiVersion

	// 时间戳
	var timeStamp = time.Now().UnixNano() / 1000000
	// 时间戳转字符串
	var timeStampStr = strconv.FormatInt(timeStamp,10);

	fmt.Println(timeStampStr)
	
	// Head String
	var headerString = "x-wac-app-key=" + config.APP_KEY + "&x-wac-timestamp=" + timeStampStr + "&x-wac-version=" + config.X_WAC_VERSION
	fmt.Println(headerString)
	// md5摘要
	var bodyMd5 = tools.MD5(jsonData);
	
	// 待签名
	var plainSignText = apiName + "|" + apiVersion + "|" + headerString + "|" + bodyMd5
	fmt.Println(plainSignText)

	// 生成的签名
	var signature = tools.HMAC256(plainSignText, config.APP_SECRET);
	fmt.Println(signature)

	// Api请求的路径
	var path = "/gw/api_entry/" + apiName + "/" + apiVersion;
	url := config.WACAI_GW_BASE_URL + path

	var binJson = []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(binJson))
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("x-wac-version", config.X_WAC_VERSION)
	req.Header.Set("x-wac-timestamp", timeStampStr)
	req.Header.Set("x-wac-app-key", config.APP_KEY)
	req.Header.Set("x-wac-signature", signature)

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}
