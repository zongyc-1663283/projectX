package main

import (
	//"fmt"
	// "net/http"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	//fmt.Println(Close(47.654425, -122.303852, 47.655502, -122.305086));
	b := Close(47.650375, -122.312309, 47.661050, -122.304864);
	fmt.Print(b);
	e := echo.New();
	e.Logger.Fatal(e.Start(":1323"));

}