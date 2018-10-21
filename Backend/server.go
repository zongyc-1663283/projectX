package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

type Info struct{
	Id 			string	`json: id`
	Long		float64	`json: long`
	Lat			float64	`json: lat`
	DestLong	float64	`json: destLong`
	DestLat		float64	`json: destLat`
	Phone		string	`json: phone`
}

type Partners struct{
	InfoList	[]Info `json: infoList`
}

var db *sql.DB
var err error
func init() {
	db, err = sql.Open("mysql",
	"wilson:admin@tcp(35.239.172.63:3306)/projectx")
	if err != nil {
		log.Fatal("no connection to Db")
	}
}



func main() {
	//fmt.Println(Close(47.654425, -122.303852, 47.655502, -122.305086));
	//b := Close(47.650375, -122.312309, 47.661050, -122.304864);
	//fmt.Print(b);
	e := echo.New();
	e.Use(middleware.Logger());
	e.Use(middleware.Recover());
	e.DELETE("/delete", Confirm);
	e.POST("/post", Post)
	e.Logger.Fatal(e.Start(":80"));


}


func Post(c echo.Context) error {
	info := new(Info)
	err := c.Bind(info)
	if err != nil {
		log.Println(err)
		return err;
	}
	info.Id = info.Id + time.Stamp // very wrong
	stmt, err := db.Prepare("INSERT INTO UserStatus(ID, CurLat, CurLong, DestLat, DestLong, Phone) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(info.Id, info.Lat, info.Long, info.DestLat, info.DestLong, info.Phone);
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT ID, CurLat, CurLong, DestLat, DestLong, Phone FROM UserStatus");
	if err != nil {
		log.Println(err);
		return err;
	}
	defer rows.Close();
	partners := []Info{};
	for rows.Next() {
		var id string
		var curLat2 float64
		var curLong2 float64
		var destLat2 float64
		var destLong2 float64
		var phone string
		//partnerInfo := new(Info)
		err = rows.Scan(&id, &curLong2, &curLat2, &destLat2, &destLong2, &phone);
		if(Close(info.DestLong, info.DestLat, destLong2, destLat2) &&
			Close(info.Long, info.Lat, curLong2, curLat2)) {
			partners = append(partners, Info{Id: id, Long: curLong2, Lat: curLat2, DestLat:destLat2, DestLong:destLong2, Phone:phone})
		}
	}
	c.JSON(http.StatusOK, Partners{InfoList:partners});
	return nil;
}

func Confirm(c echo.Context) error {
	info := new(Info)
	err := c.Bind(info)
	if err != nil {
		log.Println(err)
		return err;
	}
	_, err = db.Exec("DELETE FROM users WHERE Id = ?", info.Id)
	if err != nil {
		return err
	}
	c.String(200, "Record with id " + info.Id + "deleted");
	return nil;
}