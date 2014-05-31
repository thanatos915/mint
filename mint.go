package mint

import "testing"
import "fmt"
import "os"

type Mint struct {
	t *testing.T
}
type Testee struct {
	t        *testing.T
	actual   interface{}
	expected interface{}
	dry      bool
	not      bool
	Result   Result
}
type Result struct {
	OK      bool
	Message string
}

var (
	FailBase = 0
	FailType = 1
	Scolds   = map[int]string{
		FailBase: "Expected %sto be `%+v`, but actual `%+v`\n",
		FailType: "Expected %stype `%+v`, but actual `%T`\n",
	}
)

func Blend(t *testing.T) *Mint {
	return &Mint{
		t,
	}
}
func newTestee(t *testing.T, actual interface{}) *Testee {
	return &Testee{t: t, actual: actual, Result: Result{OK: true}}
}
func Expect(t *testing.T, actual interface{}) *Testee {
	return newTestee(t, actual)
}
func (testee *Testee) Dry() *Testee {
	testee.dry = true
	return testee
}
func (testee *Testee) Not() *Testee {
	testee.not = true
	return testee
}
func (testee *Testee) failed(failure int) *Testee {
	message := testee.toText(failure)
	if !testee.dry {
		testee.t.Errorf(message)
		testee.t.Fail()
		os.Exit(1)
	}
	testee.Result.OK = false
	testee.Result.Message = message
	return testee
}
func (testee *Testee) toText(fail int) string {
	not := ""
	if testee.not {
		not = "NOT "
	}
	return fmt.Sprintf(
		Scolds[fail],
		not,
		testee.expected,
		testee.actual,
	)
}
func judge(a, b interface{}, not bool) bool {
	if not {
		return a != b
	}
	return a == b
}

func (m *Mint) Expect(actual interface{}) *Testee {
	return newTestee(m.t, actual)
}
