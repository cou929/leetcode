# https://leetcode.com/problems/reverse-integer/

- 一桁ずつ順に処理
    - なので計算量は log10
- 32bit のオーバーフローは愚直に確認
    - 最初は int64 に入れて 2^32 と比べるズルをした
- go の int は環境依存で 32 or 64 bit らしい
    - https://golang.org/ref/spec#Numeric_types
