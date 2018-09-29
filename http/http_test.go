package http

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
)

func Test_http_get(t *testing.T) {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
	}

	fmt.Println(string(body))
}