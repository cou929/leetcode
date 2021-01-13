# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

- 配列はソート済みなので、インデックスの中央から半分にしながら、再帰的に二分木をつくればよい
  - インデックス上中央の値をノードに、その左右にインデックス上左右の部分配列からできたサブツリーをつける
- せっかくなので iterative な実装に
