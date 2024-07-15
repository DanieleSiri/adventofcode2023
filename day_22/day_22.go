package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type position struct {
    x uint64
    y uint64
    z uint64
}

type Brick struct {
    start_pos position
    end_pos position
}

func get_int(char string) uint64 {
    new_s, err := strconv.Atoi(char)
    if err != nil {
        panic(err)
    }
    return uint64(new_s)
}

func main() {
    file, err := os.Open("day_22")
    if err != nil {
        fmt.Println(err)
    }
    bricks := make([]Brick, 0)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        split := strings.Split(line, "~")

        s := strings.Split(split[0], ",")
        t := strings.Split(split[1], ",")
        new_start := position{get_int(s[0]), get_int(s[1]), get_int(s[2])}
        new_end := position{get_int(t[0]), get_int(t[1]), get_int(t[2])}
        bricks = append(bricks, Brick{new_start, new_end})
    }
    fmt.Println(bricks)
}
