# https://leetcode.com/problems/search-in-rotated-sorted-array/

- たかだか n = 5000 なので全探索したら AC だった
- 本当は二分探索を実装するのが肝だったっぽい
  - 中途半端にしかソートされていないのでうまくあつかう必要がある
    - `l < mid < r` が単純増加じゃない場合
      - `mid < target` か `target < r` なら `l = mid` でよい
    - `l < mid < r` が単純増加の場合は普通
