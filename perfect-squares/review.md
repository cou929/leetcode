# https://leetcode.com/problems/perfect-squares/submissions/

- 入力 n の最大 10^4 までの perfect square は 100 (100 * 100 = 10^4 が max)
- 1 ~ n まで走査
- その n_i での `least number of perfect square numbers that sum to n` を保存する
- `least number of perfect square numbers that sum to n` は漸化式で求められる
  - `1 + memo[n_i - perfect sq)]` が最小のものを探せば良い
  - ただし `perfect sq <= n_i`
- perfect sq は高々 100 までなので、計算量は O(n) (O(100n))
- もっと数学な解き方はありそうだけどそこは苦手なのでスルー
