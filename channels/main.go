package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tjarratt "github.com/tjarratt/babble"
)

var waitG sync.WaitGroup
var cpuUsed = 1
var maxRandomNums = 10
var counter uint64
var m sync.Mutex

func init() {
	maxCPU := runtime.NumCPU()

	cpuUsed = 4
	runtime.GOMAXPROCS(cpuUsed)

	fmt.Printf("Number of CPUs (Total=%d - Used=%d)\n", maxCPU, cpuUsed)
}

func main() {
	// ex1()
	// manageSync()
	// ex3()
	// multiplex()
	// bufferedChan()
	// ex7()
	ex7pipeline()
}

func ex7pipeline() {
	newWords := make(chan string)
	uWords := make(chan string)

	go sendWords(newWords)
	go convertWords(uWords, newWords)
	printWords(uWords)
}

func sendWords(out chan<- string) {
	babbler := tjarratt.NewBabbler()
	for i := 0; i < 5; i++ {
		out <- babbler.Babble()
	}
	close(out)
}

func convertWords(out chan<- string, in <-chan string) {
	for w := range in {
		out <- w + " --> " + strings.ToUpper(w)
	}
}

func printWords(in <-chan string) {
	for w := range in {
		fmt.Println(w)
	}
}

func ex7() {
	babbler := tjarratt.NewBabbler()
	newWord := make(chan string)
	uWords := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			newWord <- babbler.Babble()
		}
		close(newWord)
	}()

	go func() {
		for w := range newWord {
			uWords <- w + " --> " +strings.ToUpper(w)
		}
		close(uWords)
	}()

	func() {
		for w := range uWords {
			fmt.Println(w)
		}
	}()
}

func bufferedChan() {
	c := make(chan string, 3)
	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 1"
	c <- "Message 2"

	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 3"
	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))

	nums := []int{2, 3, 4}
	c2 := gen(nums...)
	out := sq(c2)

	for range nums {
		fmt.Printf("%4d ", <-out)
	}

	fmt.Println()
	for n := range sq(sq(gen(nums...))) {
		fmt.Printf("%4d ", n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func multiplex() {
	stats := make(map[string]int)

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i := 0; i < 20; i++ {
		go func() {
			time.Sleep(10 * time.Second)
			c1 <- "Hell0 from customer service #1"
		}()

		go func() {
			time.Sleep(8 * time.Second)
			c2 <- "Hell0 from customer service #2"
		}()

		go func() {
			time.Sleep(1 * time.Second)
			c3 <- "Hell0 from customer service #3"
		}()

		select {
		case msg1 := <-c1:
			stats["Nick"]++
			time.Sleep(time.Second)
			fmt.Println(msg1)
		case msg2 := <-c2:
			stats["Tim"]++
			time.Sleep(time.Second)
			fmt.Println(msg2)
		case msg3 := <-c3:
			stats["Viktor"]++
			time.Sleep(time.Second)
			fmt.Println(msg3)
		default:
			stats["Customer Waiting"]++
			time.Sleep(6 * time.Second)
			fmt.Println("No customer service is available at this time!")
		}
	}

	fmt.Printf("\n ***Customer service stats\n%v", stats)

	close(c1)
	close(c2)
	close(c3)
}

func ex6() {
	text := "please Modify Me!"

	cu := make(chan string)
	cl := make(chan string)

	fmt.Println("(m) before uppercase()")
	go uppercase(text, cu)
	fmt.Println("(m) before lovercase()")
	go lowercase(text, cl)

	sUpper, sLower := <-cu, <-cl

	fmt.Printf("sUpper=%s sLower=%s\n", sUpper, sLower)
}

func uppercase(s string, c chan string) {
	c <- strings.ToUpper(s)
}

func lowercase(s string, c chan string) {
	c <- strings.ToLower(s)
}

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "yellow"}

var b = make(chan bool)
var c = make(chan string)

func ex5() {
	f := []interface{}{sender1, sender2}

	for i := range f {
		go f[i].(func())()
	}
	go closeSenders()

	fmt.Println("before range")
	for val := range c {
		fmt.Println(val)
	}
}

func sender1() {
	for _, w := range wordSet1 {
		c <- w
	}
	b <- true
}

func sender2() {
	for _, w := range wordSet2 {
		c <- w
	}
	b <- true
}

func closeSenders() {
	<-b
	<-b
	close(c)
}

func ex4() {
	c := make(chan int)

	waitG.Add(2)
	go send(c)
	go receive(c)
	waitG.Wait()
}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for i := 1; i < 6; i++ {
		fmt.Println(<-c)
	}
	waitG.Done()
}

func ex3() {
	var c = make(chan string)

	start := time.Now()

	go message(c, "msg1", 3)
	go message(c, "msg2", 2)
	go message(c, "msg3", 4)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)
	elapsed := time.Since(start)
	fmt.Printf("\n time elapsed: %s\n", elapsed)
}

func message(c chan string, s string, d time.Duration) {
	time.Sleep(d * time.Second)

	c <- s
}

func ex2() {
	start := time.Now()
	ids := []string{"#", "@"}

	waitG.Add(len(ids)) // 4
	for i := range ids {
		go numbers(i+1, ids[i])
	}
	waitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nprogram took %s. \n", elapsed)
	fmt.Printf("\ncounter %d. \n", counter)
}

func manageSync() {
	waitG.Add(2)
	// go numbers()
	go alphabets()
	waitG.Wait()

	fmt.Println("\n main terminated")
}

func numbers(n int, id string) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= maxRandomNums; i++ {

		time.Sleep(200 * time.Millisecond)

		// m.Lock()
		// counter++
		atomic.AddUint64(&counter, 1)
		// fmt.Printf("(%d) %d %d\n", n, rand.Intn(20)+20, counter) // formula: rand.Intn(max-min)+min
		fmt.Printf("(%d) %d %d\n", n, rand.Intn(20)+20, atomic.LoadUint64(&counter)) // formula: rand.Intn(max-min)+min
		// m.Unlock()
	}
	waitG.Done()
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	waitG.Done()
}

func ex1() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !")
		c <- link
		return
	}
	fmt.Println(link, "is up !")
	c <- link
}
