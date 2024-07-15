package main

import (
    "fmt"
    "os"
    "bufio"
)

type index struct {
    x int
    y int
}

func roll (grid *[][]rune, col int, curr int) {
    if curr < 0 {
        // out of bounds
        return
    }
    if (*grid)[curr][col] == '#' || (*grid)[curr][col] == 'O' {
        return
    }
    if curr  != len((*grid))-1 {
        (*grid)[curr][col] = 'O'
        (*grid)[curr+1][col] = '.'
    }
    roll(grid, col, curr-1)
}

func get_stones (grid [][]rune) (ret int) {
    value := len(grid)
    for i:=0; i<len(grid[0]); i++ {
        for j:=0; j<len(grid); j++ {
            if grid[i][j] == 'O' {
                ret += value
            }
        }
        value--
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

//part 2
func stringify (grid [][]rune) (ret string){
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            ret += string(grid[i][j])
        }
    }
    return
}

func get_value (m map[string]int, l int, cycle int, size int) (ret int) {
    grid := make([][]rune, size)
    row := 0
    col := 0
    for k, v := range m {
        if v == cycle {
            for _, val := range k {
                grid[row] = append(grid[row], val)
                col++
                if col == l {
                    row++
                    col = 0 
                }
            }
            break
        }
    }
    ret = get_stones(grid)
    return
}

func roll_2 (grid *[][]rune, idx index, dir rune) {
    if idx.x < 0 || idx.y < 0  || idx.y >= len((*grid)) || idx.x >= len((*grid)[0]){
        // out of bounds
        return
    }
    if (*grid)[idx.x][idx.y] == '#' || (*grid)[idx.x][idx.y] == 'O' {
        return
    }
    switch dir {
    case 'N':
        if idx.x  != len((*grid))-1 {
            (*grid)[idx.x][idx.y] = 'O'
            (*grid)[idx.x+1][idx.y] = '.'
        }
        roll_2(grid, index{idx.x-1, idx.y}, 'N')
    case 'S':
        if idx.x  != 0 {
            (*grid)[idx.x][idx.y] = 'O'
            (*grid)[idx.x-1][idx.y] = '.'
        }
        roll_2(grid, index{idx.x+1, idx.y}, 'S')
    case 'W':
        if idx.y != len((*grid)[0])-1 {
            (*grid)[idx.x][idx.y] = 'O'
            (*grid)[idx.x][idx.y+1] = '.'
        }
        roll_2(grid, index{idx.x, idx.y-1}, 'W')
    case 'E':
        if idx.y != 0 {
            (*grid)[idx.x][idx.y] = 'O'
            (*grid)[idx.x][idx.y-1] = '.'
        }
        roll_2(grid, index{idx.x, idx.y+1}, 'E')
    }
}

func main() {
    file, err := os.Open("day_14")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    input_grid := make([][]rune, 0)
    input_grid_2 := make([][]rune, 0)
    row := 0
    for scanner.Scan(){
        input_grid = append(input_grid,make([]rune, 0))
        input_grid_2 = append(input_grid_2,make([]rune, 0))
        line := scanner.Text()
        for _, c := range line {
            input_grid[row] = append(input_grid[row], c)
            input_grid_2[row] = append(input_grid_2[row], c)
        }
        row++
    }
    row_len := len(input_grid[0])
    for i:=0; i<row_len; i++ {
        for j:=1; j<len(input_grid); j++ {
            if input_grid[j][i] == 'O' {
                roll(&input_grid, i, j-1) 
            }
        }
    }
//    print_visualized(input_grid)
    fmt.Println(get_stones(input_grid))

    // part 2
    rotation := [4]rune{'N','W','S','E'}
    cycle_found := make(map[string]int)
    stones := ""
    last_cycle := 0
    for n:=1; n<=100000; n++ {
        for _, r := range rotation {
            switch r {
            case 'N':
                for i:=0; i<row_len; i++ {
                    for j:=1; j<len(input_grid_2); j++ {
                        if input_grid_2[j][i] == 'O' {
                            roll_2(&input_grid_2, index{j-1, i}, 'N') 
                        }
                    }
                }
            case 'S':
                for i:=0; i<row_len; i++ {
                    for j:=len(input_grid_2)-1; j>=0; j-- {
                        if input_grid_2[j][i] == 'O' {
                            roll_2(&input_grid_2, index{j+1, i}, 'S') 
                        }
                    }
                }
            case 'W':
                for i:=0; i<len(input_grid_2); i++ {
                    for j:=1; j<row_len; j++ {
                        if input_grid_2[i][j] == 'O' {
                            roll_2(&input_grid_2, index{i, j-1}, 'W') 
                        }
                    }
                }
            case 'E':
                for i:=0; i<len(input_grid_2); i++ {
                    for j:=row_len-1; j>=0; j-- {
                        if input_grid_2[i][j] == 'O' {
                            roll_2(&input_grid_2, index{i, j+1}, 'E') 
                        }
                    }
                }
            }
        }
        stones = ""
        stones = stringify(input_grid_2)
        if _, ok := cycle_found[stones]; ok {
            last_cycle = n
            break
        } else {
            cycle_found[stones] = n
        }
    }
    // print_visualized(input_grid_2)
    res := ((1000000000 - (cycle_found[stones]))%(last_cycle - cycle_found[stones]))+cycle_found[stones]
    fmt.Println(get_value(cycle_found, row_len, res, len(input_grid_2)))
}
