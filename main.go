package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func Average(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	avg := float64(sum) / float64(len(nums))
	return int(math.Round(avg))
}

func Median(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n%2 == 0 {
		return int(math.Round(float64(nums[n/2-1]+nums[n/2]) / 2))
	}
	return nums[n/2]
}

func Variance(nums []int, mean int) int {
	var sum float64
	for _, num := range nums {
		sum += math.Pow(float64(num-mean), 2)
	}
	variance := sum / float64(len(nums))
	return int(math.Round(variance))
}

func StandardDeviation(variance int) int {
	stdDev := math.Sqrt(float64(variance))
	return int(math.Round(stdDev))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Message Error: usage >> \n ./bin/math-skills or ./run.sh math-skills \n should contient at least two argument")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error reading number:", err)
			return
		}
		nums = append(nums, num)
	}

	if len(nums) == 0 {
		fmt.Println("No data found in the file. haha ")
		return
	}

	a := Average(nums)
	m := Median(nums)
	v := Variance(nums, a)
	s := StandardDeviation(v)

	fmt.Printf("Average: %d\n", a)
	fmt.Printf("Median: %d\n", m)
	fmt.Printf("Variance: %d\n", v)
	fmt.Printf("Standard Deviation: %d\n", s)
}
