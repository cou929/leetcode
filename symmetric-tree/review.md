# https://leetcode.com/problems/symmetric-tree/

- 幅優先探索
- 同じ depth のノードを比較
- null が混じっていてもそれがシンメトリーなら OK というのにサブミットするまで気が付かなかった
  - 問題文に書いてないような...
- recursive 版は割愛。そっちのほうが簡単だと思うので
- solution はあえて同じノードを 2 つ渡し、その左右が同じか (鏡写しか) をしらべていく書き方
  - 自分のよりもスマート
