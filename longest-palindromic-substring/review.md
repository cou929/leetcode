# https://leetcode.com/problems/longest-palindromic-substring/

- インプットをループし、ループ中の各位置を回文の中央地点とみなす
- そこから左右にチェックしていく
- 回文の長さが偶数の場合もあり得るのでそのケア
    - 初期地点で、同じ文字である限り右側を伸ばす
- solution によると `Approach 4: Expand Around Center` に近かった
