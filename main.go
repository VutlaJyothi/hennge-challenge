package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func sumNegativeFourthPowers(nums []int, pos int) int {
    if pos >= len(nums) {
        return 0
    }

    currentValue := 0
    if nums[pos] < 0 {
        currentValue = nums[pos] * nums[pos] * nums[pos] * nums[pos]
    }

    return currentValue + sumNegativeFourthPowers(nums, pos+1)
}

func processTestCases(scanner *bufio.Scanner, count int, results []int) []int {
    if count == 0 {
        return results
    }

    if !scanner.Scan() {
        return append(results, -1)
    }
    xLine := scanner.Text()
    x, err := strconv.Atoi(strings.TrimSpace(xLine))
    if err != nil {
        return append(results, -1)
    }

    if !scanner.Scan() {
        return append(results, -1)
    }
    numsLine := scanner.Text()
    parts := strings.Fields(numsLine)

    if len(parts) != x {
        results = append(results, -1)
    } else {
        nums := make([]int, x)
        for i := 0; i < x; i++ {
            val, err := strconv.Atoi(parts[i])
            if err != nil {
                return append(results, -1)
            }
            nums[i] = val
        }
        sum := sumNegativeFourthPowers(nums, 0)
        results = append(results, sum)
    }

    return processTestCases(scanner, count-1, results)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return
    }
    nLine := scanner.Text()
    n, err := strconv.Atoi(strings.TrimSpace(nLine))
    if err != nil || n < 1 || n > 100 {
        return
    }

    results := processTestCases(scanner, n, []int{})

    for _, result := range results {
        fmt.Println(result)
    }
}
