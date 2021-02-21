# https://leetcode.com/problems/permutations/

- 全探索を上手く書けるかという問題で、意外とすんなりいけないやつ
- iterative に書いたが recurisve のほうがすんなり書けたかもしれない
- 採用済みの値を hash (map) で保持するか、残りの候補を slice で保持するかで、map のコピーが面倒で後者にしたが、そのおかげで空間効率は悪くなった模様
