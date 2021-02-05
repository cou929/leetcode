# https://leetcode.com/problems/subarray-sum-equals-k/

- 今回も prefix sum だが負数も混ざっているので工夫
- map に key は累積、値はその累積値になっている index の個数を保存する
- 各走査について、その時の累積 - k が map にあれば、その map の key の個数をカウントしていく
    - map にはその累積値になる全パターンが保存されているのでそれを足しこむ
    - 過去に走査した分だけが value にカウントされているので、それを足すだけで良い
    - 例えば [1,1,1,-1,-1,3], k = 2 の場合、m[2] になるのは index 1,3 の二箇所
        - index 5 の時点で、sum = 4, 4-k = 2 となり、m[2] には index 1,3 からの 2 が value として入っている
        - これがそれぞれ [1,5], [3,5] の二つの区間（部分配列）に相当する