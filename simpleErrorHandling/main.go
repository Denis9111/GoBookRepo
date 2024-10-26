// обработка ошибок на примерe проги расчета кол-ва краски
// возвращаемое значение всегда следует проверять, чтобы знать, произошла ли ошибка
package main

import (
	"fmt"
)

func main() {

	amount, err := paintNeeded(10, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

}

func paintNeeded(width float64, height float64) (float64, error) {

	if width <= 0 {
		return 0, fmt.Errorf("width must be greater than zero, width: %v", width)
	} else if height <= 0 {
		return 0, fmt.Errorf("height must be greater than zero, height: %v", height)
	}

	area := width * height
	return area / 10.0, nil
}
