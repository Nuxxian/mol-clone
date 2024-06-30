package main

import (
	"nuxxian/mol-clone/server/database"
	"nuxxian/mol-clone/server/routing"
)

func main() {
	database.Database()
	routing.Route()
}

