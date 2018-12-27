package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
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

func Test_http_get_01(t *testing.T) {
	resp, err := http.Get("http://api.sh.gns.alibaba-inc.com/gns/armory/query?ip=11.167.68.134")
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
