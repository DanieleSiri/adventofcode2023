package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

type position struct {
    x float64
    y float64
}

func find_empty(a [][]rune, e_r map[float64]int, e_v map[float64]int) {
    for row:=0; row<len(a); row++ {
        if is_empty(a[row]) {
            e_r[float64(row)] = 1
        }
    }
    for i:=0; i<len(a); i++ {
        vert_slice := make([]rune, 0)
        for k:=0; k<len(a); k++ {
            vert_slice = append(vert_slice, a[k][i])
        }
        if is_empty(vert_slice) {
            e_v[float64(i)] = 1
        }
    }
}

func is_empty(a []rune) bool {
    for _, c := range a {
        if c == '#' {
            return false
        }
    }
    return true
}

func shortest_path(expansion_coeff int, from position, to position, e_r map[float64]int, e_v map[float64]int) (total int) { 
    // rows
    for i:=math.Min(from.x, to.x); i<math.Max(from.x, to.x); i++ {
        if _, ok := e_r[i]; ok {
            total+=expansion_coeff
        } else {
            total++
        }
    }
    // columns
    for i:=math.Min(from.y, to.y); i<math.Max(from.y, to.y); i++ {
        if _, ok := e_v[i]; ok {
            total+=expansion_coeff
        } else {
            total++
        }
    }
    return
}

func main() {
    grid := make([][]rune, 0)
    galaxies := make(map[int]position)
    file, err := os.Open("day_11")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    line_number := 0
    galaxy_count := 1
    for scanner.Scan(){
        line := scanner.Text()
        tmp := make([]rune, 0)
        for i, c := range line {
            tmp = append(tmp, c)
            if c == '#' {
                galaxies[galaxy_count] = position{float64(line_number), float64(i)}
                galaxy_count++
            }
        }
        grid = append(grid, tmp)
        line_number++
    }
    // maps of empty rows and cols
    empty_rows := make(map[float64]int)
    empty_cols := make(map[float64]int)
    find_empty(grid, empty_rows, empty_cols)
    sum := 0
    // used to exclude 1 of the 2 pairs (need to consider only 1) 
    galaxies_done := make(map[position]int)
    for k, v := range galaxies {
        for l, s := range galaxies {
            curr := position{float64(l), float64(k)}
            if _, ok := galaxies_done[curr]; !ok && l != k {
                // part 2 = 1000000
                // part 1 = 2
                sum+=shortest_path(1000000, v, s, empty_rows, empty_cols)
                couple := position{float64(k), float64(l)}
                galaxies_done[couple] = 1
            }
        }
    }
    fmt.Println(sum)
}
