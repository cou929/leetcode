# https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/

- むずかった。DP ぽく解こうとして袋小路にはまった
- タグを見て分割統治と書いてあったので再帰のアプローチにたどり着けた
  - 最初に入力を全部見て条件を満たしていれば終了
  - そうでなければ、invalid な文字 (k 以下の出現頻度しか無いもの) をメモする
  - 入力を invalid な文字区切りで分割する (valid な部分文字列の集合が得られる)
  - valid な部分文字列をそれぞれ再帰で渡していく
- solution によると sliding window の解法もあった
  - 何文字使うかごとに探索する (ex. `abcabcab, k=2` の場合、1 文字だけ、2 文字 (たとえば a,b )、3 文字 (a, b, c 全部使う) それぞれで探索)
  - こういう方法で次元を落とすのは思いつかなかった
- 難しいのが、計算量は brute-force も分割統治も丸めると O(n^2) なこと
  - ただより細かく出すと brute-force は O(26 * n ^2) で分割統治は O(n^2) くらいになる
  - 係数レベルの計算量の差で TLE するか分かれる設計になっている
