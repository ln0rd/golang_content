// unit test
package test

import (
	"introduction-test/address"
	"testing"
)

type scenaryTest struct {
	insertedAddress string
	expectedReturn string
}

func TestKindOfAddress(t *testing.T)  {
	// Test 

	// rodar em paralelo
	t.Parallel()

	testScenarios := []scenaryTest{
		{"rua abc", "Rua"},
		{"avenida brasil", "Avenida"},
		{"", "invalid type"},
		{"rua Jardim sorocity", "Rua"},
	}

	// if kindOfAddressRecived != kindOfAddressExpected {
	// 	t.Errorf("Kind of address is wrong recived [ %s ] expected [ %s ]", kindOfAddressRecived, kindOfAddressExpected)
	// }

	for _, scenary := range testScenarios {
		recivedScenary := address.KindOfAddress(scenary.insertedAddress)

		if recivedScenary != scenary.expectedReturn {
			t.Errorf("Type is incorrect Recived [%s] Expected [%s]", recivedScenary, scenary.expectedReturn)
		}
	}
}