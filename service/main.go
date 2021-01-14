package main

import (
	"flag"
	"github.com/rubeers/erysichthon/service/cluster"
	"github.com/rubeers/erysichthon/service/http"
	"github.com/rubeers/erysichthon/service/storage"
	"github.com/rubeers/erysichthon/service/tcp"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "storage type")
	ttl := flag.Int("ttl", 30, "storage time to live")
	node := flag.String("node", "127.0.0.1", "node address")
	clus := flag.String("cluster", "", "cluster address")
	flag.Parse()
	log.Println("type is", *typ)
	log.Println("ttl is", *ttl)
	log.Println("node is", *node)
	log.Println("cluster is", *clus)
	c := storage.New(*typ, *ttl)
	n, e := cluster.New(*node, *clus)
	if e != nil {
		panic(e)
	}
	go tcp.New(c, n).Listen()
	http.New(c, n).Listen()
}
