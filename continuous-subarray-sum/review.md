# https://leetcode.com/problems/continuous-subarray-sum/submissions/

- 早速 prefix-sum を使いそうな問題をひきあてたのでやってみたが、中途半端にしかできなかった
- prefix-sum もどき (インデックス i にそこまでの累計を格納) を作り、あとは brute-force で該当値を探した
  - 部分配列の合計が高速に計算できたのでこれでも AC だったが、実行時間は下位 10 % とかだった
- 改善として、prefix-sum に sum mod k の値を入れてなんとかできないか悩んだが力尽きた
- [disucssion](https://leetcode.com/problems/continuous-subarray-sum/discuss/99499/Java-O(n)-time-O(k)-space) より
  - mod 版 prefix-sum で、過去に同じ mod のエントリがあれば return true
  - 過去に同じ mod のエントリがある == その時点から現時点までの間に k の定数倍だけ加算されたということになるので
    - これに気づけなかった
- コツ的なものとしては、prefix-sum を考える上で、インデックスに値をいれるのを考えたほうが良さそう
  - 今回は nums の index を prefix-sum のインデックスにもしたので、発想の転換ができていなかった
  - index が累計値やその mod の場合に何ができるか、という気づきを促せないか
