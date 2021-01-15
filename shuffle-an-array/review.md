# https://leetcode.com/problems/shuffle-an-array/

- Design という章だったので、計算量は気にせず書けということかと勘違いしてしまった
- 実際には配列のシャッフルを書くのが主眼だったっぽい
- Fisher-Yates のアルゴリズムが solution に紹介されていた
- `math/rand.Shuffle` で適当に済ませてしまったが、これが Fisher-Yates のアルゴリズムを実装していたので AC してしまった
  - https://go.googlesource.com/go/+/go1.15.6/src/math/rand/rand.go#240
- 一応自分でも再実装
