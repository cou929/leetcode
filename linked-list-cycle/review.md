# https://leetcode.com/problems/linked-list-cycle/

- 走査済みのノードにマークを付ける方針だと超簡単だったんだけど、流石にこれはだめだった
  - 入力を破壊しているので
- solution によると Floyd's Cycle Finding Algorithm が紹介されていて関心した
  - 2 つのカーソルを準備し、それぞれ別々の速度で動く
  - サイクルがある場合は 2 つがいつか重なるので、それで判定できる
  - https://en.wikipedia.org/wiki/Cycle_detection
    - 亀と野うさぎ
