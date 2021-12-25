package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeit(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("# %s duration: %+v\n", name, elapsed)
}

func getValueFreq(fd *os.File) (ones []int, zeroes []int) {
	defer timeit(time.Now(), "getValueFreq()")
	ones = make([]int, 12)
	zeroes = make([]int, 12)
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		for pos, ch := range line {
			switch ch {
			case '1':
				ones[pos]++
			case '0':
				zeroes[pos]++
			}
		}
	}
	err := scanner.Err()
	check(err)
	return
}

func calcGamma(ones, zeroes []int) (bin string) {
	defer timeit(time.Now(), "calcGamma()")
	for i := range ones {
		if ones[i] >= zeroes[i] {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return
}

func calcEpsilon(ones, zeroes []int) (bin string) {
	defer timeit(time.Now(), "calcEpsilon()")
	for i := range ones {
		if ones[i] < zeroes[i] {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return
}

func part1(fd *os.File) (result int) {
	defer timeit(time.Now(), "part1()")
	ones, zeroes := getValueFreq(fd)
	gBin := calcGamma(ones, zeroes)
	eBin := calcEpsilon(ones, zeroes)
	g, err := strconv.ParseInt(gBin, 2, 32)
	check(err)
	e, err := strconv.ParseInt(eBin, 2, 32)
	check(err)
	fmt.Printf("g: %v, bin: %v\n", g, gBin)
	fmt.Printf("e: %v, bin: %v\n", e, eBin)
	fmt.Printf("g x e: %v\n", g*e)
	result = int(g * e)
	return
}

func readLines(fd *os.File) (lines []string) {
	defer timeit(time.Now(), "readLines()")
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	check(err)
	return
}

type bitFilter func([]string, int) byte

func filterLines(lines []string, fn bitFilter) string {
	defer timeit(time.Now(), "filterLines()")
	for pos := range lines[0] {
		newLines := []string{}
		if len(lines) < 2 {
			break
		}
		bit := fn(lines, pos)
		for _, line := range lines {
			if bit == line[pos] {
				newLines = append(newLines, line)
			}
		}
		lines = newLines
	}
	return lines[0]
}

func mostCommonBit(lines []string, pos int) byte {
	// defer timeit(time.Now(), "mostCommonBit()")
	total := len(lines)
	ones := 0
	for _, line := range lines {
		if line[pos] == '1' {
			ones++
		}
	}
	if float32(ones) >= float32(total)/2 {
		return '1'
	} else {
		return '0'
	}
}

func leastCommonBit(lines []string, pos int) byte {
	// defer timeit(time.Now(), "leastCommonBit()")
	total := len(lines)
	ones := 0
	for _, line := range lines {
		if line[pos] == '1' {
			ones++
		}
	}
	if float32(ones) < float32(total)/2 {
		return '1'
	} else {
		return '0'
	}
}

func part2(fd *os.File) (result int) {
	defer timeit(time.Now(), "part2()")
	o2Lines := readLines(fd)
	co2Lines := o2Lines
	o2RatingBin := filterLines(o2Lines, mostCommonBit)
	co2RatingBin := filterLines(co2Lines, leastCommonBit)
	o2Rating, err := strconv.ParseInt(o2RatingBin, 2, 32)
	check(err)
	co2Rating, err := strconv.ParseInt(co2RatingBin, 2, 32)
	check(err)
	fmt.Printf("o2Rating: %v, bin: %v\n", o2Rating, o2RatingBin)
	fmt.Printf("co2Rating: %v, bin: %v\n", co2Rating, co2RatingBin)
	result = int(co2Rating * o2Rating)
	return
}

func main() {
	defer timeit(time.Now(), "main")
	if len(os.Args) != 2 {
		fmt.Println("please provide a filename argument")
		os.Exit(1)
	}
	filename := os.Args[1]

	fd, err := os.Open(filename)
	defer fd.Close()
	check(err)

	result1 := part1(fd)
	fmt.Printf("part1 result: %+v\n", result1)

	fd.Seek(0, io.SeekStart)

	result2 := part2(fd)
	fmt.Printf("part2 result: %+v\n", result2)
}
