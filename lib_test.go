package converterlib

import (
	"testing"
)

func TestStringToFloat64(t *testing.T) {

	myFloat := StringToFloat64("33.65")
	t.Log("StringToFloat64(\"33.65\") = ", myFloat)
	t.Errorf("Error: Testing StringToFloat64")
}

func TestFloat64ToString2DecimalPlaces(t *testing.T) {

	myFloat := Float64ToString2DecimalPlaces(33.65676)
	t.Log("Float64ToString2DecimalPlaces(33.65676) = ", myFloat)
	t.Errorf("Error: Testing Float64ToString2DecimalPlaces")
}
