# https://leetcode.com/problems/group-anagrams/solution/

- ソート済みの文字列をインデックスにしたメモを作って解いた
  - 文字列数を n, 各文字列の長さを k とすると `O(n * klog(k))`
- ソートせずに、文字の登場頻度の分布をインデックスにしたメモでもよかった
  - `O(n * k)` になる
  - ただ golang だと slice を文字列にデコードしたものなどを map のキーにしないといけなくて微妙っちゃ微妙
