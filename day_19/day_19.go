package main

import (
    "fmt"
    "os"
    "flag"
    "strings"
    "strconv"
)

var operators = map[string]func(int, int) bool {
    "<" : func(a, b int) bool { return a < b},
    ">" : func(a, b int) bool { return a > b},
}

type Node struct {
    // possibly rune
    checks []Checks
    else_result string
}

type Checks struct {
    string_to_check string
    op string
    val_checked int
    result string
}

func calculate_flow(t map[string]Node, m map[string]int, start string, res *string) string {
    if start == "A" || start == "R" {
        return start
    }
    ret := exec_check(t[start].checks, 0, m, res, t[start].else_result)
    return calculate_flow(t, m, ret, res)
}

func exec_check(curr []Checks, idx int, m map[string]int, res *string, else_res string) string {
    if (*res) == "R" || (*res) == "A" {
        return (*res)
    }
    if idx == len(curr) {
        return else_res
    }
    op := curr[idx].op
    b := curr[idx].val_checked
    if operators[op](m[curr[idx].string_to_check],b) {
        return curr[idx].result
    }
    return exec_check(curr, idx+1, m, res, else_res)
}

func main() {
    flag.Parse()
	args := flag.Args()
    var file []byte
    var err error

    if args[0] == "t" {
        file, err = os.ReadFile("day_19_t")
        if err != nil {
            fmt.Println(err)
        }
    } else {
        file, err = os.ReadFile("day_19_i")
        if err != nil {
            fmt.Println(err)
        }
    }
    sum := 0
	input := string(file)
    tree := make(map[string]Node)
    s :=  strings.Split(input, "\n\n")
    workflows := s[0]
    flux := s[1]
    for _, w := range strings.Split(workflows, "\n") {
        split_w := strings.Split(w, "{")
        nodename := split_w[0]
        split_checks := strings.Split(split_w[1], ",")
        new_node := Node{}
        for k, text := range split_checks {
            new_checks := Checks{}
            if k == len(split_checks)-1 {
                new_node.else_result = text[:len(text)-1]
                continue
            }
            inner := strings.Split(text, ":")
            new_checks.result = inner[1]
            new_checks.string_to_check = string(inner[0][0])
            new_checks.op = string(inner[0][1])
            i, _ := strconv.Atoi(inner[0][2:])
            new_checks.val_checked = i
            new_node.checks = append(new_node.checks, new_checks)
        }
        tree[nodename] = new_node
    }
    // split entries
    flux = strings.ReplaceAll(flux, "{", "")
    flux = strings.ReplaceAll(flux, "}", "")
    for _, l := range strings.Split(flux, "\n") {
        check_val := make(map[string]int)
        if len(l) == 0 {
            continue
        }
        for _, i := range strings.Split(l, ",") {
            str := strings.Split(i, "=")
            n, _ := strconv.Atoi(str[1])
            check_val[str[0]] = n
        }
        ret_val := ""
        if calculate_flow(tree, check_val, "in", &ret_val) == "A" {
            for _, el := range check_val {
                sum += el
            }
        }
    }
    fmt.Println(sum)
}
