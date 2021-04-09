# https://leetcode.com/problems/palindrome-partitioning/solution/

- 以下をメモしながら s を走査する
  - i 文字目までの全候補を m[i] に保存
- i を走査する際
  - i 文字目を使うケースで、回文になっているか調べ、なっていたら m[i-候補文字数] + 候補文字 を m[i] に保存する
- メモリをかなり使うけど計算量は比較的少ないはず
  - メモリは使いすぎだとも思う
  - はじめパターン数を返す問題だと勘違いしてこのアプローチをとったが、組み合わせの文字列スライスをすべて返す問題だったので微妙にマッチしていないと思う
  - 普通に再帰で全探索っぽくやったほうが良かったかも
    - その場合でも高々 2^16 なので十分そう
- 全然関係ない go の append の挙動ではまった
  - append は第一引数の slice の underlying array を拡張しているっぽい
  - 非破壊だと思っていた
  - `fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&ss)))` みたいな感じでデバッグして気づく
  - ここは別途きちんと押さえておく