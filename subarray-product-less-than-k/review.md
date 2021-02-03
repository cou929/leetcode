# https://leetcode.com/problems/subarray-product-less-than-k/

- 最初は prefix sum の改変で、和の累計ではなく乗じていった数の累計を保存するアプローチをとった
  - `ps[index] = products` (`ps[products] = index` ではなく)
  - prefix sum は単純増加なので、k 以上になる場所を見つければそれより右の要素数 = 求めたい組み合わせ数になる
  - 場所を求めるのはリニアサーチでとりあえず
- この方針はオーバーフローするのでだめだった
- 次は sliding window でこれは AC
- ほぼ最適解で書けたが、submission では下位二割
  - 同じコードで何度か sumit すると 60ms ほどぶれがあり、それに応じて上位 15% まであがったりした
  - 実行時間の分布が狭くて機能していない
- solution によると prefix sum のアプローチも紹介されていた
  - nums の乗算をしていくのではなく、log(nums[i]) を加算していくという変換をしていて、これはかなりなるほどだった
  - prefix sum 配列からの探索は二分探索にしていた。これは無理やりだなと思った
    - けど二分探索をさっと書けるようにはしておくと良さそう
