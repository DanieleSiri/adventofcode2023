package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

type Index struct {
    x int
    y int
}

type Number struct {
    head_idx Index
    value int
    tail_idx Index
}

func isNumber(char rune) bool {
    if 0 <= (char - '0') && (char - '0') <= 9 {
        return true
    }
    return false
}

func isDot(char rune) bool {
    if string(char) == "." {
        return true
    }
    return false
}

func calculate(n int, h Index, t Index, m map[Index]string) int {
    res := 0
    // scan all the neighbours
    for i:=h.y-1; i<=(h.y+1); i++ {
        // already found the value
        if res != 0 {
            break
        }

        // scan from head-1 to tail+1
        for j:=h.x-1; j<=(t.x+1); j++ {
            // in case we are on the row of the number we scan only the edges
            if i == h.y && (j >= h.x && j <= t.x){
                continue
            }
            if _, ok := m[Index{i,j}]; ok {
                res = n 
                break
            }
        }
    }
    return res
}

// part 2
func calculate_2(n int, h Index, t Index, m map[Index]string, res map[Index]Index) {
    result := 0
    // scan all the neighbours
    for i:=h.y-1; i<=(h.y+1); i++ {
        // found the result
        if result != 0 {
            break
        }

        // scan from head-1 to tail+1
        for j:=h.x-1; j<=(t.x+1); j++ {
            // in case we are on the row of the number we scan only the edges
            if i == h.y && (j >= h.x && j <= t.x){
                continue
            }
            if s, ok := m[Index{i,j}]; ok && s=="*"{
                result = n 
                // append to the map
                if idx, found := res[Index{i,j}]; found {
                    // extra check if we have never been here before
                    if idx.y == 0 {
                        res[Index{i,j}] = Index{idx.x, n}
                    }
                } else {
                    // first time we found the symbol
                    res[Index{i,j}] = Index{n,0}
                }
                break
            }
        }
    }
}

func main(){
    file, err := os.Open("day_3")
    if err != nil {
        fmt.Println(err)
    }
    sum := 0
    // sum_2 := 0

    symbols := make(map[Index]string)
    var numbers []Number

    defer file.Close()
    scanner := bufio.NewScanner(file)
    line_number := -1
    for scanner.Scan() {
        number_composed := ""
        line_number++
        var head Index
        var tail Index
        line := scanner.Text()
        for i, char := range line {
            if isNumber(char) {
                // found the first occurrence of a number
                if number_composed == ""{
                    // create the head index
                    head.x = i 
                    head.y = line_number
                }
                // start composing the number
                number_composed += string(char)

                // reached the first char that is not a number or the end of line
                if (i+1) >= len(line) || !isNumber(rune(line[i+1])) {
                    // create the tail index
                    tail.x = i
                    tail.y = line_number
                    // end the number calculation
                    if number_composed != "" {
                        num_value, _ := strconv.Atoi(number_composed)
                        numbers = append(numbers, Number {head, num_value, tail})
                    }
                    number_composed = ""
                }
            } else if !isDot(char){
                // create symbol map
                position := Index{line_number, i}
                symbols[position] = string(char)
            }
        }
    }

    for _, v := range numbers {
        sum += calculate(v.value, v.head_idx, v.tail_idx, symbols)
    }
    fmt.Println(sum)
    
    // part 2
    result_map := make(map[Index]Index)
    for _, v := range numbers {
        calculate_2(v.value, v.head_idx, v.tail_idx, symbols, result_map)
    }

    sum_2 := 0
    
    for _, v := range result_map {
        // for each result mapped we multiply the values found for the gears
        sum_2 += (v.x * v.y)
    }
    fmt.Println(sum_2)
}
