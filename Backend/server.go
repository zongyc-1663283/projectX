package main

import (
	"database/sql"
	"log"

	//"fmt"
	// "net/http"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type Info struct{
	Id 			string	`json: id`
	Long		float64	`json: long`
	Lat			float64	`json: lat`
	DestLong	float64	`json: destLong`
	DestLat		float64	`json: destLat`
}

var db *sql.DB
var err error
func init() {
	fmt.Print("Hellodasdsadasdsa")
	db, err := sql.Open("mysql",
	"wilson:admin@tcp(35.239.172.63:3306)/projectx")
	if err != nil {
		log.Fatal("no connection to Db")
	}
	fmt.Print(db)
}

func main() {
	//fmt.Println(Close(47.654425, -122.303852, 47.655502, -122.305086));
	b := Close(47.650375, -122.312309, 47.661050, -122.304864);
	fmt.Print(b);

	e := echo.New();
	e.Use(middleware.Logger());
	e.Use(middleware.Recover());
	e.POST("/request", Request);
	e.POST("/comfirm", Confirm);
	e.Logger.Fatal(e.Start(":1323"));

}

func Request(c echo.Context) error{
	info := new(Info)
	err := c.Bind(info)
	if err != nil {
		log.Println(err)
		return err;
	}
	c.NoContent(200);
	return nil
	//findMatch
}

func Confirm(c echo.Context) error {
	c.NoContent(200);
	return nil;
}