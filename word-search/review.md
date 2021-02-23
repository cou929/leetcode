# https://leetcode.com/problems/word-search/

- 普通に simulation, 探索
- 行儀よく iterative に実装しキューに入れる構造体を定義し毎回コピーしていたら AC だが遅くなった
- 走査済みのメモを使い回すようにし、そのために recursive に書き直したらほどほどになった
  - 絶対値だと 10 倍程度の速度、7 倍程度の空間効率の向上
