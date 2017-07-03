package main

import (
	log "github.com/Sirupsen/logrus"
	gondol "github.com/matisszilard/gondol/router"
	s "github.com/matisszilard/gondol/store"
)

func main() {
	log.Info("Starting gondol...")

	log.Info("Loading users from the database")
	users, err := s.GetUsers()
	if err != nil {
		log.Error("Something bad happend - {}", err)
		return
	}
	log.Info("Loaded users - {}", users)

	gondol.Serve()
}
