package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	_proto "go-tutorial/protobuf/proto"
	"go-tutorial/protobuf/util"
	"testing"
)

func Test_Ticker(t *testing.T) {

	msg := "Hello"
	cnt := 5

	user := &_proto.UserInfo{
		Message: msg,
		Length:  *proto.Int(len(msg)),
		Cnt:     *proto.Int(cnt),
	}

	o, _ := util.MessageToStruct(user)

	fmt.Println(o)
}
