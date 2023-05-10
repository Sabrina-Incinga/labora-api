package math

import ("testing"
	"fmt")

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
	/*got := Add(4,6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}*/
}

func BenchmarkAdd(b *testing.B){
	for i:=0;i<b.N;i++{
		Add(4,6)
	}
}

func ExampleAdd() {
    fmt.Println(Add(4, 6))
    // Output: 10
}

type factorialTest struct{
	number, expected int
}

var factorialTests = []factorialTest{
	{0, 1},
	{1, 1},
	{6, 720},
	{3, 6},
}

func TestFactorial(t *testing.T){
	for _,test := range factorialTests{
		if output := Factorial(test.number); output != test.expected{
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

