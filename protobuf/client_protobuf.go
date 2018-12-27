package main

import (
	"bufio"
	"fmt"
	stProto "go-tutorial/protobuf/proto"
	"net"
	"os"
	"time"

	//protobuf编解码库,下面两个库是相互兼容的，可以使用其中任意一个
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
)

func main() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	//发送消息
	//cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		//cnt++
		//stSend := buildUserInfo(sender)
		stSend := buildFilterConfig(sender)
		//stSend := &stProto.UserInfo{
		//	Message: sender.Text(),
		//	Length:  *proto.Int(len(sender.Text())),
		//	Cnt:     *proto.Int(cnt),
		//}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}

func buildUserInfo(sender *bufio.Scanner) *stProto.UserInfo {
	cnt++
	stSend := &stProto.UserInfo{
		Message: sender.Text(),
		Length:  *proto.Int(len(sender.Text())),
		Cnt:     *proto.Int(cnt),
	}
	return stSend
}

func buildFilterConfig(sender *bufio.Scanner) *stProto.FilterConfig {

	cnt++
	routeConfig := stProto.RouteConfig{
		Name:      "com.alibaba.middleware.hsf.HelloWorldService:1.0.0.TLS",
		Interface: "com.alibaba.middleware.hsf.HelloWorldService",
		Version:   "1.0.0",
		Group:     "HSF",
		Spec: &stProto.RouteConfigSpec{
			AppName:   "xxx_business",
			WriteMode: "center_or_unit",
			RouteIdx:  1,
		},
		Routes: []*stProto.RouteConfigRoute{
			{
				Match: &stProto.RouteConfigRouteMatch{
					Method: &stProto.RouteConfigRouteMatchMethod{
						Name: "sayHello",
						ParamsMatch: []*stProto.ParameterMatch{
							{
								Index:      0,
								Type:       "int",
								RangeMatch: &stProto.RangeMatch{Start: 100, End: 200},
							},
							{
								Index:       1,
								Type:        "string",
								PrefixMatch: "user_id:",
							},
							{
								Index:      2,
								Type:       "string",
								RegexMatch: "0.[5-9]",
							},
						},
					},
					Headers: []*stProto.HeaderMatch{
						{Name: "custom1", ExactMatch: "123"},
						{Name: "custom2", ExactMatch: "456", InvertMatch: true},
					},
				},
				Route: &stProto.RouteConfigRouteRoute{},
			}, {
				Match: &stProto.RouteConfigRouteMatch{
					Method: &stProto.RouteConfigRouteMatchMethod{Name: "fooMethod"},
					Headers: []*stProto.HeaderMatch{
						{
							Name:       "custom1",
							ExactMatch: "135",
						},
						{
							Name:        "custom5",
							ExactMatch:  "489",
							InvertMatch: true,
						},
					},
				},
				Route: &stProto.RouteConfigRouteRoute{Cluster: "cluster_xxx"},
			},
		},
	}

	routeConfigs := []*stProto.RouteConfig{&routeConfig}

	filterConfig := &stProto.FilterConfig{
		StatPrefix:        "dubbo_incoming_stats",
		ProtocolType:      "Dubbo",
		SerializationType: "Hessian2",
		RouteConfig:       routeConfigs,
	}
	return filterConfig
}

var cnt = 0
