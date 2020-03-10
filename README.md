## Algorithms and data structures

#### Run all tests (from root project directory)
```
 go test ./...
```

#### Run all tests and benchmarks (from root project directory)
```
 go test -bench=. ./... -benchmem
```

### Benchmarks
#### Results of sorting benchmarks
```
BenchmarkBubbleSort-8           13748190                80.9 ns/op             0 B/op          0 allocs/op
BenchmarkSelectionSort-8        17196655                69.8 ns/op             0 B/op          0 allocs/op
BenchmarkInsertionSort-8        100000000               10.6 ns/op             0 B/op          0 allocs/op
```

#### Results of searching benchmarks
```
BenchmarkLinearSearch-8         25225782                46.9 ns/op             0 B/op          0 allocs/op
BenchmarkBinarySearch-8         471159522               2.49 ns/op             0 B/op          0 allocs/op
```

#### Results of stack benchmarks
```
BenchmarkSimpleStack_Push-8     1000000000               0.00133 ns/op         0 B/op          0 allocs/op
BenchmarkSimpleStack_Pop-8      1000000000               0.000905 ns/op        0 B/op          0 allocs/op
BenchmarkSimpleStack_Peek-8     1000000000               0.000510 ns/op        0 B/op          0 allocs/op
```

### The process of filling the Queue
#### FIFO
```
Step/Value Queue
1          [1 0 0 0 0 0 0 0 0 0]
2          [1 2 0 0 0 0 0 0 0 0]
3          [1 2 3 0 0 0 0 0 0 0]
4          [1 2 3 4 0 0 0 0 0 0]
5          [1 2 3 4 5 0 0 0 0 0]
6          [1 2 3 4 5 6 0 0 0 0]
7          [1 2 3 4 5 6 7 0 0 0]
8          [1 2 3 4 5 6 7 8 0 0]
9          [1 2 3 4 5 6 7 8 9 0]
10         [1 2 3 4 5 6 7 8 9 10]
```

#### Priority
```
Step/Value Queue
1          [1 0 0 0 0 0 0 0 0 0]
2          [2 1 0 0 0 0 0 0 0 0]
3          [3 2 1 0 0 0 0 0 0 0]
4          [4 3 2 1 0 0 0 0 0 0]
5          [5 4 3 2 1 0 0 0 0 0]
6          [6 5 4 3 2 1 0 0 0 0]
7          [7 6 5 4 3 2 1 0 0 0]
8          [8 7 6 5 4 3 2 1 0 0]
9          [9 8 7 6 5 4 3 2 1 0]
10         [10 9 8 7 6 5 4 3 2 1]
```