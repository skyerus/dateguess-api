package main

import (
	"github.com/skyerus/history-api/pkg/api"
	"github.com/skyerus/history-api/pkg/env"
)

func main() {
	env.SetEnv()
	main := &api.App{}
	db, err := api.OpenDb()
	if err != nil {
		return
	}
	defer db.Close()
	main.Initialize(db)
	main.Run(":80")
}
