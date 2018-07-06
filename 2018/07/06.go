package main

import (
	"fmt"
	"math"
)



func main() {
	ret := maxProfit([]int{3,2,3,1,2})
	fmt.Println(ret)

	ret1 := maxProfit1([]float64{3,2,3,1,2})
	fmt.Println(ret1)
}


/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。


思路：需要遍历一遍（因为没有排序），找到前后差异，卖低买高

 */
func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	income := 0
	l := len(prices)
	i := 0

	for i < (l - 1) {
		if prices[i] < prices[i+1] {
			income += prices[i+1] - prices[i]
		}
		i++
	}
	return income
}


/**
假设有一个数组，它的第i个元素是一支给定的股票在第i天的价格。
如果你最多只允许完成一次交易(例如,一次买卖股票),设计一个算法来找出最大利润。
 */

/**
动态规划法：
	*  以第i天为分界线，计算第i天之前进行一次交易的最大收益preProfit[i]，和第i天之后进行一次交易的最大收益postProfit[i]，最后遍历一遍，
	*  max{preProfit[i] + postProfit[i]} (0≤i≤n-1)就是最大收益。第i天之前和第i天之后进行一次的最大收益求法同Best Time to Buy and Sell Stock I。
	*
贪心：
	*
	* 换一个角度来思考这个问题，最大利润实际上就是每天的交易价格，减去上一天的价格所构成的这个数组的最大子数组。
	* 只不过需要求取一个差值，第I天的盈利为该天价格减去之前最小价格，若是小于0则表示无需做交易，盈利为0
	*/

func maxProfit1(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := 0.0

	for i := 1; i < len(prices); i++ {
		cha := float64(prices[i] - min)
		max = math.Max(max, cha)

		min = math.Min(prices[i], min)
	}
	return max
}
