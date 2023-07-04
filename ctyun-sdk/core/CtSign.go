package core

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-basic/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

func (c CtyunClient) Http(method, url, ak, sk, query, body string, header map[string]string) ([]byte, error) {

	switch method {

	case "GET":
		resp, err := c.GetHttp(url, ak, sk, header, query)
		return resp, err
	case "POST":
		resp, err := c.PostHttp(url, body, ak, sk, header, query)
		return resp, err
	case "PUT":
		resp, err := c.PutHttp(url, body, ak, sk, header, query)
		return resp, err
	}
	return nil, nil
}

func (c CtyunClient) GetHttp(url, ak, sk string, headerMap map[string]string, query string) ([]byte, error) {
	var newQuery = ""
	var queryStr = strings.Split(query, "&")
	sort.Slice(queryStr, func(i, j int) bool {
		return queryStr[i] < queryStr[j] // 正序
	})
	for _, value := range queryStr {
		newQuery = newQuery + "&" + value
	}
	afterQuery := EncodeQueryStr(newQuery)
	if afterQuery != "" {
		url = url + "?" + afterQuery
	}
	fmt.Println("afterQuery:", afterQuery)
	uuId := uuid.New()
	var (
		err error
	)
	// 准备: HTTP请求
	reqBody := strings.NewReader(string(""))
	httpReq, err := http.NewRequest(http.MethodGet, url, reqBody)
	fmt.Println(url)
	if err != nil {
		c.Logger.Log(LogError, "NewRequest fail ")
		return nil, err
	}
	tim := time.Now()
	eopDate := tim.Format("20060102T150405Z")

	httpReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	httpReq.Header.Add("ctyun-eop-request-id", uuId)
	httpReq.Header.Add("Eop-Authorization", getSign("", uuId, ak, sk, http.MethodGet, tim, afterQuery))
	httpReq.Header.Add("Eop-date", eopDate)
	for temp := range headerMap {
		httpReq.Header.Add(temp, headerMap[temp])
	}
	// DO: HTTP请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		c.Logger.Log(LogError, "do http fail", err)
		return nil, err
	}
	headers := httpReq.Header
	for k, v := range headers {
		fmt.Println(k, v)
	}

	if httpRsp.StatusCode != http.StatusOK {
		c.Logger.Log(LogInfo, "request StatusCode: ", httpRsp.StatusCode)
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	c.Logger.Log(LogInfo, "rspBody: "+string(rspBody))
	if err != nil {
		c.Logger.Log(LogError, "ReadAll rspBody failed", err)
		return nil, err
	}
	return rspBody, nil
}

func (c CtyunClient) PostHttp(url, body, ak, sk string, headerMap map[string]string, query string) ([]byte, error) {
	var newQuery = ""
	var queryStr = strings.Split(query, "&")
	sort.Slice(queryStr, func(i, j int) bool {
		return queryStr[i] < queryStr[j] // 正序
	})
	for _, value := range queryStr {
		newQuery = newQuery + "&" + value
	}
	afterQuery := EncodeQueryStr(newQuery)
	if afterQuery != "" {
		url = url + "?" + afterQuery
	}
	uuId := uuid.New()
	var (
		err error
	)
	// 准备: HTTP请求
	reqBody := strings.NewReader(body)
	httpReq, err := http.NewRequest(http.MethodPost, url, reqBody)

	if err != nil {
		c.Logger.Log(LogError, "NewRequest fail: ", err)
		return nil, err
	}
	tim := time.Now()
	eopDate := tim.Format("20060102T150405Z")

	httpReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	httpReq.Header.Add("ctyun-eop-request-id", uuId)
	httpReq.Header.Add("Eop-Authorization", getSign(body, uuId, ak, sk, http.MethodPost, tim, afterQuery))
	httpReq.Header.Add("Eop-date", eopDate)
	for temp := range headerMap {
		httpReq.Header.Add(temp, headerMap[temp])
	}
	// DO: HTTP请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		c.Logger.Log(LogError, "do http fail: ", err)
		return nil, err
	}
	headers := httpReq.Header
	for k, v := range headers {
		fmt.Println(k, v)
	}

	if httpRsp.StatusCode != http.StatusOK {
		c.Logger.Log(LogInfo, "request StatusCode: ", httpRsp.StatusCode)
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	c.Logger.Log(LogInfo, "rspBody info: "+string(rspBody))

	if err != nil {
		c.Logger.Log(LogError, "ReadAll rspBody failed: ", err)
		return nil, err
	}
	return rspBody, nil
}

func (c CtyunClient) PutHttp(url, reqParam, ak, sk string, headerMap map[string]string, query string) ([]byte, error) {
	var newQuery = ""
	var queryStr = strings.Split(query, "&")
	sort.Slice(queryStr, func(i, j int) bool {
		return queryStr[i] < queryStr[j] // 正序
	})
	for _, value := range queryStr {
		newQuery = newQuery + "&" + value
	}
	afterQuery := EncodeQueryStr(newQuery)
	if afterQuery != "" {
		url = url + "?" + afterQuery
	}
	uuId := uuid.New()
	var (
		err error
	)
	// 准备: HTTP请求
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest(http.MethodPost, url, reqBody)
	fmt.Println(url)
	if err != nil {
		fmt.Printf("NewRequest fail, url: %s, reqBody: %s, err: %v", "", reqBody, err)
		return nil, err
	}
	tim := time.Now()
	eopDate := tim.Format("20060102T150405Z")

	httpReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	httpReq.Header.Add("ctyun-eop-request-id", uuId)
	httpReq.Header.Add("Eop-Authorization", getSign(reqParam, uuId, ak, sk, http.MethodPut, tim, afterQuery))
	httpReq.Header.Add("Eop-date", eopDate)
	for temp := range headerMap {
		httpReq.Header.Add(temp, headerMap[temp])
	}
	// DO: HTTP请求
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		fmt.Printf("do http fail, url: %s, reqBody: %s, err:%v", "", reqBody, err)
		return nil, err
	}
	headers := httpReq.Header
	for k, v := range headers {
		fmt.Println(k, v)
	}

	if httpRsp.StatusCode != http.StatusOK {
		fmt.Println()
		fmt.Println("请求StatusCode: ", httpRsp.StatusCode)
		fmt.Println()
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	fmt.Println("rspBody: " + string(rspBody))
	if err != nil {
		fmt.Printf("ReadAll failed, url: %s, reqBody: %s, err: %v", "", reqBody, err)
		return nil, err
	}
	return rspBody, nil
}

func getSign(body, uuId, ak, sk, method string, tim time.Time, afterQuery string) string {

	data := []byte(body)
	// hash
	hash := sha256.New()
	hash.Write(data)
	sum1 := hash.Sum(nil)
	calculateContentHash := hex.EncodeToString(sum1)

	var sigture string

	singerDate := tim.Format("20060102T150405Z")
	singerDd := tim.Format("20060102")
	CampmocalHeader := "ctyun-eop-request-id:" + uuId + "\neop-date:" + singerDate + "\n"
	//if method == http.MethodGet {
	sigture = CampmocalHeader + "\n" + afterQuery + "\n" + calculateContentHash
	fmt.Println("sigture:", sigture)
	//} else {
	//	sigture = CampmocalHeader + "\n" + "=" + "\n" + calculateContentHash
	//}
	kSecret := sk
	ktime := HmacSHA256(singerDate, kSecret)
	fmt.Println("ktime: " + hex.EncodeToString(ktime))
	kAk := HmacSHA256(ak, string(ktime))
	kdate := HmacSHA256(singerDd, string(kAk))
	signaSha256 := HmacSHA256(sigture, string(kdate))
	Signature := base64.StdEncoding.EncodeToString(signaSha256)
	signHeader := ak + " Headers=ctyun-eop-request-id;eop-date Signature=" + Signature
	fmt.Println("signHeader: " + signHeader)
	fmt.Println()
	fmt.Println()
	return signHeader
}

func getSignByByMultipartFormDataBoundary(ak, sk, afterQuery, uuId string, tim time.Time, bodyArr []byte) string {
	fmt.Println("---------------------------开始---------------------------")
	fmt.Println(string(bodyArr))
	fmt.Println("---------------------------结束---------------------------")
	// hash
	hash := sha256.New()
	hash.Write(bodyArr)
	sum1 := hash.Sum(nil)
	calculateContentHash := hex.EncodeToString(sum1)

	var sigture string

	singerDate := tim.Format("20060102T150405Z")
	singerDd := tim.Format("20060102")
	CampmocalHeader := "ctyun-eop-request-id:" + uuId + "\neop-date:" + singerDate + "\n"
	sigture = CampmocalHeader + "\n" + afterQuery + "\n" + calculateContentHash
	fmt.Println("sigture: " + sigture)
	kSecret := sk
	ktime := HmacSHA256(singerDate, kSecret)
	fmt.Println("ktime: " + hex.EncodeToString(ktime))
	kAk := HmacSHA256(ak, string(ktime))
	fmt.Println("kAk: " + hex.EncodeToString(kAk))
	kdate := HmacSHA256(singerDd, string(kAk))
	fmt.Println("kdate: " + hex.EncodeToString(kdate))
	signaSha256 := HmacSHA256(sigture, string(kdate))
	Signature := base64.StdEncoding.EncodeToString(signaSha256)
	fmt.Println("Signature: " + Signature)
	signHeader := ak + " Headers=ctyun-eop-request-id;eop-date Signature=" + Signature
	fmt.Println("signHeader: " + signHeader)
	fmt.Println()
	fmt.Println()
	return signHeader
}

func EncodeQueryStr(query string) string {
	afterQuery := ""
	if len(query) != 0 {
		n := strings.Split(query, "&")
		for _, v := range n {
			if len(afterQuery) < 1 {
				a := strings.Split(v, "=")
				if len(a) >= 2 {
					encodeStr := url.QueryEscape(a[1])
					v = a[0] + "=" + encodeStr
					afterQuery = afterQuery + v
				} else {
					encodeStr := ""
					v = a[0] + "=" + encodeStr
					afterQuery = afterQuery + v
				}
			} else {
				a := strings.Split(v, "=")
				if len(a) >= 2 {
					encodeStr := url.QueryEscape(a[1])
					v = a[0] + "=" + encodeStr
					afterQuery = afterQuery + "&" + v
				} else {
					encodeStr := ""
					v = a[0] + "=" + encodeStr
					afterQuery = afterQuery + "&" + v
				}
			}
		}
	}

	return afterQuery
}

func HmacSHA256(signature, key string) []byte {
	s := []byte(signature)
	k := []byte(key)
	m := hmac.New(sha256.New, k)
	m.Write(s)
	sum1 := m.Sum(nil)
	return sum1
}

func MapInterface2String(inputData map[string]interface{}) map[string]string {
	outputData := map[string]string{}
	for key, value := range inputData {
		switch value.(type) {
		case string:
			outputData[key] = value.(string)
		}
	}
	return outputData
}

func String2Map(header string) map[string]string {
	data := []byte(header)
	map2 := make(map[string]interface{})
	err := json.Unmarshal(data, &map2)
	if err != nil {
		fmt.Println(err)
	}
	headerMap := MapInterface2String(map2)
	return headerMap
}

func JSONMethod(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func readByte(fileName string) []byte {
	result := ""
	//打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//当函数退出时及时关闭file
	//defer在函数结束之前会调用defer后的代码
	//及时关闭file句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		result = result + str
		if err == io.EOF { //io.EOF表示文件的末尾
			return []byte(result)
		}
	}
	return nil
}
