# https://leetcode.com/problems/largest-number/description/

- 基本的には入力を辞書順の降順にソートしてから join すればよい
- ただし以下のケースでは単純な辞書順ではない
  - 比較する a, b の長さが違い途中まで同じ場合
  - 例えば `321` < `32` としたいが、これは辞書順とは逆
- 対応としては、ソートの比較関数にて、a, b を concat してから辞書順で比較することにした