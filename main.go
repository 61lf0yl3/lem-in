package main

import (
	"fmt"
	"os"

	tools "./tools"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("too many arguments or not enough")
		os.Exit(0)
	}
	filename := os.Args[1]
	g := tools.ParseFile(filename)
	g.GetAllPaths(g.Start)
	tools.SortPathsGroups1()
	//fmt.Println("Paths:", tools.Paths, "\n", len(tools.Paths), "\n")
	//fmt.Println()
	g.GetDisjointPaths()
	//fmt.Println("PathsGroups :", tools.PathsGroups, "\n", len(tools.PathsGroups), "\n")
	//fmt.Println()
	optimal := g.OptimalPath()
	fmt.Println("optimalpath", optimal)
	g.SendAnts(optimal)
}
