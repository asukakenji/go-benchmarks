package forward_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/impl/onescount/forward"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, forward.OnesCountIfConstBool)
	onescount.NaiveTest(t, forward.OnesCountIfConstUint)
	onescount.NaiveTest(t, forward.OnesCountSwitchConstUint)
	onescount.NaiveTest(t, forward.OnesCountFuncPointer)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.TableTest(t, forward.OnesCountIfConstBool)
	onescount.TableTest(t, forward.OnesCountIfConstUint)
	onescount.TableTest(t, forward.OnesCountSwitchConstUint)
	onescount.TableTest(t, forward.OnesCountFuncPointer)
}
