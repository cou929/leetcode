# https://leetcode.com/problems/longest-valid-parentheses/

- dp として、そのインデックス時点での最長の長さを保存しながら走査
- `)` だった場合に、連続する部分までさかのぼって加算する
- `(` だった場合は無条件に 0
- solution によると概ね会ってそうだったが、スタック（もどき）と dp 配列を両方使う必要はなかったっぽい
