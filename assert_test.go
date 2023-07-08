package assert_test

import (
	"strings"
	"testing"

	"github.com/Jumpaku/go-assert"
)

func recovered(t *testing.T, f func()) (err error) {
	t.Helper()
	defer func() {
		err = recover().(error)
	}()
	f()
	return nil
}
func TestParams_OK(t *testing.T) {
	assert.Params(true, "must be OK")
}

func TestParams_Panic(t *testing.T) {
	err := recovered(t, func() {
		assert.Params(false, "Panic")
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}

func TestState_OK(t *testing.T) {
	assert.State(true, "must be OK")
}

func TestState_Panic(t *testing.T) {
	err := recovered(t, func() {
		assert.State(false, "Panic")
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}

func TestUnexpected_Panic(t *testing.T) {
	err := recovered(t, func() {
		assert.Unexpected("Panic")
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}

func discard[T any](T) {}
func TestUnexpected1_Panic(t *testing.T) {
	err := recovered(t, func() {
		x := assert.Unexpected1[int]("Panic")
		discard[int](x)
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}

func TestUnexpected2_Panic(t *testing.T) {
	err := recovered(t, func() {
		x, y := assert.Unexpected2[int, string]("Panic")
		discard[int](x)
		discard[string](y)
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}

func TestUnexpected3_Panic(t *testing.T) {
	err := recovered(t, func() {
		x, y, z := assert.Unexpected3[int, string, bool]("Panic")
		discard[int](x)
		discard[string](y)
		discard[bool](z)
	})
	if err == nil {
		t.Errorf("must panic")
	}
	if !strings.Contains(err.Error(), "Panic") {
		t.Errorf(`must contain "Panic"`)
	}
}
