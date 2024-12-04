package app

import (
	"flag"
)

type Flags struct {
	MeanFlag   *bool
	MedianFlag *bool
	ModeFlag   *bool
	SdFlag     *bool
}

func (f *Flags) Parse() {
	flag.Parse()
}

func FlagsInit() *Flags {
	// Создание флагов (функия flag.Bool) возвращает указатель типа *bool
	return &Flags{
		flag.Bool("mean", false, "вывести только среднее значение"),
		flag.Bool("median", false, "вывести только медиану"),
		flag.Bool("mode", false, "вывести только моду"),
		flag.Bool("sd", false, "вывести только стандартное отклонение"),
	}

}
