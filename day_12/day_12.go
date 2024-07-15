package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func create_states(a []int) (s []rune) {
    s = append(s, '.')
    for _, n := range a {
        for range n {
            s = append(s, '#')
        }
        s = append(s, '.')
    }
    return
}

func arrangements (states []rune, s string) (ret int) {
    new_map := make(map[int]int)
    states_map := make(map[int]int)
    states_map[0] = 1
    for _, char := range s {
        for i := range states {
            if char == '?' {
                if i+1 < len(states) {
                    new_map[i+1] = new_map[i+1] + states_map[i]
                }
                if states[i] == '.' {
                    new_map[i] = new_map[i] + states_map[i]
                }
            }
            if char == '.' {
                if i+1 < len(states) && states[i+1] == '.' {
                    new_map[i+1] = new_map[i+1] + states_map[i]
                }
                if states[i] == '.' {
                    new_map[i] = new_map[i] + states_map[i]
                }
            }
            if char == '#' {
                if i+1 < len(states) && states[i+1] == '#' {
                    new_map[i+1] = new_map[i+1] + states_map[i]
                }
            }
        }
        states_map = new_map
        new_map = make(map[int]int)
    }
    ret = states_map[len(states)-1] + states_map[len(states)-2]
    return 
}

func main() {
    file, err := os.Open("day_12")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        split := strings.Split(line, " ")
        springs := split[0]
        groups_s := split[1]
        groups := make([]int, 0)
        for  _, v := range strings.Split(groups_s, ",") {
            i, _ := strconv.Atoi(v)
            groups = append(groups, i)
        }
        // part 2
        init_spring := springs
        init_group := groups
        for i:=0; i<4; i++ {
            springs += "?"
            springs += init_spring
            for _, v := range init_group {
                groups = append(groups,v)
            }
        }
        state_list := create_states(groups)
        sum += arrangements(state_list, springs)
    }
    fmt.Println(sum)
}
