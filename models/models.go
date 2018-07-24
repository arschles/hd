package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	log.Println("connecting to DB environment", env)
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"

	log.Println("creating DB")
	pop.CreateDB(DB)
	migrator, err := pop.NewFileMigrator("./migrations", DB)
	if err != nil {
		log.Fatal("couldn't create migrator:", err)
	}
	log.Println("running migrations")
	if err := migrator.Up(); err != nil {
		log.Fatal("couldn't migrate up:", err)
	}
}
