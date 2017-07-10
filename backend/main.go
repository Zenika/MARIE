package main

import "github.com/Zenika/MARIE/backend/app"

func main() {
	a := new(app.App)
	a.Initialize()
	a.Run(":8081")
}
