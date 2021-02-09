# https://leetcode.com/problems/partition-array-into-disjoint-intervals/

- 解は必ず存在し、部分配列じゃなくて入力を二分する位置を探すだけなので、変数が少なく簡単
- 右から走査しその位置での累積最大値をメモして、あとは左から解を探すだけ
- 全探索でも n^2 (tle するだろうけど)
