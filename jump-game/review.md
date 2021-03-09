# https://leetcode.com/problems/jump-game/submissions/

- 後ろから走査し、そこにたどり着けるかを調べていく
  - そこにたどり着けない場合、`need` をインクリメントしていき、その先でたどり着ける大きな値がないか探す
  - 見つからなければ false
- solution によると Greedy に相当かな
