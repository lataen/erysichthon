package http

import (
	"github.com/rubeers/erysichthon/service/cluster"
	"github.com/rubeers/erysichthon/service/storage"
	"net/http"
)

type Server struct {
	storage.Cache
	cluster.Node
}

func (s *Server) Listen() {
	http.Handle("/storage/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.Handle("/cluster", s.clusterHandler())
	_ = http.ListenAndServe(s.Addr()+":12345", nil)
}

func New(c storage.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
