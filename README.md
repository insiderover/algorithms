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
