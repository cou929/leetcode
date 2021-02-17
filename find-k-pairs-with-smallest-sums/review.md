# https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

- kth-smallest-element-in-a-sorted-matrix と似た感じだけど、妙にハマった
- 方針としては前回と同じように heap をもどきで全探索しつつやる感じ
- nums1, 2 を x, y 軸にして探索していくさまを二次元で捉えるとイメージしやすかった

```
https://leetcode.com/problems/find-k-pairs-with-smallest-sums/discuss/84550/Slow-1-liner-to-Fast-solutions のコピペだが、こういうイメージ

# # # # # ? . .
# # # ? . . . .
# ? . . . . . .   "#" means pair already in the output
# ? . . . . . .   "?" means pair currently in the queue
# ? . . . . . .
? . . . . . . .
. . . . . . . .
```

- これを凝ってやろうとして時間を使った
  - 上の例だと、一番下の行は探す必要がない
  - ヒープを使わないでループでやると、境界条件がかなりややこしいことになる
- https://leetcode.com/problems/find-k-pairs-with-smallest-sums/discuss/84550/Slow-1-liner-to-Fast-solutions みたいにヒープを使うと
  - pop して、その隣接要素だけ push すればいい、他の候補は既に push されているので、一定シンプルに書ける
  - 木の探索と近い感覚になる
- この方向が悪手なのはわかっていたが、それでもやりきれるか試してみたが、全然だめだった
- 普通にヒープを使う方も練習しておいたほうがいいかも
