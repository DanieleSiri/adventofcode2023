package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "flag"
)

type Position struct {
    x float64
    y float64
    z float64
}

type Velocity struct {
    x float64
    y float64
    z float64
}

type Hail struct {
   pos Position
   vel Velocity
}

func check(a, b Hail, limit_x, limit_y float64) bool {
    if a.vel.y*b.vel.x == b.vel.y*a.vel.x {
        return false
    }

    t1 := (b.vel.y*(a.pos.x - b.pos.x) - b.vel.x*(a.pos.y - b.pos.y)) / ((a.vel.y*b.vel.x) - (a.vel.x*b.vel.y))
    t2 := (a.vel.y*(b.pos.x - a.pos.x) - a.vel.x*(b.pos.y - a.pos.y)) / ((b.vel.y*a.vel.x) - (b.vel.x*a.vel.y))

    if t1 > 0 && ((a.pos.x+(t1*a.vel.x) > limit_x) && (a.pos.x+(t1*a.vel.x)<limit_y)) && t2 > 0 && ((a.pos.y+(t1*a.vel.y)>limit_x) && (a.pos.y+(t1*a.vel.y)<limit_y)) {
        return true
    }
    return false
}

func createHail (positions, velocities []string) Hail {
    x, _ := strconv.Atoi(positions[0])
    y, _ := strconv.Atoi(positions[1])
    z, _ := strconv.Atoi(positions[2])

    new_vel := make([]string, 0)
    for _, v := range velocities {
        s := strings.ReplaceAll(v, " ", "")
        new_vel = append(new_vel, s)
    }

    vx, _ := strconv.Atoi(new_vel[0])
    vy, _ := strconv.Atoi(new_vel[1])
    vz, _ := strconv.Atoi(new_vel[2])

    new_hail := Hail {Position{float64(x),float64(y),float64(z)}, Velocity{float64(vx),float64(vy),float64(vz)}}
    return new_hail
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 || len(args) == 1 {
		panic("please insert 2 arguments (the 2 limits to check)")
	} 
    limit_bot, _ := strconv.Atoi(args[0])
    limit_top, _ := strconv.Atoi(args[1])

    file, err := os.Open("day_24")
    if err != nil {
        fmt.Println(err)
    }
    hailstones := make([]Hail, 0)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        split_line := strings.Split(line, " @ ")
        positions := strings.Split(split_line[0], ", ")
        velocities := strings.Split(split_line[1], ", ")
        hailstones = append(hailstones, createHail(positions, velocities)) 
    }
    counter := 0
    for i:=0; i<len(hailstones)-1; i++ {
        for j:=i+1; j<len(hailstones); j++ {
            if check(hailstones[i], hailstones[j], float64(limit_bot), float64(limit_top)) {
                counter++
            }
        }
    }
    fmt.Println(counter)
    // 684195328708898
}
