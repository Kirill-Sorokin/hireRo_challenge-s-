package main

func solution(arr []int) []int {
    n := len(arr)
    if n == 0 {
        return arr
    }

    result := make([]int, n)
    index := 0

    for _, value := range arr {
        if value != 0 {
            result[index] = value
            index++
            }
        }
        return result
}
