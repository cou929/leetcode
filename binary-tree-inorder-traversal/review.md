# https://leetcode.com/problems/binary-tree-inorder-traversal/

- 実装するだけだけど、iterative にやるのが面倒
- TreeNode をラップして visited かどうか状態をもたせることにした
  - visited だったら travarse せず単に値を結果に入れる
- 要は travarse 順と結果のセット順が異なっているのを、再帰だとスッキリ書けるがループだと書きづらい問題
