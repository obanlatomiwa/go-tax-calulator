package conversions

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloatConverter(stringList []string) ([]float64, error) {
	floatValueList := make([]float64, len(stringList))
	for stringIndex, stringItem := range stringList {
		newFloatValue, err := strconv.ParseFloat(stringItem, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Can not convert string value %s to float64 value", stringItem))
		}
		floatValueList[stringIndex] = newFloatValue
	}
	return floatValueList, nil
}
