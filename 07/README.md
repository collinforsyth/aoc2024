### day `07`

Day 7 had some fun Golang optimizations.

To start, the naive approach to generate all permutations can be benchmarked as follows:
```
goos: darwin
goarch: arm64
pkg: collinforsyth/aoc2024/07
cpu: Apple M3
BenchmarkPart1-8              86          13427989 ns/op        56594816 B/op     261212 allocs/op
BenchmarkPart1-8              88          13584877 ns/op        56594842 B/op     261213 allocs/op
BenchmarkPart1-8              87          14459628 ns/op        56594839 B/op     261213 allocs/op
BenchmarkPart1-8              75          15237535 ns/op        56594850 B/op     261213 allocs/op
BenchmarkPart1-8              78          15581109 ns/op        56594928 B/op     261214 allocs/op
BenchmarkPart2-8               1        2481554083 ns/op        4600978232 B/op 42763141 allocs/op
BenchmarkPart2-8               1        2465169166 ns/op        4600976232 B/op 42763117 allocs/op
BenchmarkPart2-8               1        2389017250 ns/op        4600975368 B/op 42763111 allocs/op
BenchmarkPart2-8               1        2409625667 ns/op        4600981192 B/op 42763121 allocs/op
BenchmarkPart2-8               1        2398888000 ns/op        4600977592 B/op 42763134 allocs/op
PASS
ok      collinforsyth/aoc2024/07        19.691s
```

This is pretty slow, but we can do much better by taking advantage of Go's new Range Over Function sets in order to efficiently generate all the combinations, rather than storing everything into a single array.

```
goos: darwin
goarch: arm64
pkg: collinforsyth/aoc2024/07
cpu: Apple M3
        │ no_iter_bench.txt │           iter_bench.txt            │
        │      sec/op       │    sec/op     vs base               │
Part1-8       14.460m ± ∞ ¹   4.763m ± ∞ ¹  -67.06% (p=0.008 n=5)
Part2-8         2.410 ± ∞ ¹    1.368 ± ∞ ¹  -43.22% (p=0.008 n=5)
geomean        186.7m         80.73m        -56.75%
¹ need >= 6 samples for confidence interval at level 0.95

        │ no_iter_bench.txt │            iter_bench.txt            │
        │       B/op        │     B/op       vs base               │
Part1-8     55268.4Ki ± ∞ ¹   126.3Ki ± ∞ ¹  -99.77% (p=0.008 n=5)
Part2-8      4387.8Mi ± ∞ ¹   185.8Mi ± ∞ ¹  -95.77% (p=0.008 n=5)
geomean       486.6Mi         4.788Mi        -99.02%
¹ need >= 6 samples for confidence interval at level 0.95

        │ no_iter_bench.txt │           iter_bench.txt            │
        │     allocs/op     │  allocs/op    vs base               │
Part1-8      261.213k ± ∞ ¹   1.700k ± ∞ ¹  -99.35% (p=0.008 n=5)
Part2-8        42.76M ± ∞ ¹   27.39M ± ∞ ¹  -35.95% (p=0.008 n=5)
geomean        3.342M         215.8k        -93.54%
¹ need >= 6 samples for confidence interval at level 0.95
```

This reduces the memory footprint to 1% of the initial solution.

We can also speed up quite a bit by using a more optimized `concat()` that doesn't make any string allocations.

```
goos: darwin
goarch: arm64
pkg: collinforsyth/aoc2024/07
cpu: Apple M3
        │ iter_bench.txt │           concat_bench.txt            │
        │     sec/op     │    sec/op     vs base                 │
Part1-8     4.763m ± ∞ ¹
Part2-8    1368.2m ± ∞ ¹   391.9m ± ∞ ¹  -71.36% (p=0.008 n=5)
geomean     80.73m         391.9m        -71.36%               ²
¹ need >= 6 samples for confidence interval at level 0.95
² benchmark set differs from baseline; geomeans may not be comparable

        │  iter_bench.txt  │            concat_bench.txt            │
        │       B/op       │     B/op       vs base                 │
Part1-8      126.3Ki ± ∞ ¹
Part2-8   190274.5Ki ± ∞ ¹   126.3Ki ± ∞ ¹  -99.93% (p=0.008 n=5)
geomean      4.788Mi         126.3Ki        -99.93%               ²
¹ need >= 6 samples for confidence interval at level 0.95
² benchmark set differs from baseline; geomeans may not be comparable

        │  iter_bench.txt  │           concat_bench.txt            │
        │    allocs/op     │  allocs/op    vs base                 │
Part1-8       1.700k ± ∞ ¹
Part2-8   27388.337k ± ∞ ¹   1.700k ± ∞ ¹  -99.99% (p=0.008 n=5)
geomean       215.8k         1.700k        -99.99%               ²
¹ need >= 6 samples for confidence interval at level 0.95
² benchmark set differs from baseline; geomeans may not be comparable
```

Overall, without any pruning of the search space (still brute force), there is quite a bit of fun performance optimizations.