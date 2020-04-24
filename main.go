package main

import (
	"database/sql"
	"log"
	"os"

	casbinpg "github.com/cychiuae/casbin-pg-adapter"
	casbinmw "github.com/labstack/echo-contrib/casbin"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	ce, err := newEnforcer(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(casbinmw.Middleware(ce))

	e.GET("/admin", handlePage("Admin page accessed"))
	e.GET("/login", handlePage("Login accessed"))
	e.GET("/logout", handlePage("Logout accessed"))

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func newEnforcer(db *sql.DB) (*casbin.Enforcer, error) {
	adapter, err := casbinpg.NewAdapter(db, "casbin")
	if err != nil {
		return nil, err
	}

	ce, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		return nil, err
	}

	if err := ce.LoadPolicy(); err != nil {
		return nil, err
	}

	_, _ = ce.AddPolicy("admin", "/admin", "*")
	_, _ = ce.AddPolicy("guest", "/login", "*")
	_, _ = ce.AddPolicy("user", "/logout", "*")

	if err := ce.SavePolicy(); err != nil {
		return nil, err
	}

	return ce, err
}

func handlePage(s string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, s)
	}
}
