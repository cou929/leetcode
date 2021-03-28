# https://leetcode.com/problems/find-peak-element/submissions/

- O(n) の liner scan
  - 前から全探索し、前後の要素よりも大きかったら (= peak) そのインデックスを返すだけ
- bisect
  - 次の条件での bisect
    - mid 位置が peak だったら終了
    - mid 位置の左右の要素を比べ、大きい方に移動する
- ただいずれも solution の実装のほうがスマートだった
- liner scan
  - 現在値と次の位置とを比べ、現在値のほうが小さければ continue、そうでなければそれを解として返す
  - たしかに山が下った時点で break するし、単純増加だったらループを抜けたあとに最後の要素を返せば良い
  - 境界条件チェックがなくなりかなりシンプル

```java
public class Solution {
    public int findPeakElement(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] > nums[i + 1])
                return i;
        }
        return nums.length - 1;
    }
}
```

- bisect
  - 細かいが mid は `(left + right) / 2` でよい
    - 自分の実装の式展開でも求められる
  - liner scan と同じ要領で mid + 1 よりも mid が大きければ左へ戻り、そうでなければ右へ進む方針だと if 文がシンプルになる
  - 右に進む際に mid + 1 を left に代入することで無限ループを防ぐ

```java
public class Solution {
    public int findPeakElement(int[] nums) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int mid = (l + r) / 2;
            if (nums[mid] > nums[mid + 1])
                r = mid;
            else
                l = mid + 1;
        }
        return l;
    }
}
```

- bisect はやはり手になじませるかコピペできるようにしておいたほうが良さそう (コンテスト出るならだけど)
