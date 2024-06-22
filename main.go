package main

import (
	"example.com/awesome/database"
	"example.com/awesome/database/migration"
	"example.com/awesome/route"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
)

func main() {
	_ = database.ConnectDb()
	m := gormigrate.New(database.DbManager, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migration.AddUserTable(),
	})

	action := flag.String("action", "", "Action to perform: migrate or rollback")
	flag.Parse()

	switch *action {
	case "migrate":
		if err := m.Migrate(); err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}
		log.Println("Migration did run successfully")

	case "rollback":
		if err := m.RollbackLast(); err != nil {
			log.Fatalf("Could not rollback: %v", err)
		}
		log.Println("Rollback did run successfully")

	default:
		r := gin.Default()

		route.UserRoutes(r)

		err := r.Run(":8080")
		if err != nil {
			log.Fatal("error starting server", err)
		}
	}

}
