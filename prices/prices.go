package prices

import (
	"fmt"

	"example.com/tax-calculator/conversions"
	"example.com/tax-calculator/iomanager"
)

type TaxIncludedPrice struct {
	InputPrices []float64 `json:"input_prices"`
	TaxRate float64 `json:"tax_rate"`
	TaxIncludedPrices map[string]string `json:"tax-included-prices"`
	FileManager iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPrice(io iomanager.IOManager,taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		FileManager: io,
		TaxRate: taxRate,
	}
}

func (processor *TaxIncludedPrice) Process() error {
	err := processor.readPricesFromFile()

	if err != nil {
		return err
	}

	result := make(map[string] string)

	for _, price := range processor.InputPrices {
		result[fmt.Sprintf("%.2f ", price)] = fmt.Sprintf("%.2f", price * (1 + processor.TaxRate))
	}
	processor.TaxIncludedPrices = result
	return processor.FileManager.WriteToJson(processor)
}

func (processor *TaxIncludedPrice) readPricesFromFile() error{
	
	lines, err := processor.FileManager.ReadLinesFromFile()

	if err != nil {
		return err
	}

	inputPrices, err := conversions.StringToFloatConverter(lines)

	if err != nil {
		return err
	}

	processor.InputPrices = inputPrices

	return nil
}