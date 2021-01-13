# https://leetcode.com/problems/validate-binary-search-tree/

- イテレーティブな幅優先探索で、最大、最小値も保持しながらチェックする方針
    - 幅優先なのは実装の練習のため
- 解答には深さ優先探索の探索順に関する記述があり勉強になった
    - preoder, postorder, inorder (left, node, right の順に見る)
    - 今回の場合 inorder で見ていくと必ず単調増加になるはず（ならなければ偽を返す）
    - これは気づかなかった
