package main

import (
	"fmt"
)

const MOD = 1_000_000_007

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j <= n; j++ {
			if dp[i][j] == 0 {
				continue
			}

			if s[i] == '(' || s[i] == '[' || s[i] == '{' {
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % MOD
			} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
				if j > 0 {
					dp[i+1][j-1] = (dp[i+1][j-1] + dp[i][j]) % MOD
				}
			} else if s[i] == '?' {
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % MOD
				if j > 0 {
					dp[i+1][j-1] = (dp[i+1][j-1] + dp[i][j]) % MOD
				}
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % MOD
				if j > 0 {
					dp[i+1][j-1] = (dp[i+1][j-1] + dp[i][j]) % MOD
				}
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % MOD
				if j > 0 {
					dp[i+1][j-1] = (dp[i+1][j-1] + dp[i][j]) % MOD
				}
			}
		}
	}

	fmt.Println(dp[n][0] % 10)
}
