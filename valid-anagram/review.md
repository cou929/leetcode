# https://leetcode.com/problems/valid-anagram/

- sort して比較のアプローチ
    - 汚いけど一文字ずつの []string にバラしてソートし、ジョインして比較した
    - 計算量は本質部分に限って出した（厳密には空間などはもっと使ってるんだけど、それは言語の差異ということで）
- 別アプローチは hash でメモして探索
    - time: O(2n), space: O(n) になりそう
- input がアスキーじゃなくて unicode だとどうかという問いも書かれていた
    - golang で言う byte と rune とか、エンコーディングや内部表現の話かと思ったが
    - ascii と unicode では入力の範囲が違うと言うのが言いたかったらしかった
    - hash を使えば関係ない
