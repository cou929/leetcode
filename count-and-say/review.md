# https://leetcode.com/problems/count-and-say/

- naive に実装
- 計算量をうまく考えられなかった
    - 解答より、ワーストケースを考えればよかった
    - ワーストケースは同じ値が全く連続しない場合で、長さが二倍になる
    - 最初が 1 なので計算量は 2^n
- 正規表現のアプローチもあるらしい
    - パターンの連続をキャプチャする方法は知らなかった
        - 今まで実際に必要が発生しなかったのか