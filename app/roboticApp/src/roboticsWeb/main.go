package main

const (
	PORT = "8080"
)

func main() {
	r := registerRoutes()
	r.Run(":" + PORT)
}
