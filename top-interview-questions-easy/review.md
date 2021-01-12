# https://leetcode.com/problems/palindrome-linked-list/

- time: O(n), space: O(1) で解けるか? というお題だった
- 引数を破壊する方法しか思いつかなかった
  - linked-list の中央のノードを探し、前半を reverse する
  - 中央から左右に一要素ずつ見ていき、全部一致なら true
- これでいいんだろうか...
  - もう head からは何も辿れないので、引数の linked-list は使い物にならなくなる
