package main

import (
    "fmt"
    "bufio"
    "os"
    "flag"
    "strings"
    "strconv"
    "math"
)

var map_hex = map[byte]string {
    0 : "R",
    1 : "D",
    2 : "L",
    3 : "U",
}

type position struct {
    x int
    y int
}

var directions = map[string]position {
    "U" : {-1,0},
    "D" : {1,0},
    "L" : {0,-1},
    "R" : {0, 1},
}

func dig(dir string, curr position, grid *[]position, amount int) position {
    new_idx := position{}
    new_idx.x = curr.x + (directions[dir].x * amount)
    new_idx.y = curr.y + (directions[dir].y * amount)
    (*grid) = append((*grid), new_idx)
    return new_idx
}

func full_area(perim int, tr []position) int {
    // pick's theorem + shoelace formula
    // summing up interior points + perimeter gives the answer
    trench_area := area(tr)
    interior_points := (trench_area + 1) - (perim/2)
    return perim + interior_points
}

func area(grid []position) int {
    a := 0
    for i:=0; i<len(grid); i++ {
        if i==len(grid)-1 {
            a += determinant(grid[i], grid[0])
        } else {
            a += determinant(grid[i], grid[i+1])
        }
    }
    return int(math.Abs(float64(a))/2)
}

func determinant (a, b position) int {
    return (a.x * b.y) - (b.x * a.y)
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
        file, err = os.Open("day_18_t")
        if err != nil {
            fmt.Println(err)
        }
    } else {
        file, err = os.Open("day_18_i")
        if err != nil {
            fmt.Println(err)
        }
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    trench := make([]position, 0)
    last_pos := position{0,0}
    new_lp := position{0,0}
    new_trench := make([]position, 0)
    perimeter := 0
    new_perimeter := 0
    for scanner.Scan(){
        line := scanner.Text()
        l := strings.Split(line, " ")
        i, _ := strconv.Atoi(l[1])
        perimeter += i
        last_pos = dig(l[0], last_pos, &trench, i)
        // part 2
        new_str := strings.ReplaceAll(l[2], "(", "")
        new_str = strings.ReplaceAll(new_str, ")", "")
        new_str = strings.ReplaceAll(new_str, "#", "")
        hex, _ := strconv.ParseInt(new_str[:len(new_str)-1], 16, 64)
        new_lp = dig(map_hex[new_str[len(new_str)-1]-'0'], new_lp, &new_trench, int(hex))
        new_perimeter += int(hex)
    }
    fmt.Println(full_area(perimeter, trench))
    fmt.Println(full_area(new_perimeter, new_trench))
}
