# https://leetcode.com/problems/missing-number/

- 想定される総和から実際の総和をひくだけでよい
  - nums には 0..n のすべて distinct な整数が入るので
- solution ではガウスの解放と紹介されていた
  - 等差数列の総和は `n * (n + 1) / 2` で求められるというあれ
  - 自分は気づかずにループで足し込んだ
