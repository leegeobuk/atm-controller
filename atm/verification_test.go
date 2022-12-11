package atm

import (
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
)

func TestATM_verifyCardNumber(t *testing.T) {
	// given
	const iter = 3
	simpleBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	simpleATM := New[int](simpleBank, cashBin)

	sb := strings.Builder{}
	for i := 0; i < iter; i++ {
		sb.WriteByte('1')
		sb.WriteByte('\n')
	}

	invalidCardNumber := strings.NewReader(sb.String())
	validCardNumber := strings.NewReader("1234123412341234\n")

	// when
	_, invalid := simpleATM.verifyCardNumber(invalidCardNumber, iter)
	_, valid := simpleATM.verifyCardNumber(validCardNumber, iter)

	// then
	if invalid {
		t.Errorf("verifyCardNumber(%s, %d) = %t, want %t", "invalidCardNumber", iter, invalid, false)
	}

	if !valid {
		t.Errorf("verifyCardNumber(%s, %d) = %t, want %t", "validCardNumber", iter, valid, true)
	}
}

func TestATM_verifyPIN(t *testing.T) {
	// given
	const iter = 3
	simpleBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	simpleATM := New[int](simpleBank, cashBin)

	sb := strings.Builder{}
	for i := 0; i < iter; i++ {
		sb.WriteByte('1')
		sb.WriteByte('\n')
	}

	invalidPIN := strings.NewReader(sb.String())
	validPIN := strings.NewReader("1234\n")

	// when
	_, invalid := simpleATM.verifyPIN(invalidPIN, iter)
	_, valid := simpleATM.verifyPIN(validPIN, iter)

	// then
	if invalid {
		t.Errorf("verifyPIN(%s, %d) = %t, want %t", "invalidPIN", iter, invalid, false)
	}

	if !valid {
		t.Errorf("verifyPIN(%s, %d) = %t, want %t", "validPIN", iter, valid, true)
	}
}
