package converterlib

import (
	"testing"
)

func TestStringToFloat64(t *testing.T) {

	myFloat := StringToFloat64("33.65")
	t.Log("StringToFloat64(\"33.65\") = ", myFloat)
	t.Errorf("Error: Testing StringToFloat64")
}
