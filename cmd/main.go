package main

import "dam4rus/nifi-tui/internal/pkg/app"

func main() {
	app := app.NewApp("https://localhost:8443/nifi-api", "aa518870-4cc8-481e-8f69-95aa12a9c4dd", "MUAQnbOWaJ0BhAIhIf+jadsllNPm4tPr")
	app.EnterProcessGroup("root")
	if err := app.Run(); err != nil {
		panic(err)
	}
}
