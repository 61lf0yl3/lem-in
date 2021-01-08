package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// CheckErr check errors
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NbAnts is the number of ants
var NbAnts int

// ParseFile parse file os.Args[1]
func ParseFile(filename string) Lemin {
	var g Lemin
	data, err := ioutil.ReadFile(filename)
	CheckErr(err)
	lines := strings.Split(string(data), "\n")
	var err1 error
	NbAnts, err1 = strconv.Atoi(lines[0])
	CheckErr(err1)
	lines = lines[1:]
	flag := "room"
	for _, line := range lines {
		if line == "##start" {
			flag = "start"
			continue
		} else if line == "##end" {
			flag = "end"
			continue
		} else if len(line) == 0 || line[0] == '#' { // if it comment or empty line
			continue
		}
		splitted := strings.Split(line, " ") //split our line to remove space and, splice from 2

		if len(splitted) == 1 { //there are no space in the, which means it's link between rooms
			linkrooms := strings.Split(splitted[0], "-")
			// connect one room to another and vice versa
			g.Connect(linkrooms[0], linkrooms[1])
			continue
		}
		if flag == "start" {
			// it's start room
			g.Start = splitted[0]
			flag = "room"
		}
		if flag == "end" {
			//it's end room
			g.End = splitted[0]
			flag = "room"
		}
	}
	return g
}
