package main

import (
	"context"
	"github.com/kom1ssar/tech_em/internal/app"
)

func main() {
	app := app.NewApp(context.Background())

	app.Run()
}
