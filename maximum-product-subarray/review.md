# https://leetcode.com/problems/maximum-product-subarray/

- 符号が変わるごとにそこを開始地点の候補として記録し、各候補について乗じていきながら、最大値を保存していく
  - nums を操作しながら最適値を保存するので、brute-force よりも計算量は減らせる
- discussion によると走査しながら最大値、最小値を保存していくだけで良かった
  - 最適解を更新しながら走査する
  - すべて 0 以上の整数から nums が構成されるとすると、最大値を保存するだけでいい
  - 負の整数も入りうるので、最小値も持っておけば、それが符号反転時に最大値になる
- これは思いつかなかった
  - 複数の値を更新するタイプの dp という考え方の切り口かな
