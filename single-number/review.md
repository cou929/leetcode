# https://leetcode.com/problems/single-number/description/

- 数学的な解法があるんだろうなと思いつつ普通にやった
- 数学的な解法
    - `2*(a+b+c)−(a+a+b+b+c)=c`
- xor を使って同じ数値同士を打ち消すアプローチはなるほどだった
    - `a xor b xor a` => b が残る
- go の sort はクイックソート
    - 入力のサイズによってはヒープソートに切り替えているぽい
    - https://github.com/golang/go/blob/0c5afc4fb7e3349ec4efdce6554f83554e3d087c/src/sort/sort.go#L198
        - 実装は追っておらず関数名でそう認識しただけ
