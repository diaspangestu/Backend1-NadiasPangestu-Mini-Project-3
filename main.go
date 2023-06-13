package main

import (
	"fmt"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/db"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/admin"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/customer"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/modules/superadmin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"path"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.New()

	// open connection db
	db := db.InitDB()

	migrationSource := &migrate.FileMigrationSource{
		Dir: path.Join("."),
	}
	sqlDb, _ := db.DB()
	n, err := migrate.Exec(sqlDb, "mysql", migrationSource, migrate.Up)
	fmt.Println(n, err)

	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	// Admin Handler
	adminHandler := admin.NewRouter(db)
	adminHandler.Handle(router)

	// Superadmin Handler
	superadminHandler := superadmin.NewRouter(db)
	superadminHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
