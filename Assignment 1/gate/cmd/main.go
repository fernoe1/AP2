package main

import (
	"github.com/fernoe1/AP2/assignment-1/gate/internal/app"
)

func main() {
	a := app.InitApp(":8080")
	a.Start()
}
