# https://leetcode.com/problems/course-schedule-ii/solution/

- 有向のグラフを作って探索する
- ヘッドとなるノードからの最長距離が長い順に並べれば良い
- 入力の中に複数のグラフが出来うるのと、循環があったらエラーを返すのも必要
- dfs で実装した
- セオリーを一から考えた感があったので、グラフアルゴリズムの簡単な部分は復習しといた方が良さそう
- sort パッケージの使い方を間違っていたのも要復習
    - https://play.golang.org/p/cjr4bxWOKlx
