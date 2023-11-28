package main

import (
	"Blogs"
	"Blogs/internal"
	"fmt"
	"log"
	"os"
)

var (
	cfg *Blogs.Config
)

func init() {
	log.Println("Initializing...")
	cfg = Blogs.Load()
	log.Println("Initializing done...")
	fmt.Println(cfg)
}

func main() {
	if len(os.Args) > 1 {
		internal.HandleCLA(os.Args, cfg)
	} else {
		internal.Run(cfg)
	}
	fmt.Println("OK")
}
