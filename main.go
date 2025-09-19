package assert;


//asserts that something is true
//panics if assertion fails
func T(val bool) {
	if !val {
		panic(NewAssertionError(StatementFalseError{}));
	}
}

//asserts that two values are supposed to be equal
//panics if assertion fails
//error message includes the values being compared
func Eq[T comparable](first, second T) {
	if first!=second {
		panic(NewAssertionError(NotEqualError[T]{Pair[T]{
			first:first,
			second: second,
		}}));
	}
}

//asserts that two values are supposed to be not equal
//panics if assertion fails
//error message includes the values being compared
func Ueq[T comparable](first, second T) {
	if first==second {
		panic(NewAssertionError(NotEqualError[T]{Pair[T]{
			first:first,
			second: second,
		}}));
	}
}