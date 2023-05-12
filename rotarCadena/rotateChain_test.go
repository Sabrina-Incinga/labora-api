package rotateChain

import "testing"

type rotateChainTest struct{
	originalChain string;
	expectedChain string;
}

var rotateChainTestsLeft []rotateChainTest = []rotateChainTest{
	{"ABC", "CAB"},
	{"ABCD", "DABC"},
	{"XYZA", "AXYZ"},
}

func TestRotateChainLeft(t *testing.T){
	for _, item := range rotateChainTestsLeft {
		if output := RotateChainLeft(item.originalChain); output != item.expectedChain {
			t.Errorf("La cadena esperada %s no coincide con la obtenida %s", item.expectedChain, output)
		}
	}
}

var rotateChainTestsRight []rotateChainTest = []rotateChainTest{
	{"ABC", "BCA"},
	{"ABCD", "BCDA"},
	{"XYZA", "YZAX"},
}

func TestRotateChainRight(t *testing.T){
	for _, item := range rotateChainTestsRight {
		if output := RotateChainRight(item.originalChain); output != item.expectedChain {
			t.Errorf("La cadena esperada %s no coincide con la obtenida %s", item.expectedChain, output)
		}
	}
}

func BenchmarkRotateChainLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateChainLeft("ABC")
	}
}

func BenchmarkRotateChainRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateChainRight("BCA")
	}
}