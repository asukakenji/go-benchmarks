package permutation_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0"
	"github.com/asukakenji/go-benchmarks/permutation/b0a"
	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/b1a"
	"github.com/asukakenji/go-benchmarks/permutation/b2"
	"github.com/asukakenji/go-benchmarks/permutation/b3"
	"github.com/asukakenji/go-benchmarks/permutation/b4"
	"github.com/asukakenji/go-benchmarks/permutation/b5"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []struct {
		name string
		f    func([]uint, func([]uint))
	}{
		{"b0.Permutation0ParamOrder0", b0.Permutation0ParamOrder0},
		{"b0.Permutation0ParamOrder1", b0.Permutation0ParamOrder1},
		{"b0.Permutation0ParamOrder2", b0.Permutation0ParamOrder2},
		{"b0.Permutation0ParamOrder3", b0.Permutation0ParamOrder3},
		{"b0.Permutation0ParamOrder4", b0.Permutation0ParamOrder4},
		{"b0.Permutation0ParamOrder5", b0.Permutation0ParamOrder5},
		{"b0a.Permutation0ParamOrder0", b0a.Permutation0ParamOrder0},
		{"b0a.Permutation0ParamOrder1", b0a.Permutation0ParamOrder1},
		{"b0a.Permutation0ParamOrder2", b0a.Permutation0ParamOrder2},
		{"b0a.Permutation0ParamOrder3", b0a.Permutation0ParamOrder3},
		{"b0a.Permutation0ParamOrder4", b0a.Permutation0ParamOrder4},
		{"b0a.Permutation0ParamOrder5", b0a.Permutation0ParamOrder5},
		{"b1.Permutation1ParamOrder0", b1.Permutation1ParamOrder0},
		{"b1.Permutation1ParamOrder1", b1.Permutation1ParamOrder1},
		{"b1.Permutation1ParamOrder2", b1.Permutation1ParamOrder2},
		{"b1.Permutation1ParamOrder3", b1.Permutation1ParamOrder3},
		{"b1.Permutation1ParamOrder4", b1.Permutation1ParamOrder4},
		{"b1.Permutation1ParamOrder5", b1.Permutation1ParamOrder5},
		{"b1a.Permutation1ParamOrder0", b1a.Permutation1ParamOrder0},
		{"b1a.Permutation1ParamOrder1", b1a.Permutation1ParamOrder1},
		{"b1a.Permutation1ParamOrder2", b1a.Permutation1ParamOrder2},
		{"b1a.Permutation1ParamOrder3", b1a.Permutation1ParamOrder3},
		{"b1a.Permutation1ParamOrder4", b1a.Permutation1ParamOrder4},
		{"b1a.Permutation1ParamOrder5", b1a.Permutation1ParamOrder5},
		{"b2.Permutation2ParamOrder00", b2.Permutation2ParamOrder00},
		{"b2.Permutation2ParamOrder01", b2.Permutation2ParamOrder01},
		{"b2.Permutation2ParamOrder02", b2.Permutation2ParamOrder02},
		{"b2.Permutation2ParamOrder03", b2.Permutation2ParamOrder03},
		{"b2.Permutation2ParamOrder10", b2.Permutation2ParamOrder10},
		{"b2.Permutation2ParamOrder11", b2.Permutation2ParamOrder11},
		{"b2.Permutation2ParamOrder12", b2.Permutation2ParamOrder12},
		{"b2.Permutation2ParamOrder13", b2.Permutation2ParamOrder13},
		{"b2.Permutation2ParamOrder20", b2.Permutation2ParamOrder20},
		{"b2.Permutation2ParamOrder21", b2.Permutation2ParamOrder21},
		{"b2.Permutation2ParamOrder22", b2.Permutation2ParamOrder22},
		{"b2.Permutation2ParamOrder23", b2.Permutation2ParamOrder23},
		{"b2.Permutation2ParamOrder30", b2.Permutation2ParamOrder30},
		{"b2.Permutation2ParamOrder31", b2.Permutation2ParamOrder31},
		{"b2.Permutation2ParamOrder32", b2.Permutation2ParamOrder32},
		{"b2.Permutation2ParamOrder33", b2.Permutation2ParamOrder33},
		{"b2.Permutation2ParamOrder40", b2.Permutation2ParamOrder40},
		{"b2.Permutation2ParamOrder41", b2.Permutation2ParamOrder41},
		{"b2.Permutation2ParamOrder42", b2.Permutation2ParamOrder42},
		{"b2.Permutation2ParamOrder43", b2.Permutation2ParamOrder43},
		{"b2.Permutation2ParamOrder50", b2.Permutation2ParamOrder50},
		{"b2.Permutation2ParamOrder51", b2.Permutation2ParamOrder51},
		{"b2.Permutation2ParamOrder52", b2.Permutation2ParamOrder52},
		{"b2.Permutation2ParamOrder53", b2.Permutation2ParamOrder53},
		{"b3.Permutation3ParamOrder0", b3.Permutation3ParamOrder0},
		{"b3.Permutation3ParamOrder1", b3.Permutation3ParamOrder1},
		{"b3.Permutation3ParamOrder2", b3.Permutation3ParamOrder2},
		{"b3.Permutation3ParamOrder3", b3.Permutation3ParamOrder3},
		{"b3.Permutation3ParamOrder4", b3.Permutation3ParamOrder4},
		{"b3.Permutation3ParamOrder5", b3.Permutation3ParamOrder5},
		{"b4.Permutation4ParamOrder0", b4.Permutation4ParamOrder0},
		{"b4.Permutation4ParamOrder1", b4.Permutation4ParamOrder1},
		{"b4.Permutation4ParamOrder2", b4.Permutation4ParamOrder2},
		{"b4.Permutation4ParamOrder3", b4.Permutation4ParamOrder3},
		{"b4.Permutation4ParamOrder4", b4.Permutation4ParamOrder4},
		{"b4.Permutation4ParamOrder5", b4.Permutation4ParamOrder5},
		{"b5.Permutation5Inc", b5.Permutation5Inc},
	}

	for _, impl := range implementations {
		for _, c := range tcommon.Cases {
			sCopy := make([]uint, len(c.S))
			copy(sCopy, c.S)
			got := [][]uint{}
			f := func(x []uint) {
				xCopy := make([]uint, len(x))
				copy(xCopy, x)
				got = append(got, xCopy)
			}
			impl.f(c.S, f)
			if !reflect.DeepEqual(c.S, sCopy) {
				t.Errorf("%s changed the input: %v, expected %v", impl.name, c.S, sCopy)
			}
			sort.Slice(got, func(i, j int) bool {
				if len(got) == 0 {
					return false
				}
				for index, count := 0, len(got[0]); index < count; index++ {
					if got[i][index] < got[j][index] {
						return true
					}
					if got[i][index] > got[j][index] {
						return false
					}
				}
				return false
			})
			if !reflect.DeepEqual(got, c.Expected) {
				t.Errorf("%s(%v, <func>) = %v, expected %v", impl.name, c.S, got, c.Expected)
			}
		}
	}
}
