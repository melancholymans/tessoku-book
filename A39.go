package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
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
	n := ni()
	sl := make([]IntPair, n)
	for i := 0; i < n; i++ {
		sl[i][0] = ni()
		sl[i][1] = ni()
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][1] < sl[j][1]
	})
	current_time := 0
	rlt := 0
	for i := 0; i < n; i++ {
		if current_time <= sl[i][0] {
			current_time = sl[i][1]
			rlt += 1
		}
	}
	fmt.Fprintln(wtr, rlt)
}
