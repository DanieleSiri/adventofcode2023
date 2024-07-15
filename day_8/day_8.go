package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func walk(instr string, curr byte, curr_idx int,  m map[string][2]string, s string, steps int)  int{
    // string to check if ZZZ
    if s == "ZZZ"{
        return steps
    }
    steps++
    // calculate next curr
    var next_curr byte
    next := curr_idx + 1
    if next >= len(instr) {
        next_curr = instr[0]
        next = 0
    } else {
        next_curr = instr[next]
    }
    if curr == 'L' {
        return walk(instr, next_curr, next, m, m[s][0], steps)
    } else {
        return walk(instr, next_curr, next, m, m[s][1], steps)
    }
}

func main() {
    instructions := ""
    nodes := make(map[string][2]string)
    file, err := os.Open("day_8")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line  := scanner.Text()
        if instructions == "" {
            instructions = line
            continue
        }

        if line != "" {
            s := strings.Split(line, " = ")
            r := strings.ReplaceAll(strings.ReplaceAll(s[1], ")", ""), "(", "")
            nodes[s[0]] = [2]string(strings.Split(r, ", "))
        }
    }
    res := walk(instructions, instructions[0], 0, nodes, "AAA", 0)
    fmt.Println(res)
}
