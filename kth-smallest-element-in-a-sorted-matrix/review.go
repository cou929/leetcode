# https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

- heap に一列つっこんで、そこから pop するごとに次の候補も heap に追加していく
- 実際には heap ぽいことを全探索で実現
    - キューに入っているものを全て調べて最小を取り出している
        - priority がついていないキューと同じ
- go にも標準でヒープパッケージがあったらしい
    - https://golang.org/pkg/container/heap/
    - インタフェースの実装から必要なので面倒だけど
- 正直ヒープのことは完全に忘れているので、そろそろ体系的にアルゴリズムを復習しても良さそう
- もう一つの解法は二分探索
    - 二次元の領域を半分ずつ狭めていく
    - 0,0 が low, n,n が high で、その値のちょうど半分を mid に
    - mid を得るには何ステップ必要か調べる（全探索で）
    - k 以上かかったなら high = mid という具合
    - これの計算量どうなんだろう。全部配列に入れてソートかけた方が早そうな
        - k の計算をメモしたりするのかな