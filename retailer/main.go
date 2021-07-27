package main

import (
	"freshers-bootcamp/retailer/routes"
)

func main() {
	DBname := "first_go"
	r := routes.Setup(DBname)
	r.Run()
}
