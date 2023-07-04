package core

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CtyunClient is the base struct of service clients
type CtyunClient struct {
	Credential  Credential
	Config      Config
	ServiceName string
	Revision    string
	Logger      Logger
}

type SignFunc func(*http.Request) error

func (c CtyunClient) Send(request RequestInterface) ([]byte, error) {
	method := request.GetMethod()
	builder := GetParameterBuilder(method, c.Logger)
	jsonReq, _ := json.Marshal(request)
	query, err := builder.BuildURL(jsonReq)
	if err != nil {
		return nil, err
	}
	reqUrl := fmt.Sprintf("%s://%s/%s", c.Config.Scheme, c.Config.Endpoint, request.GetURL())
	body, err := builder.BuildBody(jsonReq)
	if err != nil {
		return nil, err
	}
	c.Logger.Log(LogInfo, "request method: "+method)
	c.Logger.Log(LogInfo, "request reqUrl: "+reqUrl)
	c.Logger.Log(LogInfo, "request body: "+body)
	c.Logger.Log(LogInfo, "request query: "+query)
	return c.doSend(method, reqUrl, body, query, request.GetHeaders())
}

func (c CtyunClient) doSend(method, url, body string, query string, header map[string]string) ([]byte, error) {

	req, err := c.Http(method, url, c.Credential.AccessKey, c.Credential.SecretKey, query, body, header)

	if err != nil {
		c.Logger.Log(LogFatal, err.Error())
		return nil, err
	}

	return req, nil
}

func (c CtyunClient) setHeader(req *http.Request, header map[string]string) {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("JdcloudSdkGo/%s %s/%s", Version, c.ServiceName, c.Revision))

	base64Headers := []string{HeaderCtcloudPrefix + "-pin", HeaderCtcloudPrefix + "-erp", HeaderCtcloudPrefix + "-security-token",
		HeaderCcloudPrefix + "-pin", HeaderCcloudPrefix + "-erp", HeaderCcloudPrefix + "-security-token"}

	for k, v := range header {
		if includes(base64Headers, strings.ToLower(k)) {
			v = base64.StdEncoding.EncodeToString([]byte(v))
		}

		req.Header.Set(k, v)
	}

	for k, v := range req.Header {
		c.Logger.Log(LogInfo, k, v)
	}
}
