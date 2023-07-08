package assert

import "fmt"

// Params panics with an error having formatted message if condition is false.
// This assertion represents that parameters must satisfy the condition.
func Params(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Errorf(`invalid params: `+format, args...))
	}
}

// State panics with an error having formatted message if condition is false.
// This assertion represents that a state of an object must satisfy the condition.
func State(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Errorf(`invalid state: `+format, args...))
	}
}

// Unexpected panics with an error having formatted message when it is called.
// This assertion represents that execution reaching this code is unexpected.
func Unexpected(format string, args ...any) {
	panic(fmt.Errorf(`unexpected state: `+format, args...))
}

// Unexpected1 panics with an error having formatted message when it is called.
// This assertion represents that execution reaching this code is unexpected.
// This function call can be used as a value to be returned, for example:
// ```go
// type Enum int
//
// const (
//
//	Enum0 Enum = iota
//	Enum1
//
// )
//
//	func F(enum Enum) string {
//	    switch enum {
//	    case Enum0:
//	         return "A"
//	    case Enum1:
//	        return "B"
//	    default:
//	        return Unexpected1[string]("unexpected enum value")
//	    }
//	}
func Unexpected1[T any](format string, args ...any) T {
	panic(fmt.Errorf(`unexpected state: `+format, args...))
}

// Unexpected2 panics with an error having formatted message when it is called.
// This assertion represents that execution reaching this code is unexpected.
// This function call can be used as values to be returned, for example:
// ```go
// type Enum int
//
// const (
//
//	Enum0 Enum = iota
//	Enum1
//
// )
//
//	func F(enum Enum) (string, int) {
//	    switch enum {
//	    case Enum0:
//	        return "A", int(enum)
//	    case Enum1:
//	        return "B", int(enum)
//	    default:
//	        return Unexpected2[string, int]("unexpected enum value")
//	    }
//	}
func Unexpected2[T1, T2 any](format string, args ...any) (T1, T2) {
	panic(fmt.Errorf(`unexpected state: `+format, args...))
}

// Unexpected3 panics with an error having formatted message when it is called.
// This assertion represents that execution reaching this code is unexpected.
// This function call can be used as values to be returned, for example:
// ```go
// type Enum int
//
// const (
//
//	Enum0 Enum = iota
//	Enum1
//
// )
//
//	func F(enum Enum) (string, int, bool) {
//	    switch enum {
//	    case Enum0:
//	        return "A", int(enum), true
//	    case Enum1:
//	        return "B", int(enum), false
//	    default:
//	        return Unexpected3[string, int, bool]("unexpected enum value")
//	    }
//	}
func Unexpected3[T1, T2, T3 any](format string, args ...any) (T1, T2, T3) {
	panic(fmt.Errorf(`unexpected state: `+format, args...))
}
