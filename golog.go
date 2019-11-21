package gotest

import "log"

func init()  {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("Message")
	log.Fatalln("Fatal Message")
	log.Panicln("Panic Message")
}
