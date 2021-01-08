package main

import (
	"fmt"
	"os"

	lib "./lib"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		n, g := lib.ParseFile(filename)

		group := g.GetBestPath(n)
		fmt.Println("Paths", lib.Paths, "\n", "len Paths:", len(lib.Paths), "\n")
		fmt.Println("PathGroups", lib.PathGroups, "\n", "len PathGroups:", len(lib.PathGroups), "\n")
		fmt.Println("bestpath", group)
		//g.SendAnts(group, n)
	} else if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments.")
	}
}
