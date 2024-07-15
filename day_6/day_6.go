package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math"
)

func calculate(T, D int) int{
    // quadratic formula to resolve the equation
    D++
    b1 := int(math.Floor((float64(T) + math.Sqrt(math.Pow(float64(T), 2) - (4 * float64(D))))/2))
    b2 := int(math.Ceil((float64(T) - math.Sqrt(math.Pow(float64(T), 2) - (4 * float64(D))))/2))
    return b1 - b2 + 1
}

func main() {
    time := make(map[int]int)
    distance := make(map[int]int)
    file, err := os.Open("day_6")
    if err != nil {
        fmt.Println(err)
    }
    result := 1
    
    // part 2
    time_2 := 0
    distance_2 := 0

    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        line_number := 0
        // part 2
        new_line := strings.Replace(line, " ", "", -1)
        if strings.Contains(line, "Time") {
            for _, v := range strings.Split(line, " ") {
                i, err := strconv.Atoi(v)
                if err == nil {
                    line_number++
                    time[line_number] = i
                }
            }
            // part 2
            time_2, _ = strconv.Atoi(strings.Split(new_line, ":")[1])
        } else {
            for _, v := range strings.Split(line, " ") {
                i, err := strconv.Atoi(v)
                if err == nil {
                    line_number++
                    distance[line_number] = i
                }
            }
            // part 2
            distance_2, _ = strconv.Atoi(strings.Split(new_line, ":")[1])
        }

    }
    for i:=1; i<=len(time); i++ {
        result *= calculate(time[i], distance[i])
    }
    fmt.Println(result)
    
    // part 2
    fmt.Println(calculate(time_2, distance_2))
}
