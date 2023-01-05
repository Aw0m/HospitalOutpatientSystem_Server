package utils

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = jsoniter.Unmarshal(body, &result)
	}

	return result, err
}
