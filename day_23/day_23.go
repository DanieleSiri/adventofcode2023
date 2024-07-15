package main

import (
    "fmt"
    "bufio"
    "os"
)

type index struct {
    x int
    y int
}

var directions = map[rune]index {
    'N' : {-1,0},
    'S' : {1,0},
    'E' : {0,1},
    'W' : {0,-1},
}

var sign_direction = map[rune]rune {
    '>' : 'E',
    '<' : 'W',
    'v' : 'S',
}

func walk(gr map[index]rune, curr, end, prev index, seen map[index]int) {
    seen[curr] = 1
    if curr == end {
        return
    }
    next_idx := index{}
    if v, ok := sign_direction[gr[curr]]; ok {
        next_idx.x = curr.x + directions[v].x
        next_idx.y = curr.y + directions[v].y
    } else {
        for _, v := range directions {
            next_idx.x = curr.x + v.x
            next_idx.y = curr.y + v.y
            // check if index is in the grid and if we have not seen it or we came from it
            if _, ok := seen[next_idx]; ok || next_idx == prev {
                continue
            }
            if _, ok := gr[next_idx]; ok {
                break
            }
        }
    }
    walk(gr, next_idx, end, curr, seen)
}

func main() {
    file, err := os.Open("day_23")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    row := 0
    grid := make(map[index]rune)
    for scanner.Scan(){
        line := scanner.Text()
        for k, c := range line {
            if c != '#' {
                idx := index{}
                idx.x = row
                idx.y = k
                grid[idx] = c
            }
        }
        row++
    }
    path_list := make([]map[index]int, 0)
    path := map[index]int {}
    for {
        walk(grid, index{0,1}, index{22, 21}, index{-1,-1}, path)
        path_list = append(path_list, path)
        break
    }
    fmt.Println(len(path))
}
