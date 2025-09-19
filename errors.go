package assert

import (
	"fmt"
)

//AssertionError implements error
//also supports unwrapping by the errors package to reveal underlying errors
type AssertionError struct {
	underlying_error error;
};

//implements error.Error()
func (e AssertionError) Error() string {
	const error_phrasing = "assertion error";
	return error_phrasing+": "+e.underlying_error.Error();
}

//implements the Unwrap method as specified in the 
//errors package https://pkg.go.dev/errors
func (e AssertionError) Unwrap() error {
	return e.underlying_error;
}

func NewAssertionError(underlying_error error) error {
	return &AssertionError{underlying_error};
}

//StatementFalse implements error
type StatementFalseError struct {};

//implements error.Error()
func (e StatementFalseError) Error() string {
	const error_phrasing = "statement not true";
	return error_phrasing;
}

type Pair[T comparable] struct {
	first, second T;
}

func (p Pair[T]) First() T {return p.first;}
func (p Pair[T]) Second() T {return p.second;}

//NotEqualError implements error
//stores the values that need comparison in the underlying pair,
//access through the .First() and .Second() methods
type NotEqualError[T comparable] struct{Pair[T]};

func (e NotEqualError[T]) Error() string {
	return fmt.Sprintf("%v is not equal to %v, expecting equality",e.first,e.second);
}

//EqualError implements error
//stores the values that need comparison in the underlying pair,
//access through the .First() and .Second() methods
type EqualError[T comparable] struct{Pair[T]};

func (e EqualError[T]) Error() string {
	return fmt.Sprintf("%v is equal to %v, expecting inequality",e.first,e.second);
}
