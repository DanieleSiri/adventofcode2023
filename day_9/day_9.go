package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func all_zeros(a []int) bool {
    for _, v := range a {
        if v != 0 {
            return false
        }
    }
    return true
}

func calculate (h []int) (r int) {
    if len(h) == 1 || all_zeros(h) {
        return r 
    }
    var curr_res []int
    for i:=1; i<len(h); i++ {
        curr_res = append(curr_res, h[i] - h[i-1])
    }
    return h[len(h)-1] + calculate(curr_res)
}

// part 2
func calculate_2 (h []int) (r int) {
    if len(h) == 1 || all_zeros(h) {
        return r 
    }
    var curr_res []int
    for i:=1; i<len(h); i++ {
        curr_res = append(curr_res, h[i] - h[i-1])
    }
    return h[0] - calculate_2(curr_res)
}

func main() {
    file, err := os.Open("day_9")
    if err != nil {
        fmt.Println(err)
    }
    sum := 0
    sum_2 := 0
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        var history []int
        line := scanner.Text()
        for _, v := range strings.Split(line, " ") {
            n, _ := strconv.Atoi(v)
            history = append(history, n)
        }
        sum += calculate(history)
        sum_2 += calculate_2(history)
    }
    fmt.Println(sum)
    fmt.Println(sum_2)
}
