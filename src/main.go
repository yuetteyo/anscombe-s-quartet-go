/* Программа вычисляет:
median - если количество чисел нечетно, то медиана это среднее число в отсортированных данных
	   - если количестов чисел четное, то медиана - это среднее между двумя числами посередине
mean   - среднее значение
mode   - это наиболее часто повторяющееся число в последовательности,
	     если таких чисел несколько, то это наименьшее из них
	    (using some structure for storing number counts, and some Go standard container may help.)
SD     - Standard Deviation, что переводится как стандартное отклонение
*/

package main

import (
	"fmt"
	"s21anscombeapp/app"
)

func main() {

	flags := app.FlagsInit()
	flags.Parse()

	fmt.Println("Start typing numbers divided by \"enter\".")
	fmt.Println("Press \"q\" or \"Q\", when you are ready for calculations.")

	numbers := app.Read()

	fmt.Println("Input data:", numbers)

	mode := app.CalculateMode(numbers)
	median := app.CalculateMedian(numbers)
	mean := app.CalculateMean(numbers)
	sd := app.CalculateSd(numbers)

	app.PrintStats(*flags, mean, median, mode, sd)

}
