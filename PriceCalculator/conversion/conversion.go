package conversion

import (
	"errors"
	"strconv"
)

// StringsToFloats converts a slice of strings to a slice of floats
func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("error converting string to float")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
