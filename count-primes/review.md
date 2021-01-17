# https://leetcode.com/problems/count-primes/

- エラトステネスの篩
- solution 見てから付け足した最適化
    - sqrt(n) までしか調べなくて良い
        - 素数かどうか = 合成数かどうか
        - n が p で割り切れる場合、n = p * q
        - ここで素数判定においては p, q の順列は考えなくていいので、p <= q だけ調べれば良い
        - そうなると p の最大値は sqrt(n) になる
    - 素数 p が見つかってその倍数を非素数としてマークする際、 p*p からはじめて p*p+p, p*p+2p, ... と処理していけば良い
        - p ~ p*p の範囲はこれまでの処理でマーク済みなので
- 計算量は O(nloglog n) らしいが理解しきれていない