# https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/

- https://leetcode.com/problems/insert-delete-getrandom-o1/ の発展問題
- 同じ方針でちょっと拡張した
- insert された値の配列と、値ごとの位置のマップを持つ基本コンセプトは一緒
- 削除時には削除対象と配列の最後の要素をスワップして位置情報を崩さないのも一緒
- 今回は値の重複ありなので `map[int]int` ではなく、`map[int][]int` にして同じ値の挿入位置すべてを記憶する方針に変更
- それにともなって Remove の実装が込み入っている
- なんか勘違いしてそうだが、Remove の際、swap した要素の map を sort する必要が出てしまっている
  - 一応 AC はしているが、厳密に O(1) ではなくなってしまっている
  - swap 対象は、位置のリストの最後のものと入れ替えないといけないので、どうしてもソートが必要になる気がしたが、solution などはそうなっていないっぽい
    - solution や discussion は別言語で、hash map 相当のデータ構造の挙動もそれぞれ違う。そこまで調べてないので、どこを勘違いしているのかまだ理解していない
