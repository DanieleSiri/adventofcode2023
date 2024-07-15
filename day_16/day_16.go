package main

import (
    "fmt"
    "bufio"
    "os"
)

var opposites = map[rune]rune {
    'N': 'S',
    'S': 'N',
    'W': 'E',
    'E': 'W',
}

type position struct {
    x int
    y int
}

var directions = map[rune]position {
    'N' : {-1,0},
    'S' : {1,0},
    'W' : {0,-1},
    'E' : {0, 1},
}

func get_90_deg (from rune, turn rune) (to rune) {
    if turn == '\\' {
        switch from {
        case 'N':
            to = 'E'
        case 'S':
            to = 'W'
        case 'E':
            to = 'N'
        case 'W':
            to = 'S'
        }
    } else {
        switch from {
        case 'N':
            to = 'W'
        case 'S':
            to = 'E'
        case 'E':
            to = 'S'
        case 'W':
            to = 'N'
        }
    }
    return
}

func print_visualized(grid [][]rune) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            fmt.Print(string(grid[i][j])) 
        }
        fmt.Print("\n")
    }
}

func walk(grid [][]rune, curr position, from rune, to rune, seen map[position]int, visited map[position]rune) {
    if curr.x < 0 || curr.y < 0 || curr.x >= len(grid) || curr.y >= len(grid[0]){
        return
    }
    // avoiding closed loops
    if v, ok := visited[curr]; ok && v == from {
        return
    }
    c := grid[curr.x][curr.y]
    seen[curr] = 1
    visited[curr] = from
    switch c {
    case '\\','/':
        to = get_90_deg(from, c)
    case '|':
        if from != 'N' && from != 'S' {
            // split
            to_l := get_90_deg (from, '/')
            to = get_90_deg (from, '\\')
            next_l := position {curr.x + directions[to_l].x, curr.y + directions[to_l].y}
            walk(grid, next_l, opposites[to_l], to_l, seen, visited)
        }
    case '-':
        if from != 'W' && from != 'E' {
            //split
            to_l := get_90_deg (from, '/')
            to = get_90_deg (from, '\\')
            next_l := position {curr.x + directions[to_l].x, curr.y + directions[to_l].y}
            walk(grid, next_l, opposites[to_l], to_l, seen, visited)
        }
    }
    next := position {curr.x + directions[to].x, curr.y + directions[to].y}
    walk(grid, next, opposites[to], to, seen, visited)
}

func check_len(grid [][]rune, curr position, from rune, to rune, m int) (max int) {
    r := make(map[position]int)
    v := make(map[position]rune)
    walk(grid, curr, from, to , r, v)  
    ret := len(r)
    if ret > m {
        max = ret
    } else {
        max = m
    }
    return
}

func main() {
    file, err := os.Open("day_16")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    input_grid := make([][]rune, 0)
    for scanner.Scan(){
        line := scanner.Text()
        tmp := make([]rune, 0)
        for _, c := range line {
            tmp = append(tmp, c)
        }
        input_grid = append(input_grid, tmp)
    }
    rays := make(map[position]int)
    vis := make(map[position]rune)
    walk(input_grid, position{0,0}, 'W', 'E', rays, vis)
    fmt.Println(len(rays))
    // part 2
    max_len := 0
    for i:=0; i<len(input_grid); i++ {
        for j:=0; j<len(input_grid[0]); j++ {
            if i==0 {
                max_len = check_len(input_grid, position{i,j}, 'N', 'S', max_len)
            }
            if j==0 {
                max_len = check_len(input_grid, position{i,j}, 'W', 'E', max_len)
            }
            if j==len(input_grid)-1 {
                max_len = check_len(input_grid, position{i,j}, 'E', 'W', max_len)
            }
            if i==len(input_grid[0])-1 {
                max_len = check_len(input_grid, position{i,j}, 'S', 'N', max_len)
            }
        }
    }
    fmt.Println(max_len)
}
