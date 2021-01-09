# https://leetcode.com/problems/implement-strstr/

- 単純に探索
- 以下は submit してからテストケースで落ちて気づいた
    - needle の方が長い
    - needle の探索時に haystack の残りをはみ出す
    - needle の探索分を haystack のカーソルに足していたが、それでは見つからないケース
        - "misisippi" から "isippi" を探すケース
        - 最初の i が違うとわかった時、index 3 の s から探索を再開する必要がある
- これらは先に気付きたい
    - でないと実戦でもテストが書けない
