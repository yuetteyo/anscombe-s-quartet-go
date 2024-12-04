package app

import (
	"fmt"
)

func PrintStats(f Flags, mean float64, median float64, mode float64, sd float64) {
	if *f.MeanFlag || *f.MedianFlag || *f.ModeFlag || *f.SdFlag {
		if *f.MeanFlag {
			fmt.Printf("Mean %v\n", mean)
		}
		if *f.MedianFlag {
			fmt.Printf("Median %v\n", median)
		}
		if *f.ModeFlag {
			fmt.Printf("Mode %v\n", mode)
		}
		if *f.SdFlag {
			fmt.Printf("SD %v\n", sd)
		}
	} else {
		fmt.Printf("Mean %.2f\n", mean)
		fmt.Printf("Median %.2f\n", median)
		fmt.Printf("Mode %.2f\n", mode)
		fmt.Printf("SD %.2f\n", sd)
	}
}
