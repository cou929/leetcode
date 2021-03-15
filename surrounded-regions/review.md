# https://leetcode.com/problems/surrounded-regions/submissions/

- TLE しまくりだがなんとか AC
- 方針
  - 'O' を見つけ次第 dfs で領域探索
  - 領域の各位置の配列と、matrix の端に達するかどうかを bool で返す
  - 達していなかったらその領域を 'X' で置き換える
- 遅かったところ
  - 計算量は O(nm) とたかを括っていたが、毎回 slice の領域確保をしていたら TLE した
  - 最初は無駄な探索をしていると踏んでデバッグしていたら時間を浪費した
    - じつはそうではなく (計算量の見積もりは正しかった)、ひとつひとつのオペレーション (特にメモリ確保) が高コストという、普通のチューニング問題になっていた
- submissions を見て
  - matrix の外周から dfs し "だめな領域" だけマークしておく
  - 次に全探索し、だめな領域は "O" のまま、そうでない "O" は単純に "X" に置き換える
  - 内部の "良い領域" の DFS をしなくてよくて高速
  - たしかにそうだ。思いつかなかった
  - けど、大雑把な計算量はどちらも O(nm) なので、「計算量の削減以外のチューニングを求められることはない」という先入観が自分の中にあったのが原因だと思う
    - 最近はどうか知らないけど昔の TopCoder だとそういう問題はなかった記憶 (たぶんコンテストが 1 時間だから?)
  - こういう TopCoder SRM の過去の先入観で損していることが多い
    - srm だとこのレベルの実装複雑さの問題は出ないだろう => leetcode だと出た。とか