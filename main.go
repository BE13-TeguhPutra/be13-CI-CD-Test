package main

import (
	"be13/clean-arch/config"
	"be13/clean-arch/factory"
	"be13/clean-arch/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()

	mysql.DBMigration(db)

	factory.InitFactory(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
