package myutil

import (
	"fmt"
	"testing"
)

type rectTestData struct {
	l     uint
	w     uint
	area  uint
	perim uint
}

var reactTestValues = []rectTestData{
	{2, 3, 6, 10}, {5, 10, 50, 30}, {4, 2, 8, 12}}

func TestRect(t *testing.T) {
	for _, el := range reactTestValues {
		tempArea := Area(el.l, el.w)
		tempPerim := Perim(el.l, el.w)

		if tempArea != el.area {
			t.Errorf("[Test Name: %s * Input %d, %d * Expected: %d * Got: %d]\n",
				t.Name(), el.l, el.w, el.area, tempArea)
		}

		if tempPerim != el.perim {
			t.Errorf("[Test Name: %s * Input %d, %d * Expected: %d * Got: %d]\n",
				t.Name(), el.l, el.w, el.perim, tempPerim)
		}
	}

	fmt.Printf("** Test Rect - ALL PASSED (Number of test cases: %d)\n", len(reactTestValues))
}

// ======================================================================================== //

type avgTestData struct {
	val []float32
	avg float32
}

var averageTestValues = []avgTestData {
	{[]float32{4, 6}, 5.00},
	{[]float32{3, -3}, 0},
	{[]float32{3,-3,5,-5,4}, 0.8},
	{[]float32{1,1,1,100}, 25.75},
}

func TestAverage(t *testing.T) {
	for _, el := range averageTestValues {
		tempAvg := Average(el.val)

		if tempAvg != el.avg {
			t.Errorf("[Test Name: %s * Input %.2f, Expected: %.2f * Got: %.2f]\n",
				t.Name(), el.val, el.avg, tempAvg)
		}
	}

	fmt.Printf("** Test Rect - ALL PASSED (Number of test cases: %d)\n", len(averageTestValues))
}