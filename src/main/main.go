package main

import (

	"os"
	"strings"
	"server"
)

const DEBUG = true



func main() {
	server.InitConfiguration()
	db := server.ConnectDatabase()
	defer db.Close()
	//
	db.Init()
	server.SystemConfig.InitDbConfig(db.Mysql.Tables)
	//
	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args[1:]
		for _, arg := range args {
			if strings.HasPrefix(arg, "tables=") {
				data := strings.Replace(arg, "tables=", "", 1)
				for _, value := range strings.Split(data, ",") {
					server.ARGS_REQUESTED_TABLES = append(server.ARGS_REQUESTED_TABLES, value)
				}
			}
		}

	} else {
		args = []string{"Initial", "Listener"}
	}

	if server.StringInSlice("Initial", args) {
		if server.StringInSlice("Listener", args) {
			go server.R2G(db)
		} else {
			server.R2G(db)
		}
	}

	if server.StringInSlice("Listener", args) {
		server.RunCommand(db)
	}
}

