package main

func main() {
	auth := &AuthHandler{}
	log := &LogHandler{}
	data := &DataHandler{}

	auth.SetNext(log)
	log.SetNext(data)

	auth.Handle("auth")
}
