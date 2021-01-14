# https://leetcode.com/problems/first-bad-version/

- シンプルな二分探索
- 実装の練習になった
    - 左右の範囲を更新していく
    - 境界値
        - 更新時に左右どちらかは +1 必要
        - 左端か右端に答えがあるケースでチェックすると良い
