package main

import (
    "fmt"
    "bufio"
    "os"
)

// changed based on input for testing and real problem
var max_steps = 64

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

func update_grid(gr *[][]rune, l []index) {
    new_gr := gr
    for _, el := range l {
        for _, v := range directions {
            new_x := el.x + v.x
            new_y := el.y + v.y
            if new_x < 0 || new_y < 0 || new_x >= len((*gr)) || new_y >= len((*gr)[0]) {
                continue
            }
            if (*gr)[new_x][new_y] == '#' {
                continue
            } else {
                (*new_gr)[el.x][el.y] = '.'
                (*new_gr)[new_x][new_y] = 'O'
            }
        }
    }
    gr = new_gr
}

func walk (gr *[][]rune, start index, steps_left int) {
    if steps_left == 0 {
        return
    }
    // first step
    indexes := make([]index, 0)
    if steps_left == max_steps {
        indexes = append(indexes, start)
    } else {
        for i:=0; i<len((*gr)); i++ {
            for j:=0; j<len((*gr)[0]); j++ {
                if (*gr)[i][j] == 'O' {
                    indexes = append(indexes, index{i,j})
                }
            }
        }
    }
    update_grid(gr, indexes)
    walk(gr, start, steps_left - 1)
}

func print_visualized(grid [][]rune) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            fmt.Print(string(grid[i][j])) 
        }
        fmt.Print("\n")
    }
}

func get_plots(grid [][]rune) (ret int) {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == 'O' {
                ret++
            }
        }
    }
    return
}

func main() {
    file, err := os.Open("day_21")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    grid := make([][]rune, 0)
    start_idx := index{}
    row := 0
    for scanner.Scan(){
        line := scanner.Text()
        tmp := make([]rune, 0)
        for k, c := range line {
            if c=='S' {
                start_idx.x = row
                start_idx.y = k
            }
            tmp = append(tmp, c)
        }
        grid = append(grid, tmp)
        row++
    }
    walk(&grid, start_idx, max_steps)
//    print_visualized(grid)
    fmt.Println(get_plots(grid))
}
