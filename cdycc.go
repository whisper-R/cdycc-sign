package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main() {

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWD")

	cookie, err := sendFirstRequest(username, password)
	if err != nil {
		fmt.Println("Error sending first request:", err)
		return
	}

	responseBody, err := sendSecondRequest(cookie)
	if err != nil {
		fmt.Println("Error sending second request:", err)
		return
	}
	//fmt.Println(responseBody)
	fmt.Println(string(responseBody))
}

// 发送第一个请求
func sendFirstRequest(username, password string) (string, error) {
	url := "https://www.cdycc.cn/zb_users/plugin/freeUser/cmd.php?act=verify"
	payload := fmt.Sprintf("username=%s&edtPassWord=%s", username, password)

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("Host", "www.cdycc.cn")
	req.Header.Add("Connection", "close")
	req.Header.Add("Content-Length", "60")
	req.Header.Add("sec-ch-ua", `";Not A Brand";v="99", "Chromium";v="88"`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Referer", "https://www.cdycc.cn/?Login")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cookieStrings []string
	for _, cookie := range resp.Cookies() {
		cookieStrings = append(cookieStrings, cookie.String())
	}

	setCookie := strings.Join(cookieStrings, "; ")
	return setCookie, nil
}

// 发送第二个请求
func sendSecondRequest(cookie string) ([]byte, error) {
	url := "https://www.cdycc.cn/zb_users/plugin/freeUser/common/sign.php"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Host", "www.cdycc.cn")
	req.Header.Add("Connection", "close")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("sec-ch-ua", `";Not A Brand";v="99", "Chromium";v="88"`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	req.Header.Add("Origin", "https://www.cdycc.cn")
	req.Header.Add("Referer", "https://www.cdycc.cn/?User")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")

	// 设置上一个请求中获取到的 cookie
	req.Header.Set("Cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}