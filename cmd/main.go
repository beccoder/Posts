package main

import (
	"Blogs/internal"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		internal.HandleCLA(os.Args)
	} else {
		internal.Run()
	}
}
