package server

func StartServer() {
	app := NewApp()
	app.Run(":8080")
}
