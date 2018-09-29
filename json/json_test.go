package json

import (
	"testing"
	"encoding/json"
)

func Test_http_get(t *testing.T){

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


json.Unmarshal([]byte(s) , serviceList)
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
