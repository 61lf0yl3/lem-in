package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ParseFile extracts number of ants, rooms, rooms connections from given file
func ParseFile(filename string) (int, Graph) {
	var g Graph

	data, err := ioutil.ReadFile(filename)
	check(err)
	Task = string(data)

	lines := strings.Split(string(data), "\n")

	n, err1 := strconv.Atoi(lines[0])
	check(err1)
	if n == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	flag := "room"
	lines = lines[1:]

	for _, line := range lines {
		if line == "##start" {
			flag = "start"
			continue
		} else if line == "##end" {
			flag = "end"
			continue
		} else if len(line) == 0 || line[:1] == "#" { // if it is comment
			continue
		}

		splitted := strings.Split(line, " ")

		if len(splitted) == 1 { // if it is a link
			rooms := strings.Split(splitted[0], "-")
			g.Connect(rooms[0], rooms[1])
			continue
		}

		if flag == "start" {
			g.Start = splitted[0]
		} else if flag == "end" {
			g.End = splitted[0]
		}
		flag = "room"
	}

	return n, g
}
