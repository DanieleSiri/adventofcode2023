package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func get_value(char rune, curr int) (ret int) {
    ascii := int(char)
    ret = ((ascii + curr) * 17) % 256
    return
}

func main() {
    file, err := os.Open("day_15")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    res := 0 
    sum := 0
    for scanner.Scan(){
        line := scanner.Text()
        chars := make([]string, 0)
        chars = strings.Split(line, ",")
        for _, s := range chars {
            for _, c := range s {
                res = get_value(c, res)
            }
            sum += res
            res = 0
        }
        fmt.Println(sum)
    }
}
