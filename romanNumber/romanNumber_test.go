package main

import (
	"fmt"
	"testing"
)

func TestRomanNumber(t *testing.T) {

	t.Run("Test roman number", func(t *testing.T) {
		var result, rExpect string
		result = romanNumber(1)
		rExpect = "I"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}

		result = romanNumber(4)
		rExpect = "IV"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}

		result = romanNumber(5)
		rExpect = "V"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}

		result = romanNumber(9)
		rExpect = "IX"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}

		result = romanNumber(10)
		rExpect = "X"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}

		result = romanNumber(40)
		rExpect = "XL"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}
		result = romanNumber(50)
		rExpect = "L"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}
		result = romanNumber(90)
		rExpect = "XC"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}
		result = romanNumber(100)
		rExpect = "C"
		fmt.Printf("result %#v expect %#v\n", result, rExpect)
		if result != rExpect {
			t.Errorf("expect % #v but got % #v", rExpect, result)
		}
	})
}
