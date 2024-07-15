package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

var result_map = map[string]int {
    "red" : 12,
    "green" : 13,
    "blue" : 14,
}

func parse_games(s string) (x, y string) {
    res1 := strings.Split(s, ":")
    x = res1[0]
    y = res1[1]
    return
}


func is_valid(color string, n int) bool {
    if n <= result_map[color] {
        return true
    } else {
        return false
    }
}

func calculate(roundid string, s string) int{
    s = strings.Replace(s, ";", ",", -1)
    split_string := strings.Split(s, " ")

    // start from 2 because first match is " " and second is the number we are looking for
    // increment by 2 to get the color and coming back with i-1 for the number
    for i:=2; i < len(split_string); i=i+2 {
        str := strings.Replace(split_string[i], ",", "", -1)
        n, _ := strconv.Atoi(split_string[i-1])
        if !is_valid(str, n) {
            return 0
        }
    }
    gameid, _ := strconv.Atoi(strings.Split(roundid, " ")[1])
    return gameid
 }
 
// part 2
func calculate_2(s string) int {
    max_map := map[string] int {
        "red" : 0,
        "blue" : 0,
        "green" : 0,
    }

    s = strings.Replace(s, ";", ",", -1)
    split_string := strings.Split(s, " ")

    // start from 2 because first match is " " and second is the number we are looking for
    // increment by 2 to get the color and coming back with i-1 for the number
    for i:=2; i < len(split_string); i=i+2 {
        str := strings.Replace(split_string[i], ",", "", -1)
        n, _ := strconv.Atoi(split_string[i-1])
        if n > max_map[str] {
            max_map[str] = n
        }
    }
    res := 1
    for _, v := range max_map {
        res *= v
    }

    return res
}

 func main(){
     file, err := os.Open("day_2")
     if err != nil {
         fmt.Println(err)
     }
     sum := 0
     sum_2 := 0
     defer file.Close()

     scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        round, rolls := parse_games(scanner.Text())
        sum += calculate(round, rolls)    
        sum_2 += calculate_2(rolls)
    }
    
    fmt.Println(sum)
    fmt.Println(sum_2)
}
