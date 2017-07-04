package math

func Prime() func() uint64 {
	n, primes := uint64(1), []uint64{}
	return func() uint64 {
	Outer:
		for {
			n++
			for _, prime := range primes {
				if n%prime == 0 {
					continue Outer
				}
				if prime*prime > n {
					break
				}
			}
			primes = append(primes, n)
			return n
		}
	}
}
