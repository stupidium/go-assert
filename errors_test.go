package assert

import (
	"fmt"
	"testing"
)

func TestTypeCorrectness(t *testing.T) {
	var err error;
	//ensure all error types implement the error interface
	err = AssertionError{};
	err = StatementFalseError{};
	err = NotEqualError[int]{};
	err = EqualError[int]{};
	//ensure the error constructors return the correct types
	//compiler check isn't sufficient as the return value type is an interface
	err = testConstructor[*AssertionError](NewAssertionError(nil));
	if err!=nil {t.Error(err);}
	err = testConstructor[StatementFalseError](NewStatementFalseError())
	if err!=nil {t.Error(err);}
	err = testConstructor[*NotEqualError[int]](NewNotEqualError(1,2));
	if err!=nil {t.Error(err);}
	err = testConstructor[*EqualError[int]](NewEqualError(1,2));
	if err!=nil {t.Error(err);}
}

func testConstructor[T any](value interface{}) error {
	var a T;
	if _,ok:=value.(T); !ok {
		return fmt.Errorf("%T constructor returns an undesired type:%T",a,value);
	}
	return nil;
}

//shows a few examples of error messages
//use the -v flag if you wanna see the output of this function
func Test_ShowErrorMessages(t *testing.T) {
	t.Log("Below are some logs to showcase example error messages. They are not real error message");
	var logBoth = func(val error) {
		t.Log(stringifyLog(val));
		t.Log(stringifyLog(NewAssertionError(val)));
	};
	logBoth(NewStatementFalseError());
	logBoth(NewEqualError(1,2));
	logBoth(NewNotEqualError(1,2));
}


func stringifyLog(val interface{}) string {
	return fmt.Sprintf("%T stringifies to \"%v\"", val, val);
}