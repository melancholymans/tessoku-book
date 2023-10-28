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
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	defer wtr.Flush()
	h, w := ni2()
	sl := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		sl[i] = make([]int, w+1)
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			sl[i][j] = ni()
		}
	}
	q := ni()
	slq := make([]point, q)
	for i := 0; i < q; i++ {
		slq[i].x1, slq[i].y1, slq[i].x2, slq[i].y2 = ni(), ni(), ni(), ni()
	}
	slz := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		slz[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			slz[i][j] = 0
		}
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			slz[i][j] = slz[i][j-1] + sl[i][j]
		}
	}
	for j := 1; j <= w; j++ {
		for i := 1; i <= h; i++ {
			slz[i][j] = slz[i-1][j] + slz[i][j]
		}
	}
	for i := 0; i < q; i++ {
		x1 := slq[i].x1
		y1 := slq[i].y1
		x2 := slq[i].x2
		y2 := slq[i].y2
		rst := slz[x2][y2] +
			slz[x1-1][y1-1] -
			slz[x1-1][y2] -
			slz[x2][y1-1]
		fmt.Fprintln(wtr, rst)
	}
}
