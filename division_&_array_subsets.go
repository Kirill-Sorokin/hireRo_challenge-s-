package main

func solution(nums []int, target int) int {
    remainderCount := make([]int, target)
    for _, num := range nums {
        remainderCount[num % target]++
    }

    maxSubsetSize := min(remainderCount[0], 1)

    for i := 1; i <= target/2; i++ {
        if i != target-i {
            maxSubsetSize += max(remainderCount[i], remainderCount[target-i])
        } else {
            maxSubsetSize += min(remainderCount[i], 1)
        }
    }

    return maxSubsetSize
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
