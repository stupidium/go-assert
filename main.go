package assert;


//asserts that something is true
//panics if assertion fails
func T(val bool) {
	if !val {
		panic(NewAssertionError(NewStatementFalseError()));
	}
}

//asserts that two values are supposed to be equal
//panics if assertion fails
//error message includes the values being compared
func Eq[T comparable](first, second T) {
	if first!=second {
		panic(
			NewAssertionError(
				NewNotEqualError(first,second),
			),
		);
	}
}

//asserts that two values are supposed to be not equal
//panics if assertion fails
//error message includes the values being compared
func Ueq[T comparable](first, second T) {
	if first==second {
		panic(
			NewAssertionError(
				NewEqualError(first,second),
			),
		);
	}
}