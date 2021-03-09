# https://leetcode.com/problems/merge-intervals/submissions/

- 入力をソートすることで、途中解の最後の要素と次の候補とだけを比べればよくなる
- 計算量はほぼソート分だけになる
- solution によるともっとナイーブな解として、オーバーラップしている要素をつないだグラフを作るアプローチが紹介されていた
  - グラフに要素を挿入する際に全探索が必要なので計算量は `O(n^2)` になるっぽい