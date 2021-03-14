# https://leetcode.com/problems/increasing-triplet-subsequence/submissions/

- インデックスを 2 つ作ってそれを使って探索する方針にした
  - その位置より左側にある最小値を保持するインデックスと、その位置より右側にある最大値を保持するインデックスの 2 つを作る
  - nums を走査し、左により小さい数値、右により大きい数値があれば true
  - 3 パスだが O(N)
- submission を見るとインデックスは不要だった
  - 左から走査し、low, mid を更新しながら進み、`low < mid < cur` が成り立つか調べるだけでよい
  - 初期値は `low = nums[0]` `mid = inf`
    - inf 用に `math.MaxInt64` を使っていてなるほど。この設問の最大値は 32bit signed int
  - `low < cur < mid` なら mid を更新する
  - cur < low なら low だけを更新する
    - ここがポイント。いわゆるリセット的な状態になり、一時的に `low > mid` になる。つまり次の mid が見つかるまで true になりえない (= これまでの low, mid は捨てている)
    - 自分も最初同じようなアプローチを試してうまく行かなかったが、この部分の工夫が足りなかった
      - 単純にそこまでの low, mid をメンテしていく感じだったので、リセットパターンに対応できていなかった
      
  参考に引用
  
  ```go
  func increasingTriplet(nums []int) bool {
    n := len(nums)
    if n< 3{
        return false
    }
     
    lo, mid := nums[0], math.MaxInt64
    for _, v := range nums{
        
        if mid<v{
            return true
        }
        
        if lo<v && mid>v{
           mid = v 
        }
        
        if lo>v{
            lo = v
        }
        
    }

    return false
}
```
