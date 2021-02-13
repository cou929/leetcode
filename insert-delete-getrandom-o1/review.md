# https://leetcode.com/problems/insert-delete-getrandom-o1/

- hash map と配列を保持、更新していく
- hash にはインデックスは set に含まれる数値、値は配列上の位置
- Remove の時、対象の数多と配列末尾をスワップして、配列末尾を削除すれば O(1)
- get は乱数で配列から一つ選べばいい
- swap 方針を最初思いつかず、削除時には hash からだけ消して、get 時には見つかるまでループする方針にした
    - これでも ac したが、削除が o(1) でない
