package main

import (
	log "github.com/Sirupsen/logrus"
	inject "github.com/facebookgo/inject"
	gondol "github.com/matisszilard/gondol/router"
	"github.com/matisszilard/gondol/store"
	"github.com/matisszilard/gondol/store/rethinkstore"
)

func main() {
	log.Info("Starting gondol...")

	s := store.Load()
	rs := rethinkstore.Load("localhost:32769")
	err := inject.Populate(s, rs.Users)
	if err != nil {
		panic(err)
	}

	gondol.Serve()
}
