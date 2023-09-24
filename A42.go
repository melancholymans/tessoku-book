package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		b, e := ioutil.ReadFile("./input")
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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	n, k := ni2()
	sl := make([]IntPair, n)
	for i := 0; i < n; i++ {
		sl[i][0] = ni()
		sl[i][1] = ni()
	}
	total := 0
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			total = max(total, getscore(a, b, k, n, sl))
		}
	}
	fmt.Fprintln(wtr, total)
}

func getscore(a, b, k, n int, sl []IntPair) int {
	total := 0
	for i := 0; i < n; i++ {
		if (a <= sl[i][0] && sl[i][0] <= a+k) && (b <= sl[i][1] && sl[i][1] <= b+k) {
			total += 1
		}
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
