package intswitch_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount"
	"github.com/asukakenji/go-benchmarks/math/bits/internal/onescount/intswitch"
)

func TestOnesCountAllNaive(t *testing.T) {
	onescount.NaiveTest(t, intswitch.OnesCountIfConstBool)
	onescount.NaiveTest(t, intswitch.OnesCountIfConstUint)
	onescount.NaiveTest(t, intswitch.OnesCountSwitchConstUint)
	onescount.NaiveTest(t, intswitch.OnesCountFuncPointer)
}

func TestOnesCountAllTable(t *testing.T) {
	onescount.TableTest(t, intswitch.OnesCountIfConstBool)
	onescount.TableTest(t, intswitch.OnesCountIfConstUint)
	onescount.TableTest(t, intswitch.OnesCountSwitchConstUint)
	onescount.TableTest(t, intswitch.OnesCountFuncPointer)
}
