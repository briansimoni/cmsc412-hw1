package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// the main function will create a graph object
// read the graph.txt file
// and then output the unnormalized weighted degree centrality
// to wdegree.txt
func main() {

	file, err := os.Open("graph.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	// Scan the first line and get the number of nodes
	scanner.Scan()
	n := firstLine(scanner.Text())

	// Create a new graph based on the number of nodes
	g := NewGraph(n)

	for scanner.Scan() {
		e := getEdge(scanner.Text())
		g.nodes[e.fromNode-1].AddEdge(e)

		// its undirected, so I'll just switch the to and from node and add another edge
		e1 := undirectedEdge(e)
		g.nodes[e1.fromNode-1].AddEdge(e1)
	}

	fileString := ""
	for i := range g.nodes {
		fileString += floatToString(g.nodes[i].WeightedDegreeCentrality()) + "\n"
	}

	f, err := os.Create("wdegree.txt")
	check(err)
	f.Write([]byte(fileString))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// parse the line of text of the file, and extract an edge from a line
func getEdge(line string) edge {

	e := strings.Split(line, " ")
	fromNode, _ := strconv.Atoi(e[0])
	toNode, _ := strconv.Atoi(e[1])
	weight, err := strconv.ParseFloat(e[2], 64)
	check(err)

	edge := edge{fromNode, toNode, weight}
	return edge

}

// parse the first line and return the number of nodes in the graph
func firstLine(line string) int {
	l := strings.Split(line, " ")
	n, _ := strconv.Atoi(l[0])
	return n
}

// flips the to and from node
func undirectedEdge(e edge) edge {
	fromNode := e.toNode
	toNode := e.fromNode
	return edge{fromNode, toNode, e.Weight}
}

// convert the float value to string and truncate 0s
func floatToString(input_num float64) string {
	// to convert a float number to a string
	n := strconv.FormatFloat(input_num, 'f', 6, 64)
	s := ""
	r := true
	for i := len(n) - 1; i >= 0; i-- {
		if string(n[i]) == "0" && r {
			continue
		} else if string(n[i]) == "." {
			r = false
			continue
		} else {
			s = string(n[i]) + s
		}
	}
	return s
}
