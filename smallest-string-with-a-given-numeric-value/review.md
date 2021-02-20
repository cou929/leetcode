# https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/

- greedy
- 先頭からできるだけ a をつめていく。つまり後ろからできるだけ大きいアルファベットを入れていくという知識で解いた
- 後ろから z を、足りなければ入れられる最大値を入れていく
- 入れられる最大値は、その位置より前が全て a だと仮定した残りの k
