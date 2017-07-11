package rethinkstore

import (
	log "github.com/Sirupsen/logrus"
	r "gopkg.in/gorethink/gorethink.v3"
)

const (
	driver = "rethinkdb"

	// GondolDatabaseName holds the database name of the gondol project
	gondolDatabaseName = "gondol"

	// UsersTableName name of the table used for storing users
	UsersTableName = "users"
)

// Store rethink DB store
type Store struct {
	session *r.Session
	Users   *Users
}

// Load datastore settings based on enviroment
func Load(address string) *Store {
	log.Infof("rethinkstore load - using database server: %s", address)
	var s Store
	s.session = s.new(driver, address, gondolDatabaseName)
	s.Users = &Users{s.session}
	return &s
}

// New RethinkDB connection
func (s *Store) new(driver, address, database string) *r.Session {

	//TODO add automatic connection retry
	session, err := r.Connect(r.ConnectOpts{
		Address:  address,
		Database: database,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	setUpDatabase(database, session)
	return session
}

func setUpDatabase(database string, session *r.Session) {
	createDB(database, session)

	createTables(session)

	log.Info("Database setup completed")
}

func createDB(database string, session *r.Session) {
	res, err := r.DBList().Run(session)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Close()

	var databases []string

	err = res.All(&databases)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !contains(databases, database) {
		_, err = r.DBCreate(database).RunWrite(session)

		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("Created database [%s]", database)
	} else {
		log.Debugf("Database [%s] already exists", database)
	}
}

func createTables(session *r.Session) {
	res, err := r.TableList().Run(session)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Close()

	var tables []string

	err = res.All(&tables)
	if err != nil {
		log.Error(err.Error())
	}

	createTable(UsersTableName, tables, session)
}

func createTable(name string, tabels []string, session *r.Session) {
	if !contains(tabels, name) {
		_, err := r.TableCreate(name).RunWrite(session)
		if err != nil {
			log.Error(err.Error())
		}

		log.Infof("Created table [%s]", name)
	} else {
		log.Debugf("Table [%s] already exists", name)
	}
}

func contains(strings []string, search string) bool {
	for _, a := range strings {
		if a == search {
			return true
		}
	}

	return false
}

func getOneByID(s *r.Session, table string, id string, v interface{}) error {
	res, err := r.Table(table).Get(id).Run(s)
	if err != nil {
		return err
	}
	defer res.Close()

	err = res.One(v)
	if err != nil {
		return err
	}

	return nil
}

func getAll(s *r.Session, table string, v interface{}) error {
	res, err := r.Table(table).Run(s)
	if err != nil {
		return err
	}
	defer res.Close()
	err = res.All(v)
	if err != nil {
		return err
	}
	return nil
}
