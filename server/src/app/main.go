package main

func main() {
	application := NewApplication()
	application.configureRoutes()
	application.startWebApp()
}
