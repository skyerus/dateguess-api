package main

import (
	"github.com/skyerus/ims-api/pkg/api"
)

func main() {
	main := &api.App{}
	db, err := api.OpenDb()
	if err != nil {
		return
	}
	defer db.Close()
	main.Initialize(db)
	main.Run(":80")
}
