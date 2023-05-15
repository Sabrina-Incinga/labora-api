package repository

import (
	"fmt"
	"sync"
	"testing"
)

type totalPriceTest struct{
	price float64;
	quantity int64;
	expected float64;
}

var totalPriceTests []totalPriceTest = []totalPriceTest{
	{1000, 2, 2000},
	{1500, 1, 1500},
	{10500.99, 2, 21001.98},
	{1599, 10, 15990},
}

func TestCalculateTotalPrice(t *testing.T){
	for _, test := range totalPriceTests {
		if output:=calculateTotalPrice(test.price, test.quantity); output != test.expected {
			t.Errorf("El valor esperado %f difiere del obtenido %f", test.expected, output)
		}
	}
}

func BenchmarkCalculateTotalPrice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateTotalPrice(1550.35, 3)
	}
}

func ExampleCalculateTotalPrice(){
	fmt.Println(calculateTotalPrice(1150.10, 2))
	//Output: 2300.2
}

func TestIncrementViewCount(t *testing.T) {
	var wg sync.WaitGroup

	initialViewCount := getViewCount(1)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			GetItemById(1)
			wg.Done()
		}()
	}

	wg.Wait()

	finalViewCount := getViewCount(1)
	
	if finalViewCount != initialViewCount + 100 {
		t.Errorf("El valor esperado %d difiere del obtenido %d", initialViewCount + 100, finalViewCount)
	}
	
}