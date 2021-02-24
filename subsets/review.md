# https://leetcode.com/problems/subsets/solution/

- solution では Cascading と呼ばれる解法だった
- これまで見つけた解の全てに対して、現在の値を append したものを新たな解として追加していく
- solution によると backtracking と Lexicographic (Binary Sorted) Subsets というのがあった
- backtracing は再帰での探索
  - 再帰で呼び出した後に、自分をもとに戻す感じのが backtracking ぽいが、今度本でしっかりやろう
- Lexicographic (Binary Sorted) Subsets はビットマスクを全パターン生成する
  - ビットマスクは `2^n ~ 2^(n+1)` までの範囲でよい
  - クヌースが考えたアルゴリズムらしい
