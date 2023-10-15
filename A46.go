package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := os.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}
func ns() string {
	sc.Scan()
	return sc.Text()
}
func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

type point struct {
	x float64
	y float64
}

func main() {
	defer wtr.Flush()
	order := make([]int, 160)
	n := ni()
	sl := make([]point, n)
	for i := 0; i < n; i++ {
		sl[i].x, sl[i].y = float64(ni()), float64(ni())
	}
	playGreedy(sl, order, n)
	for i := 0; i < n+1; i++ {
		fmt.Fprintln(wtr, order[i])
	}
}

func getDistance(sl []point, p, q int) float64 {
	return math.Sqrt((sl[p].x-sl[q].x)*(sl[p].x-sl[q].x) + (sl[p].y-sl[q].y)*(sl[p].y-sl[q].y))
}

func playGreedy(sl []point, order []int, n int) {
	currentPlace := 0
	visited := make([]bool, n)
	order[0] = 1
	visited[0] = true
	for i := 1; i < n; i++ {
		minDist := 10000.0
		minID := -1
		for j := 1; j < n; j++ {
			if visited[j] == true {
				continue
			}
			newDist := getDistance(sl, currentPlace, j)
			if minDist > newDist {
				minDist = newDist
				minID = j
			}
		}
		visited[minID] = true
		order[i] = minID + 1
		currentPlace = minID
	}
	order[n] = 1
}
