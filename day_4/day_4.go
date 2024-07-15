package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "os/exec"
    "strconv"
)

func main() {
    file, err := os.Open("day_4")
    if err != nil {
        fmt.Println(err)
    }

    sum := 0
    sum_2 := 0

    // part 2
    winnings_2 := make(map[int]int)
    out, e := exec.Command("wc", "-l", "day_4").Output()
    if e != nil {
        fmt.Println(e)
    }
    size, _ := strconv.Atoi(strings.Split(string(out), " ")[0])

    // initializing the map as 1 for each copy of the card
    for i:=1; i<=int(size); i++ {
        winnings_2[i] = 1
    }
    line_number := 0

    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := 0
        winnings := map[string]int{} 
        // part 2
        line_number++
        offset := 1

        input := strings.Split(scanner.Text(), ": ")
        tables := strings.Split(input[1], " | ")
        for _, v := range (strings.Split(tables[0], " " )) {
            // creating map of winning numbers
            if v != "" {
                winnings[v] = 1
            }
        }
        for _, v := range (strings.Split(tables[1], " ")){
            if _, ok := winnings[v]; ok {
                // found a winning number, so we increase by 1 if it's the first time
                // or we multiply if we have been here before'
                if row == 0 {
                    row++
                } else {
                    row *= 2
                }
                // part 2
                // we repeat the addition for each copy we had from the previous rounds
                for k:=0; k<winnings_2[line_number]; k++ {
                    winnings_2[line_number+offset]++
                }
                offset++
            }
        }
        sum += row
    }
    fmt.Println(sum)

    // part 2
    for _, v := range winnings_2 {
        sum_2 += v
    }

    fmt.Println(sum_2)
}
