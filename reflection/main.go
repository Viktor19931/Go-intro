package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {
	testCase4()
}

func testCase1() {
	t := reflect.TypeOf(32)
	fmt.Println(t)

	fmt.Println(reflect.TypeOf(10.23))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf([...]int{1, 2}))
	fmt.Println(reflect.TypeOf([]int{1, 2}))

	var f io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(f))
}

func testCase2() {
	i := 10
	iv := reflect.ValueOf(i)

	fmt.Println(iv)
	fmt.Println(iv.String())

	t := iv.Type()
	fmt.Println(t.String())
}

func testCase3() {
	f := 10.32

	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)

	fmt.Printf("fv=%v, fv.Kind()=%v, fv.Type()=%v \n", fv, fv.Kind(), fv.Type())
	fmt.Printf("fv.Strinr()=%v fv.Float()=%v fv.Interface()=%v fv.Interface.(float64)=%v\n",
		fv.String(), fv.Float(), fv.Interface(), fv.Interface().(float64))
	fmt.Printf("ft=%v, ft.Kind()=%v fv.String()=%v \n", ft, ft.Kind(), ft.String())
}

func testCase4() {
	type myFloat float64
	var f myFloat = 10.34
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)

	fmt.Printf("fv=%v, fv.Kind()=%v, fv.Type()=%v \n", fv, fv.Kind(), fv.Type())
	fmt.Printf("fv.String()=%v fv.Float()=%v fv.Interface=%v \n", fv.String(), fv.Float(), fv.Interface())
	fmt.Printf("ft=%v, ft.Kind()=%v, ft.String()=%v \n", ft, ft.Kind(), ft.String())
}

func getType(value interface{}) string {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int32, reflect.Int64: // and rest
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8: // and rest
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 3, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	default:
		return v.Type().String() + " value"
	}
}
