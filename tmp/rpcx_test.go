package main

import (
	"context"
	"testing"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/client"
	"log"
)

type Args struct{
	A int
	B int
}

type Reply struct{
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error{
	reply.C = args.A * args.B
	return nil
}

func Test_server(t *testing.T){
	s := server.NewServer()
	s.RegisterName("Arith",new(Arith),"")
	s.Serve("tcp", ":8972")
}



func Test_client(t *testing.T){
	d := client.NewPeer2PeerDiscovery("tcp@"+"127.0.0.1:8972", "")

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}

	err := xclient.Call(context.Background(), "Mul", args, reply)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

