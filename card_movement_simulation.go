package main

func solution(cards [][]int, moves [][]int, query int) []int {
    cardPosition := make(map[int][]int)
    for _, card := range cards {
        cardID, row, col := card[0], card[1], card[2]
        cardPosition[cardID] = []int{row, col}
    }

    // Process moves
    for _, move := range moves {
        cardID, startRow, startCol, destRow, destCol := move[0], move[1], move[2], move[3], move[4]

        if pos, exists := cardPosition[cardID]; exists && pos[0] == startRow && pos[1] == startCol {
            // Move the card to the new position
            cardPosition[cardID] = []int{destRow, destCol}

            // Move other cards in the destination column down
            for id, position := range cardPosition {
                if id != cardID && position[1] == destCol && position[0] >= destRow {
                    cardPosition[id][0]++
                }
            }

            // Move other cards in the original column up
            for id, position := range cardPosition {
                if id != cardID && position[1] == startCol && position[0] > startRow {
                    cardPosition[id][0]--
                }
            }
        }
    }

    // Return the final position of the queried card
    if pos, exists := cardPosition[query]; exists {
        return pos
    }
    return []int{}
}
