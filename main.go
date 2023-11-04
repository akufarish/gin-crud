package main

import (
	"github.com/akufarish/gin-crud/db"
	"github.com/akufarish/gin-crud/services"
)

func main()  {
	db.Init()
	services.Init()
}