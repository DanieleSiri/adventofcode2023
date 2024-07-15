package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func calculate(s string) int {
    first_digit := int(s[strings.IndexAny(s, "123456789")]-'0')
    last_digit := int(s[strings.LastIndexAny(s, "123456789")]-'0')
    res := (10 * first_digit) + last_digit
    return res
}

func main() {
    file, err := os.Open("day_1")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    sum := 0
    sum_2 := 0
    for scanner.Scan(){
        line := scanner.Text()
        sum += calculate(line)
        // part 2
        r := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
        // need to replace it twice for cases like "twone" that are 2 numbers
        line = r.Replace(r.Replace(line))
        sum_2 += calculate(line)
    }
    fmt.Println(sum)
    fmt.Println(sum_2)
}
