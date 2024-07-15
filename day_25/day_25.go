package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type Vertex struct {
    value string
    edges map[string]*Edge
}

type Edge struct {
    weight int
    vert *Vertex
}

type Graph struct {
    vertices map[string]*Vertex
}

func (this *Graph) AddEdge(srcKey, destKey string, weight int) {
	// check if src & dest exist
	if _, ok := this.vertices[srcKey]; !ok {
		return
	}
	if _, ok := this.vertices[destKey]; !ok {
		return
	}

	// add edge src --> dest
	this.vertices[srcKey].edges[destKey] = &Edge{weight: weight, vert: this.vertices[destKey]}
}

func (this *Graph) AddVertex(key, val string) {
    if _, ok := this.vertices[key]; ok {
        return
    }
	this.vertices[key] = &Vertex{value: val, edges: map[string]*Edge{}}
}

func printVisualized(graph Graph) {
    for k, v := range graph.vertices {
        fmt.Println(k, (*v).edges)
    }
}

func main() {
    file, err := os.Open("day_25")
    if err != nil {
        fmt.Println(err)
    }
    graph := Graph{vertices: map[string]*Vertex{}}
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        splitline := strings.Split(line, ": ")
        right := strings.Split(splitline[1], " ")
        if _, ok := graph.vertices[splitline[0]]; !ok {
            graph.AddVertex(splitline[0], splitline[0])
        }
        for _, el := range right {
            graph.AddVertex(el, el)
            graph.AddEdge(splitline[0], el, 1)     
        }
    }
    printVisualized(graph)
}
