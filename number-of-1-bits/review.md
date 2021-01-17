# https://leetcode.com/problems/number-of-1-bits/

- 2 の mod または bit 演算で 1 bit ごと調べた
- solution より
  - 下から最初のビットが立っている位置について、n & n-1 はそれ以上の bit には影響を及ぼさず、その最下位の立っているビットだけを 0 にする
  - これを活かして `while cur != 0` のあいだ `cur = cur & (cur - 1)` するだけでいける
  - これは思いつかない
