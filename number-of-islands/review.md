# https://leetcode.com/problems/number-of-islands/description/

- nxm サイズのメモに記録しながらカウントする
- 同一の島かは、1 を見つけたら前後左右に再帰的に繋がっている領域を拡張する
- grid が何故か byte なのに気づかず時間を使った
