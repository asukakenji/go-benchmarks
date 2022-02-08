# Permutation

## b0

The package `b0` passes the tail sub-slice as a parameter to the recursive
function. Therefore, the signature of the recursive function is:

```go
func [T any] (original, subslice []T, consumer func([]T));
```

## b0a

The package `b0a` is different from `b0` in the sense that `len(s)` is cached in
a variable common to both branches. No significant speed up or slow down is
observed. This may due to the fact that the Go compiler does a good job in
static analysis. This result agrees with the comparison between packages `b1a`
and `b1`.

## b1

The package `b1` is different from `b0` in the sense that an index, instead of a
sub-slice, is passed as a parameter in the recursion. Therefore, the signature
of the recursive function is:

```go
func [T any] (original []T, depth int, consumer func([]T));
```

This makes `b1` about `10%` faster than `b0`.

## b1a

The package `b1a` is different from `b1` in the sense that `len(s)` is cached in
a variable common to both branches. No significant speed up or slow down is
observed. This may due to the fact that the Go compiler does a good job in
static analysis. This result agrees with the comparison between packages `b0a`
and `b0`.

## b2

The package `b2` is different from `b1` in the sense that `len(s)` is passed
as the parameter `n`. Therefore, the signature of the recursive function is:

```go
func [T any] (n int, original []T, depth int, consumer func([]T));
```

This is actually slower than `b1` according to the benchmark results.

## b3

The package `b3` is different from `b1a` in the sense that a closure is used
rather than a regular recursive function. This makes `b3` about `15%` slower
than `b1a` due to the overhead of closure.

## b4

The package `b4` is different from `b1` in the sense that the recursion is done
backwards, and therefore loop variables compare to `0` instead of `len(s)`. This
makes `b4` about `1~2%` faster than `b1`. However, the code looks much cleaner
with fewer `len(s)` calls.

## b5

The package `b5` is different from all previous packages in the sense that it
implements the permutation algorithm iteratively (without recursion). It is
surprisingly slower than the recursive ones.
