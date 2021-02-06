# https://leetcode.com/problems/queens-that-can-attack-the-king/

- やるだけだった
- 書きやすさのため queens のメモに座標を表す独自構造体を使ったが、ここを `[][]int` とかにすればメモリ使用量はもっと抑えられそう
- 8 方向への探索をミスを避けるためにコピペコードにしたが、ふつうに各方位への dx dy を for 文でまわすともっとスッキリ書ける
