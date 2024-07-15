package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "slices"
)

var points = map[byte]int{
    '2': 1,
    '3': 2,
    '4': 3,
    '5': 4,
    '6': 5,
    '7': 6,
    '8': 7,
    '9': 8,
    'T': 9,
    'J': 10,
    'Q': 11,
    'K': 12,
    'A': 13,
}
// part 2
var points_2 = map[byte]int{
    'J': 0,
    '2': 1,
    '3': 2,
    '4': 3,
    '5': 4,
    '6': 5,
    '7': 6,
    '8': 7,
    '9': 8,
    'T': 9,
    'Q': 10,
    'K': 11,
    'A': 12,
}

type card_values struct {
    card string;
    num int
}

// utils function to return values from a map
func Values[M ~map[K]V, K comparable, V any](m M) []V {
    r := make([]V, 0, len(m))
    for _, v := range m {
        r = append(r, v)
    }
    return r
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func order_map (m map[int][]card_values) (ordered_keys []int){
    for k := range m {
        ordered_keys = append(ordered_keys, k)
    }
    slices.Sort(ordered_keys)
    return
}

func calculate(c string, n int) (ret int) {
    r := make(map[rune]int)
    for _, l := range c {
        if  _, ok := r[l]; !ok {
            r[l] = 1
        } else {
            r[l]++
        }
    }
    switch win := len(r); win {
    case 5:
        // high card
        ret = 0
    case 4:
        // pair
        ret = 1
    case 3:
        // three of a kind or two pair
        if intInSlice(3, Values(r)) {
            ret = 3
        } else {
            ret = 2
        }
    case 2:
        // four or full house
        if intInSlice(3, Values(r)) {
            ret = 4
        } else {
            ret = 5
        }
    case 1:
        //five
        ret = 6
    }
    return
}

// part 2
func calculate_2(c string, n int) (ret int) {
    r := make(map[rune]int)
    for _, l := range c {
        if  _, ok := r[l]; !ok {
            r[l] = 1
        } else {
            r[l]++
        }
    }
    // in case it's a joker we increase the maximum number of pairs
    if _, ok := r['J']; ok && r['J'] != 5{
        max := 0
        var key rune
        for k, v := range r {
            if v >= max && k != 'J' {
                max = v
                key = k
            }
        }
        r[key] += r['J']
        delete(r,'J')
    }

    switch win := len(r); win {
    case 5:
        // high card
        ret = 0
    case 4:
        // pair
        ret = 1
    case 3:
        // three of a kind or two pair
        if intInSlice(3, Values(r)) {
            ret = 3
        } else {
            ret = 2
        }
    case 2:
        // four or full house
        if intInSlice(3, Values(r)) {
            ret = 4
        } else {
            ret = 5
        }
    case 1:
        //five
        ret = 6
    }
    return
}

func winner (card1 card_values, card2 card_values, p map[byte]int) (min card_values) {
    for i:=0; i < len(card1.card); i++ {
        // compare the cards
        if p[card1.card[i]] != p[card2.card[i]] {
//            fmt.Println(card1, card2)
//            fmt.Println(card1.card[i], card2.card[i])
            if p[card1.card[i]] < p[card2.card[i]] {
               min = card1
               break
            } else {
                min = card2
                break
            }
        }
    }
    return 
}

// bubble sorting the cards
func sort_list(card_list []card_values, p map[byte]int) {
    for i:=0; i < len(card_list)-1; i++ {
        for j:=0; j < len(card_list)-i-1; j++ {
            if winner(card_list[j], card_list[j+1], p) == card_list[j+1] {
                tmp := card_list[j]
                card_list[j] = card_list[j+1]
                card_list[j+1] = tmp
            }
        }
    }
}

func order(m map[int][]card_values, p map[byte]int) (ordered_result []card_values) {
    keys := order_map(m)
    for k := range keys {
        sort_list(m[k], p)
        for _, l := range m[k] {
            ordered_result = append(ordered_result, l)
        }
    }
    return
}

func main() {
    results := make(map[int][]card_values)
    file, err := os.Open("day_7")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    results_2 := make(map[int][]card_values)
    for scanner.Scan(){
        line := scanner.Text()
        row := strings.Split(line, " ")
        num, _ := strconv.Atoi(row[1])
        new_card := card_values{row[0], num}
        index := calculate(row[0], num)
        // part 2
        index_2 := calculate_2(row[0], num)
        // calculate the poker result and then append to the map of results
        results_2[index_2] = append(results_2[index_2], new_card)
        results[index] = append(results[index], new_card)
    }
    sum := 0
    for i, v := range order(results, points) {
        sum += (i + 1) * v.num
    }
    fmt.Println(sum)
    // part 2
    sum_2 := 0
    for i, v := range order(results_2, points_2) {
        sum_2 += (i + 1) * v.num
    }
    fmt.Println(sum_2)
}
