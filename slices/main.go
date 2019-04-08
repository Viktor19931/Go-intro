package main

import "fmt"

type slice []int

func main() {
	// arrays()
	slices()
}

func slices() {
	days := [...]string{"fds", "fds", "dsfdsf", "dsfsdf", "fdsdf", "fdsfsd", "sdfsdf"} // array

	kk := days[:5] // slicing days[0:4]

	kk = append(kk, "WWW")

	fmt.Printf("slice=%v\n", kk)
	fmt.Printf("length=%d\n", len(kk))
	fmt.Printf("capacity=%d\n", cap(kk))
	// fmt.Printf("kk[0]=%s\n", kk[0])

	k := days[5:] // slicing days[5:7]

	fmt.Printf("slice=%v\n", k)

	s := "ABCDEFGHIJKL"
	slc := s[1:5]
	fmt.Println(s, slc)

	var f []float32

	f = append(f, 1.2)
	f = append(f, 2.4, 4.8, 8.16)
	// fmt.Println(f)
	f = append(f, f...)
	fmt.Println(f)

	team1 := []string{"Messi", "Pele", "Maradonna", goalkeeper()}

	team := append([]string{"Maldini", "Cafu", "Carlos", "Beckenbauer"}, team1...)
	team = append(team, midfielders()...)

	for i, name := range team {
		fmt.Println(i, name)
	}

	nums := make([]int, 2, 3) // or  // nums := new([3]int)[0:2]

	fmt.Println(nums)

	s1 := []int{1, 3, 5}
	s2 := make([]int, 2)

	copy(s2, s1)  // copy s1 to s2 

	s3 := append(s2, 2, 6)

	fmt.Println(s1, s2, s3)
	for i := range s3 {
		fmt.Println(s3[i], " ")
	}

	m2 := []int{2, 3, 4, 9, 2, 3, 4}
	m1 := slice{1, 2, 3, 4, 5}
	slices := [][]int{m2, {5, 6}, {7, 8, 9, 10}}
	var flatSlice []int

	for i := range slices {
		flatSlice = append(flatSlice, slices[i]...)
	}

	fmt.Println(flatSlice)
	fmt.Println(m1)
	m1 = m1.deleteIndex(2)
	fmt.Println(m1)

	// сut
	// copy(a[i:], a[j:])
	// for k, n := len(a)-j+i, len(a); k < n; k++ {
	// 	a[k] = nil // or the zero value of T
	// }
	// a = a[:len(a)-j+i]
}


func (s slice) deleteIndex(i int) slice {
	k := s
	copy(k[i:], k[i+1:])
	k[len(k)-1] = 0 // or the zero value of T
	return k[:len(k)-1]
	
}

func midfielders() []string {
	return []string{"Iniesta", "Zidane", "Platini"}
}

func goalkeeper() string {
	return "DBuffon"
}

func arrays() {
	var nums [3]int

	var sum1 int
	var sum2 int

	// num1 := [...]string{"num1", "num2"} // dinamic array
	// num2 := [3]int{12, -34, 100}        // static array

	x := [3]float32{3.2, 2.54, 3.4}

	y := [...]int{3: 10, 20}  // [0 0 0 10 20]

	a := [2][3]int{{2, 4, 6}, {4, 6, 8}}

	a[0][1] = 5 // 4 --> 5
	a[1][1] = 10 // 6 --> 10

	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(a)

	var total float32

	for _, val :=range x {
		total += val
	}

	fmt.Println(total)

	for i := range nums {
		sum1 += i
		sum2 += nums[i]
	}

	fmt.Printf("nums=%v type=%T len=%d\n ", nums, nums, len(nums))
	fmt.Println(sum1, sum2)
}
