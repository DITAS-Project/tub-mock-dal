package dal

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	strict bool
	port   int
}

func NewDalServer(strict bool, port int) *Server {

	return &Server{
		strict: strict,
		port:   port,
	}
}

func (dal *Server) StartAndListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", dal.port))

	if err != nil {

	}

	server := grpc.NewServer()

	RegisterDalServer(server, dal)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (dal *Server) Query(context.Context, *QueryRequest) (*QueryReply, error) {
	if dal.strict {
		//TODO: validate token
	}

	result := make(map[string]*Rows)
	result["msg"] = &Rows{Value: []string{"Hello Dal!"}}

	reply := &QueryReply{
		Error:  nil,
		Result: &Result{Result: result},
	}

	return reply, nil
}
