package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    file, err := os.Open("day_5")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){

    }
}