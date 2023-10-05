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

func main() {
	defer wtr.Flush()
	n, q := ni2()
	state := 1
	sl := make([]int, n)
	for i := 0; i < n; i++ {
		sl[i] = i + 1
	}
	for i := 0; i < q; i++ {
		t := ni()
		if t == 1 {
			x := ni()
			y := ni()
			if state == 1 {
				sl[x-1] = y
			}
			if state == 2 {
				sl[n-x] = y
			}
		}
		if t == 2 {
			if state == 1 {
				state = 2
			} else {
				state = 1
			}
		}
		if t == 3 {
			x := ni()
			if state == 1 {
				fmt.Fprintln(wtr, sl[x-1])
			}
			if state == 2 {
				fmt.Fprintln(wtr, sl[n-x])
			}
		}
	}
}
