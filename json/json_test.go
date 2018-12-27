package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_json_02(t *testing.T) {

	s := `{
    "data":{
        "appName":"common-app",
        "app_use_type":"DAILY",
        "dns_ip":"11.167.68.134",
        "fromPeer":false,
        "id":33238740,
        "ipGroup":"",
        "nodegroup":"ali-service-platform_testhost",
        "nodename":"ali-service-platform011167068134.zth",
        "site":"ZTH",
        "sn":"6673dc4c-b9d7-4eae-af33-09e291ebfbaa",
        "state":"working_online",
        "tags":[],
        "timestamp":1543820579762,
        "unit":"CENTER_UNIT.center",
        "uuid":195511430
    },
    "success":true
}`

	gns := &GNS{}
	json.Unmarshal([]byte(s), gns)
	fmt.Println(gns.Data)
	fmt.Println(gns.Success)
}

func Test_http_get(t *testing.T) {

	s := `[
    {
        "serviceName": "caogu.2",
        "clusterMap": {
            "DEFAULT": {
                "hosts": [
                    {
                        "valid": true,
                        "metadata": {},
                        "port": 8083,
                        "ip": "20.18.7.11",
                        "weight": 1
                    }
                ]
            }
        },
        "metadata": {}
    },
    {
        "serviceName": "caogu.1",
        "clusterMap": {
            "DEFAULT": {
                "hosts": [
                    {
                        "valid": true,
                        "metadata": {},
                        "port": 8083,
                        "ip": "20.18.7.12",
                        "weight": 1
                    },
                    {
                        "valid": true,
                        "metadata": {},
                        "port": 8083,
                        "ip": "20.18.7.11",
                        "weight": 1
                    }
                ]
            }
        },
        "metadata": {}
    }
]`

	serviceList := &[]Service{}

	json.Unmarshal([]byte(s), serviceList)
}

type Service struct {
	Name    string              `json:"serviceName"`
	Cluster map[string]*Cluster `json:"clusterMap"`
	Meta    map[string]string   `json:"metadata"`
}

type Cluster struct {
	Hosts []*Host `json:"hosts"`
}

type Host struct {
	Valid  bool              `json:"valid"`
	Meta   map[string]string `json:"metadata"`
	Port   int               `json:"port"`
	Ip     string            `json:"ip"`
	Weight int               `json:"weight"`
}

type GNS struct {
	Data    Data   `json:"data"`
	Success string `json:"success"`
}

type Data struct {
	AppName string `json:"appName"`
	Site    string `json:"site"`
}
