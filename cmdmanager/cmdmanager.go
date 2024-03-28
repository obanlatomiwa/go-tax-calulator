package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadLinesFromFile() ([]string, error) {
	fmt.Println("Please enter your prices, HIT ENTER to confirm each price, when done HIT 0")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)

	}
	return prices, nil
}


func (cmd CMDManager) WriteToJson(data interface{}) error{
	fmt.Println(data)
	return nil
}

