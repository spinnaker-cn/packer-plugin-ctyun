package core

import (
	"encoding/json"
	"errors"
	"fmt"
	urllib "net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var baseRequestFields []string

func init() {
	req := CtyunRequest{}
	reqType := reflect.TypeOf(req)
	for i := 0; i < reqType.NumField(); i++ {
		baseRequestFields = append(baseRequestFields, reqType.Field(i).Name)
	}
}

type ParameterBuilder interface {
	BuildURL(paramJson []byte) (string, error)
	BuildBody(paramJson []byte) (string, error)
}

func GetParameterBuilder(method string, logger Logger) ParameterBuilder {
	if method == MethodGet || method == MethodDelete || method == MethodHead {
		return &WithoutBodyBuilder{logger}
	} else {
		return &WithBodyBuilder{logger}
	}
}

// WithBodyBuilder supports PUT/POST/PATCH methods.
// It has path and body (json) parameters, but no query parameters.
type WithBodyBuilder struct {
	Logger Logger
}

func (b WithBodyBuilder) BuildURL(paramJson []byte) (string, error) {

	return "", nil
}

func (b WithBodyBuilder) BuildBody(paramJson []byte) (string, error) {
	paramMap := make(map[string]interface{})
	err := json.Unmarshal(paramJson, &paramMap)
	if err != nil {
		b.Logger.Log(LogError, err.Error())
		return "", err
	}

	// remove base request fields
	for k := range paramMap {
		if includes(baseRequestFields, k) {
			delete(paramMap, k)
		}
	}

	body, _ := json.Marshal(paramMap)
	b.Logger.Log(LogInfo, "Body=", string(body))
	return string(body), nil
}

// WithoutBodyBuilder supports GET/DELETE methods.
// It only builds path and query parameters.
type WithoutBodyBuilder struct {
	Logger Logger
}

func (b WithoutBodyBuilder) BuildURL(paramJson []byte) (string, error) {
	paramMap := make(map[string]interface{})
	err := json.Unmarshal(paramJson, &paramMap)
	if err != nil {
		return "", err
	}

	// remove base request fields
	for k := range paramMap {
		if includes(baseRequestFields, k) {
			delete(paramMap, k)
		}
	}

	var url = ""

	for k, v := range paramMap {

		switch v.(type) {
		case string:
			temp := v.(string)
			url += fmt.Sprintf("%s=%s&", k, temp)
		case float32:
			url += fmt.Sprintf("%s=%f&", k, v.(float32))
		case float64:
			temp := strconv.FormatFloat(v.(float64), 'f', -1, 32)
			url += fmt.Sprintf("%s=%s&", k, temp)
		case int:
			temp := strconv.Itoa(v.(int))
			url += fmt.Sprintf("%s=%s&", k, temp)
		case int32:
			temp := string(v.(int32))
			url += fmt.Sprintf("%s=%s&", k, temp)
		case int64:
			temp := strconv.FormatInt(v.(int64), 10)
			url += fmt.Sprintf("%s=%s&", k, temp)
		}

		//url += j + "=" + string(v.(string)) + "&"

	}
	url = url[0 : len(url)-1]
	fmt.Println(url)
	//body, _ := json.Marshal(paramMap)

	return url, nil
}

func (b WithoutBodyBuilder) BuildBody(paramJson []byte) (string, error) {
	return "", nil
}

func replaceUrlWithPathParam(url string, paramMap map[string]interface{}) (string, error) {
	r, _ := regexp.Compile("{[a-zA-Z0-9-_]+}")
	matches := r.FindAllString(url, -1)
	for _, match := range matches {
		field := strings.TrimLeft(match, "{")
		field = strings.TrimRight(field, "}")
		value, ok := paramMap[field]
		if !ok {
			return "", errors.New("Can not find path parameter: " + field)
		}

		valueStr := fmt.Sprintf("%v", value)
		url = strings.Replace(url, match, valueStr, -1)
	}

	return url, nil
}

func buildQueryParams(paramMap map[string]interface{}, url string) urllib.Values {
	values := urllib.Values{}
	accessMap(paramMap, url, "", values)
	return values
}

func accessMap(paramMap map[string]interface{}, url, prefix string, values urllib.Values) {
	for k, v := range paramMap {
		// exclude fields of JDCloudRequest class and path parameters
		if shouldIgnoreField(url, k) {
			continue
		}

		switch e := v.(type) {
		case []interface{}:
			for i, n := range e {
				switch f := n.(type) {
				case map[string]interface{}:
					subPrefix := fmt.Sprintf("%s.%d.", k, i+1)
					accessMap(f, url, subPrefix, values)
				case nil:
				default:
					values.Set(fmt.Sprintf("%s%s.%d", prefix, k, i+1), fmt.Sprintf("%s", n))
				}
			}
		case nil:
		default:
			values.Set(fmt.Sprintf("%s%s", prefix, k), fmt.Sprintf("%v", v))
		}
	}
}

func shouldIgnoreField(url, field string) bool {
	flag := "{" + field + "}"
	if strings.Contains(url, flag) {
		return true
	}

	if includes(baseRequestFields, field) {
		return true
	}

	return false
}
