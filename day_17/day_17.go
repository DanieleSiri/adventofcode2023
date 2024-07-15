package main

import (
    "fmt"
    "bufio"
    "os"
    "flag"
    "sort"
)

// implementing a BFS
type Node struct {
    key position
    neighbors map[position]int
}

type position struct {
    x int
    y int
}

func BFS (grid map[position]Node, start position, consecutive int) map[position]int {
    visited := make(map[position]bool, 0)
    visited[start] = true
    distances := make(map[position]int)
    for key := range grid {
        distances[key] = 10000
    }
    distances[start] = 0

    var vertices []Node 
    for _, vert := range grid {
        vertices = append(vertices, vert)
    }

    for len(vertices) != 0 {
        sort.SliceStable(vertices, func(i, j int)bool {
            return distances[vertices[i].key] < distances[vertices[j].key]
        })
        vertex := vertices[0]
        vertices = vertices[1:]
        visited[vertex.key] = true
        
        for adjacent, cost := range vertex.neighbors {
            if _, ok := visited[adjacent]; ok {
                continue
            }
            alt := distances[vertex.key] + cost
            if vertex.key.x == 0 && vertex.key.y == 2 {
                fmt.Println(vertex.neighbors)
                fmt.Println(distances[vertex.key],adjacent, distances[adjacent], alt)
            }
            if alt < distances[adjacent] {
                distances[adjacent] = alt
            }
        }
    }
    return distances
}

var directions = map[rune]position {
    'N' : {-1,0},
    'S' : {1,0},
    'W' : {0,-1},
    'E' : {0, 1},
}

func get_neighbors(p position, l int, val [][]int, neighb map[position]int) {
    for _, dir := range directions {
        new_pos := position{p.x + dir.x, p.y + dir.y}
        if new_pos.x>=0 && new_pos.x<l && new_pos.y>=0 && new_pos.y<l {
           neighb[new_pos] = val[new_pos.x][new_pos.y]
        }
    }
}

func build_graph(g [][]int, m map[position]Node) {
    for i:=0; i<len(g); i++ {
        for j:=0; j<len(g[i]); j++ {
            idx := position{i, j}
            node := Node{}
            node.key = idx
            node.neighbors = make(map[position]int)
            get_neighbors(idx, len(g[i]), g, node.neighbors)
            m[idx] = node
        }
    }
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("no input file provided")
	}
    var file *os.File
    var err error

    if args[0] == "t" {
        file, err = os.Open("day_17_t")
        if err != nil {
            fmt.Println(err)
        }
    } else {
        file, err = os.Open("day_17_i")
        if err != nil {
            fmt.Println(err)
        }
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    grid := make([][]int, 0)
    graph := make(map[position]Node)
    row := 0
    for scanner.Scan(){
        line := scanner.Text()
        tmp := make([]int, 0)
        for _, c := range line {
            tmp = append(tmp, int(c-'0'))
        }
        grid = append(grid, tmp)
        row++
    }
    build_graph(grid, graph) 
    fmt.Println(BFS(graph, position{0,0}, 3))
}
