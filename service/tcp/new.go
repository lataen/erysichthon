package tcp

import (
	"github.com/rubeers/erysichthon/service/cluster"
	"github.com/rubeers/erysichthon/service/storage"
	"net"
)

type Server struct {
	storage.Cache
	cluster.Node
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", s.Addr()+":12346")
	if e != nil {
		panic(e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}
		go s.process(c)
	}
}

func New(c storage.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
