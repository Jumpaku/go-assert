# go-assert
A lightweight Go library for asserting conditions related to parameters and object state.

## Overview

go-assert is a lightweight Go library designed to clearly declare assertions for the conditions related to parameters and object state within your codebase.

The library provides several key functions:

1. **assert.Params**: This function is used to assert conditions on input parameters. If a provided condition is not met, the function panics with an error message indicating invalid parameters.

2. **assert.State**: This function is used to assert conditions related to object state, such as, object properties, constraints, or invariants. If the condition is false, the function panics with an error message indicating invalid state.

3. **assert.Unexpected**: This function is used to assert unexpected code paths. If the function is called, it panics with an error message indicating that the execution has reached an unexpected point. This is useful for signaling unexpected situations during program execution.

4. **assert.UnexpectedN**: These functions are variations of the "Unexpected" function, but they can also be used at return statements of functions.

The "go-assert" library promotes code quality by ensuring that parameters and object states meet expected conditions. It provides concise and expressive syntax for asserting conditions, helping developers enforce best practices and reduce the likelihood of bugs.

## Functions

```go
package assert

import "fmt"

// Params panics with an error containing a formatted message if the condition is false.
// This assertion ensures that the parameters satisfy the condition.
func Params(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Errorf(`invalid params: `+format, args...))
	}
}

// State panics with an error containing a formatted message if the condition is false.
// This assertion ensures that the state of an object satisfies the condition.
func State(condition bool, format string, args ...any) {
	if !condition {
		panic(fmt.Errorf(`invalid state: `+format, args...))
	}
}

// Unexpected panics with an error containing a formatted message when it is called.
// This assertion represents that reaching this code is unexpected.
func Unexpected(format string, args ...any) {
	panic(fmt.Errorf(`unexpected code execution: `+format, args...))
}

// Unexpected1 panics with an error containing a formatted message when it is called.
// This assertion represents that reaching this code is unexpected.
// This function call can be used as a value to be returned.
func Unexpected1[T any](format string, args ...any) T {
	panic(fmt.Errorf(`unexpected code execution: `+format, args...))
}

// Unexpected2 panics with an error containing a formatted message when it is called.
// This assertion represents that reaching this code is unexpected.
// This function call can be used as values to be returned.
func Unexpected2[T1, T2 any](format string, args ...any) (T1, T2) {
	panic(fmt.Errorf(`unexpected code execution: `+format, args...))
}

// Unexpected3 panics with an error containing a formatted message when it is called.
// This assertion represents that reaching this code is unexpected.
// This function call can be used as values to be returned.
func Unexpected3[T1, T2, T3 any](format string, args ...any) (T1, T2, T3) {
	panic(fmt.Errorf(`unexpected code execution: `+format, args...))
}

```