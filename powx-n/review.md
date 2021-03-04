# https://leetcode.com/problems/powx-n/submissions/

- オーバーフローを考えなくて良い問題設定
- 力技で ac してしまった
  - 普通に n 回ループで回す
  - x == 1, -1, 0 の場合だけ特別扱いして、TLE 狙いのテストケースを回避
- これだとあまりにもなので普通にもやる
- pow(x, n/2) を計算して、その 2 倍を return することで、時間効率を logn にする
