package main

// maxProfit returns the maximum profit possible for the stock
func maxProfit(stockPricesByDay []int) int {
	maxProfit := 0
	for current := 0; current < len(stockPricesByDay); current++ {
		if stockPricesByDay[current] < 0 {
			maxProfit = 0
			break
		}
		for next := current + 1; next < len(stockPricesByDay); next++ {
			profit := stockPricesByDay[next] - stockPricesByDay[current]
			if profit > maxProfit {
				maxProfit = profit
			}

		}
	}
	return maxProfit

	//  {[]int{1, 2, 3, 4, 5, 6, 7}, 6},
	// 	{[]int{6, 5, 4, 3, 2, 1}, 0},
	// 	{[]int{5, 10, 15, 11, 19, 7, 3, 1}, 14},
	// 	{[]int{300, 500, 1500, 20, 40, 50, 100, 1100}, 1080},
	// {[int{95, 100, 3, 9, 5, 1, 2, 5, 2, 4}]}
	// find smallest, then find largest after smallest, difference between both is 'max profit" for now.
	// any profit in negative is zero.
	// Look for largest number before smallest number and then find smallest number before
	//

}
