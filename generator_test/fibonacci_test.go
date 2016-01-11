package generator_test

import (
	"testing"
)

var fibonacci_test_cases = [...]struct {
	n        int
	expected uint64
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{11, 89},
	{12, 144},
	{13, 233},
	{14, 377},
	{15, 610},
	{16, 987},
	{17, 1597},
	{18, 2584},
	{19, 4181},
	{20, 6765},
	{21, 10946},
	{22, 17711},
	{23, 28657},
	{24, 46368},
	{25, 75025},
	{26, 121393},
	{27, 196418},
	{28, 317811},
	{29, 514229},
	{30, 832040},
	{31, 1346269},
	{32, 2178309},
	{33, 3524578},
	{34, 5702887},
	{35, 9227465},
	{36, 14930352},
	{37, 24157817},
	{38, 39088169},
	{39, 63245986},
	{40, 102334155},
	{41, 165580141},
	{42, 267914296},
	{43, 433494437},
	{44, 701408733},
	{45, 1134903170},
	{46, 1836311903},
	{47, 2971215073},
	{48, 4807526976},
	{49, 7778742049},
	{50, 12586269025},
	{51, 20365011074},
	{52, 32951280099},
	{53, 53316291173},
	{54, 86267571272},
	{55, 139583862445},
	{56, 225851433717},
	{57, 365435296162},
	{58, 591286729879},
	{59, 956722026041},
	{60, 1548008755920},
	{61, 2504730781961},
	{62, 4052739537881},
	{63, 6557470319842},
	{64, 10610209857723},
	{65, 17167680177565},
	{66, 27777890035288},
	{67, 44945570212853},
	{68, 72723460248141},
	{69, 117669030460994},
	{70, 190392490709135},
	{71, 308061521170129},
	{72, 498454011879264},
	{73, 806515533049393},
	{74, 1304969544928657},
	{75, 2111485077978050},
	{76, 3416454622906707},
	{77, 5527939700884757},
	{78, 8944394323791464},
	{79, 14472334024676221},
	{80, 23416728348467685},
	{81, 37889062373143906},
	{82, 61305790721611591},
	{83, 99194853094755497},
	{84, 160500643816367088},
	{85, 259695496911122585},
	{86, 420196140727489673},
	{87, 679891637638612258},
	{88, 1100087778366101931},
	{89, 1779979416004714189},
	{90, 2880067194370816120},
	{91, 4660046610375530309},
	{92, 7540113804746346429},
	{93, 12200160415121876738},
}

func Fibonacci0(n int) uint64 {
	if n < 0 {
		panic("n < 0")
	}
	a, b := uint64(0), uint64(1)
	for n != 0 {
		a, b = a+b, a
		n--
	}
	return a
}

func Fibonacci1(callback func(int, uint64, bool) bool) {
	for a, b, n := uint64(0), uint64(1), int(0); n < 94; n++ {
		if !callback(n, a, true) {
			return
		}
		a, b = a+b, a
	}
	for {
		if !callback(0, 0, false) {
			return
		}
	}
}

func Fibonacci2() func() (int, uint64, bool) {
	a, b, n := uint64(0), uint64(1), int(0)
	return func() (int, uint64, bool) {
		if n0, a0, ok := n, a, n < 94; ok {
			a, b = a+b, a
			n++
			return n0, a0, true
		}
		return 0, 0, false
	}
}

func Fibonacci3() <-chan struct {
	n   int
	fib uint64
} {
	ch := make(chan struct {
		n   int
		fib uint64
	})
	go func() {
		for a, b, n := uint64(0), uint64(1), int(0); n < 94; n++ {
			ch <- struct {
				n   int
				fib uint64
			}{n, a}
			a, b = a+b, a
		}
		close(ch)
	}()
	return ch
}

func TestFibonacci0(t *testing.T) {
	for _, c := range fibonacci_test_cases {
		got := Fibonacci0(c.n)
		if got != c.expected {
			t.Errorf("Fibonacci0(%d) = %d, expected %d", c.n, got, c.expected)
		}
	}
}

func TestFibonacci1(t *testing.T) {
	counter := 0
	callback := func(n int, got uint64, ok bool) bool {
		if counter < 94 {
			if n != counter || got != fibonacci_test_cases[counter].expected || ok != true {
				t.Errorf("Fibonacci1() = (%d, %d, %t), expected (%d, %d, %t)", n, got, ok, counter, fibonacci_test_cases[counter].expected, true)
			}
		} else {
			if n != 0 || got != 0 || ok != false {
				t.Errorf("Fibonacci1() = (%d, %d, %t), expected (%d, %d, %t)", n, got, ok, 0, 0, false)
			}
		}
		counter++
		return counter < 100
	}
	Fibonacci1(callback)
}

func TestFibonacci2(t *testing.T) {
	closure := Fibonacci2()
	for counter := 0; counter < 100; counter++ {
		n, got, ok := closure()
		if counter < 94 {
			if n != counter || got != fibonacci_test_cases[counter].expected || ok != true {
				t.Errorf("Fibonacci2() = (%d, %d, %t), expected (%d, %d, %t)", n, got, ok, counter, fibonacci_test_cases[counter].expected, true)
			}
		} else {
			if n != 0 || got != 0 || ok != false {
				t.Errorf("Fibonacci2() = (%d, %d, %t), expected (%d, %d, %t)", n, got, ok, 0, 0, false)
			}
		}
	}
}

func TestFibonacci3(t *testing.T) {
	channel := Fibonacci3()
	for counter := 0; counter < 100; counter++ {
		got, ok := <-channel
		if counter < 94 {
			if got.n != counter || got.fib != fibonacci_test_cases[counter].expected || ok != true {
				t.Errorf("Fibonacci3() = (%d, %d, %t), expected (%d, %d, %t)", got.n, got.fib, ok, counter, fibonacci_test_cases[counter].expected, true)
			}
		} else {
			if got.n != 0 || got.fib != 0 || ok != false {
				t.Errorf("Fibonacci3() = (%d, %d, %t), expected (%d, %d, %t)", got.n, got.fib, ok, 0, 0, false)
			}
		}
	}
}

func BenchmarkFibonacci1a(b *testing.B) {
	callback := func(_ int, _ uint64, ok bool) bool {
		return ok
	}
	for i := 0; i < b.N; i++ {
		Fibonacci1(callback)
	}
}

func BenchmarkFibonacci1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callback := func(_ int, _ uint64, ok bool) bool {
			return ok
		}
		Fibonacci1(callback)
	}
}

func BenchmarkFibonacci2a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		closure := Fibonacci2()
		for {
			if _, _, ok := closure(); ok {
				continue
			}
			break
		}
	}
}

func BenchmarkFibonacci2b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		closure := Fibonacci2()
		for {
			if _, _, ok := closure(); !ok {
				break
			}
		}
	}
}

func BenchmarkFibonacci3a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channel := Fibonacci3()
		for {
			if _, ok := <-channel; ok {
				continue
			}
			break
		}
	}
}

func BenchmarkFibonacci3b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channel := Fibonacci3()
		for {
			if _, ok := <-channel; !ok {
				break
			}
		}
	}
}
