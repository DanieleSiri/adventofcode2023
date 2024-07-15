//package main
//
//import (
//    "fmt"
//    "os"
//    "strings"
//    "sync"
//    "strconv"
//)
//
//func min (a, b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}
//
//func calculateDestination (destination, source, jump, k, v int, m map[int]int, t map[int]bool) {
//        // value already found
//        if _, ok := t[k]; ok && t[k]==true{
//            return
//        }
//        if v < source || v > (source + jump-1){
//            return
//        } else {
//            // get the diff from the key to source to get the jump distance for the destination mapping
//            offset := v - source
//            new_dest := destination + offset
//            m[k] = new_dest
//            t[k] = true
//        }
//}
//
//func calculate(s string, m map[int]int, t map[int]bool) {
//    //  splitting the string for the numbers
//    split_string := strings.Split(s, " ")
//    destination, _ := strconv.Atoi(split_string[0])
//    source, _ := strconv.Atoi(split_string[1])
//    jump, _ := strconv.Atoi(split_string[2])
//    var wg sync.WaitGroup
//    for k, v := range m {
//        wg.Add(1)
//        go func(l, r int) {
//            calculateDestination(destination, source, jump, l, r, m, t)
//            wg.Done()
//        }(k, v)
//    }
//    wg.Wait()
//}
//
//func main() {
//    seeds := make(map[int]int)
//    // part 2
//    seeds_2 := make(map[int]int)
//    file, err := os.ReadFile("day_5")
//    if err != nil {
//        fmt.Println(err)
//    }
//    truth_table := make(map[int]bool)
//    file_without_empty := strings.Replace(string(file), "\n\n","\n",-1)
//    for _, line := range strings.Split(file_without_empty, "\n") {
//        if line == "" || strings.Contains(line, "map"){
//            for k := range truth_table {
//                truth_table[k] = false
//            }
//            continue
//        }
//        if strings.Contains(line, "seeds") {
//            line = strings.Split(line, "seeds: ")[1]
//            for _, el := range strings.Split(line, " ") {
//                num, _ := strconv.Atoi(el)
//                // initializing high so first number is always the minimum
//                seeds[num] = num
//                truth_table[num] = false
//            }
//            // part 2
//            ranges := strings.Split(line, " ")
//            ranges_int := make([]int, 0)
//            for _, el := range ranges {
//                n, _ := strconv.Atoi(el)
//                ranges_int = append(ranges_int, n)
//            }
//            for i:=0; i<len(ranges_int); i=i+2 {
//                for j:=ranges_int[i]; j<ranges_int[i]+ranges_int[i+1]; j++ {
//                    seeds_2[j] = j
//                    truth_table[j] = false 
//                }
//            }
//        } else {
//            calculate(line, seeds, truth_table)
//            calculate(line, seeds_2, truth_table)
//        }
//    }
//    min_num := -1
//    for _, v := range seeds {
//        if min_num == -1 {
//            min_num = v
//        } else {
//            min_num = min(v, min_num)
//        }
//    }
//    // part 2
//    min_num_2 := -1
//    for _, v := range seeds_2 {
//        if min_num_2 == -1 {
//            min_num_2 = v
//        } else {
//            min_num_2 = min(v, min_num_2)
//        }
//    }
//    fmt.Println(min_num)
//    fmt.Println(min_num_2)
//}
