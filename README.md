# best_time_sell_and_buy

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

## Examples


```
Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
```
## 解析

要找到最佳買入以及賣出點 所算出的最大 profit

關鍵就是 賣出點為某個買入點之後的最大值

## 初步解法

要找到某個買入點之後的最大值(也就是該時間之前不去看)

最簡單的作法 把矩陣反過來找

假設 prices = [7,1,5,3,6,4]

對於 i = 5 , prices[i] = 6 這個買入點來說 最大的賣出價就是 4

對於 i = 1 , prices[i] = 1 這個買入點來說 最大的賣出價就是 6

初始化 i = len(prices) -1 , max_profit = 0, max_sell = 0

每次 對於 i 不是最後一個的值 檢查是否小於 max_sell

如果 prices[i] 小於 max_sell , 計算 profit = max_sell - prices[i] 

如果 profit > max_profit , 更新 max_profit = profit

如果 prices[i] > max_sell, 更新 max_sell = prices[i]

## 作法1

```golang
func maxProfit(prices []int) int {
    // for each i , check i < j
    max_profit := 0
    sell_max := 0
    for idx := len(prices) - 1; idx >= 0 ; idx-- {
        profit := 0
        if idx != len(prices) - 1 {
            if sell_max > prices[idx] {
                profit = sell_max - prices[idx]
            }
            if profit > max_profit {
                max_profit = profit
            }
        }
        // update sell_max
        if prices[idx] > sell_max {
            sell_max = prices[idx]
        }
    }
    return max_profit
}
```
## two points 作法

對於每個 sell 比需大於 buy

所以可以 先給定 buy = 0, sell = 1, max_sell = 0, max_profit = 0

然後 loop prices array 也就是 sell < prices 長度 

每次 檢查 prices[sell] > prices[buy]

如果 prices[sell] > prices[buy] , 計算 profit = prices[sell] - prices[buy]

如果 max_profit < profit, 更新 max_profit = profit

如果 prices[sell] <= prices[buy], 更新 buy = sell // 因為此時的賣價 是當下最低

更新 sell = sell+1

## 作法2 程式碼

```golang
package best_sell_buy

func maxProfit(prices []int) int {
	max_profit := 0
	sell := 1
	buy := 0
	max_idx := len(prices) - 1
	for sell <= max_idx {
		profit := 0
		if prices[sell] > prices[buy] {
			profit = prices[sell] - prices[buy]
			if max_profit < profit {
				max_profit = profit
			}
		} else {
			buy = sell
		}
		sell++
	}
	return max_profit
}

```