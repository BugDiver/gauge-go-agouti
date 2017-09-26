package stepImpl

import (
	"fmt"

	"github.com/getgauge-contrib/gauge-go/testsuit"
)

var assert = func(expected, actual interface{}) {
	t := testsuit.T
	if actual != expected {
		t.Fail(fmt.Errorf(`expected  "%s" to be equal "%s"`, expected, actual))
	}
}
