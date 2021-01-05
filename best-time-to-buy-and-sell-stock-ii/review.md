# [Best Time to Buy and Sell Stock II \- LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)

- 深さ優先探索 & メモ化で一応 ac
- 全探索するにしても、solution は幅優先だった
    - こっちのほうがコールスタックが浅くなりそう
    - 計算量は `O(n^n)` とのことだが、自分のコードでは `O(2^n)` の認識だった
- あとは peak vallery aporeach
    - 要は dp
- ついつい探索木 & 枝刈り or メモ化のアプローチをしてしまうが、探索木を造った時点でそこから dp には行けない気がする (経験則)
    - それよりも設問をよく見て別のモデル化を検討したほうが良さそう
    - 今回の場合は prices を二次元にプロットすれば見えるものがあったかもしれない
- 深さ優先探索、幅優先探索は手に馴染ませないと時間がかかる
    - 再起じゃなくてループで書けるようになっておきたい
    - 当時はできていた
