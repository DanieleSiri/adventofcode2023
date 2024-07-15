package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

type position struct {
    x int
    y int
}

var directions = map[rune]position {
    'N':{0,-1},
    'E':{1,0},
    'S':{0,1},
    'W':{-1,0},
}

var symbol_map = map[rune][2]rune {
    '|' : {'N','S'},
    '-' : {'E','W'},
    'L' : {'N','E'},
    'J' : {'N','W'},
    '7' : {'S','W'},
    'F' : {'S','E'},
}

type pipe struct {
    symbol rune
    directions map[rune]int
}

func get_opposite(a rune) rune {
    switch a {
    case 'W':
        return 'E'
    case 'E':
        return 'W'
    case 'N':
        return 'S'
    case 'S':
        return 'N'
    }
    return '\n'
}

func walk (p_list map[position]pipe, curr position, from rune, path *[]position) int{
    // next value does not exist (found a .)
    if _, ok := p_list[curr]; !ok {
        return -1
    }
    // we finished the loop (came back to S)
    if v := p_list[curr]; v.symbol == 'S' {
       return 0
    }
    // value exists but the pipe is of the wrong type and does not support the direction we came from
    v := p_list[curr]; 
    if _, ok := v.directions[from]; !ok {
        return -1
    }
    var to rune
    // direction we're going
    for k := range v.directions {
        if k != from {
            to = k
        }
    }
    next_pos := position{curr.x + directions[to].x, curr.y + directions[to].y}
    // for part 2
    *path = append(*path, curr)
    return walk(p_list, next_pos, get_opposite(to), path) + 1
}

// for part 2
func determinant (x, y position) int {
    return (x.x * y.y) - (y.x * x.y)
}

func main() {
    file, err := os.Open("day_10")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    line_number := 0
    pipe_list := make(map[position]pipe)
    starting_pos := position{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        for i, c := range line {
            p := pipe{}
            if _, ok := symbol_map[c]; !ok && c!= 'S'{
                // . 
                continue
            }
            dir := make(map[rune]int)
            for _, v := range symbol_map[c] {
                dir[v] = 1
            }
            if c == 'S' {
                starting_pos.x = i
                starting_pos.y = line_number
            }
            p.directions = dir
            p.symbol = c
            pipe_list[position{i, line_number}] = p
        }
        line_number++
    }
    // for part 2
    seen := []position{}
    // number of steps to the farthest 
    // we could implement a check for every direction walked to use the first step of recursion but we'll skip for now and look at our starting point in the file to determine where to go
    n := math.Ceil(float64(walk(pipe_list, position{starting_pos.x+1, starting_pos.y}, 'W', &seen))/2)
    fmt.Println(n)
    
    // create a map of our pipeline so that access is quicker
    seen_map := make(map[position]int)
    // add starting position to map
    seen_map[starting_pos] = 1
    for _, v := range seen {
        seen_map[v] = 1
    }
    tmp_dets := determinant(seen[len(seen)-1], seen[0])
    for i:=1; i<len(seen); i++ {
        tmp_dets += determinant(seen[i-1], seen[i])
    }
    // pick's theorem
    tmp_dets = int(math.Abs(float64(tmp_dets)))
    l := tmp_dets/2 + 1 - ((int(n)*2)/2)
    fmt.Println(l)
}
