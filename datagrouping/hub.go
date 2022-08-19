package datagrouping

import "strings"

func DataGrouping(op string) {
	switch {
	case strings.Contains(op, "array"):
		callArrays()
	case strings.Contains(op, "slice"):
		callSlices()
	case strings.Contains(op, "map"):
		callMaps()
	case strings.Contains(op, "ex"):
		callExercises()
	}
}

func callArrays() {
	normalArray()
	directDeclarationArray()
	sliceComparingToArray()
}

func callSlices() {
	normalSlice()
	changeSliceItemValue()
	forRangeSlice()
	appendSlice()
	sliceSlice()
	sliceAllItems()
	forAllSliceItems()
	removeSliceItems()
	appendUsage()
	sliceMake()
	multiDimensionalSlice()
	underlyingArray()
}

func callMaps() {
	normalMap()
	commaOk()
	rangeMap()
	deleteMapElement()
}

func callExercises() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
}
