package main

import (
	"fmt"
	"gostudy/mylib"
	"gostudy/mylib/under"
	"io"
	"log"
	"os"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

// func add(x, y int) (int, string) {
// 	return x + y, strconv.Itoa(x + y)
// }
// func foo(params ...int) {
// 	fmt.Println(len(params), params)
// 	for _, param := range params {
// 		fmt.Println(param)
// 	}
// }

// func getOsName() string {
// 	return "mac"
// }

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("lesson.log")
	// 	var i int = 1
	// 	var f64 float64 = 1.2
	// 	var s string = "test"
	// 	var t, f bool = true, false
	// 	fmt.Println(i, f64, s, t, f)

	// 	xi := 1
	// 	xf := 1.2
	// 	x := "test"
	// 	xt := true
	// 	fmt.Println(xi, xf, x, xt)
	// 	fmt.Println(string(x[0]))
	// 	fmt.Println(strings.Replace(x, "t", "s", 1))
	// 	fmt.Println(x)
	// 	x = "2"
	// 	tx, fx := true, false
	// 	fmt.Printf("%T %v", tx, fx)
	// 	n := []int{1, 2, 3}
	// 	fmt.Println(n)
	// 	var nx [2]int
	// 	nx[0] = 1
	// 	nx[1] = 2
	// 	fmt.Println(nx)
	// 	fmt.Println(n)
	// 	n = append(n, 4, 5)
	// 	m := map[string]int{"apple": 100}
	// 	fmt.Println(m)
	// 	fmt.Println(n)
	// 	fmt.Println(add(1, 2))
	// 	foo(1, 2, 3)
	// 	ss := []int{1, 2, 3, 4}
	// 	foo(ss...)

	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	sum := 1
	// 	for sum < 3 {
	// 		sum += sum
	// 		fmt.Println(sum, "sum")
	// 	}
	// 	l := []string{"python", "go", "java"}
	// 	for i, v := range l {
	// 		fmt.Println(i, v)
	// 	}
	// 	mm := map[string]int{"apple": 100, "banana": 200}
	// 	for k, v := range mm {
	// 		fmt.Println(k, v)
	// 	}
	// 	osd := "mac"
	// 	switch osd {
	// 	case "mac":
	// 		fmt.Println("mac")
	// 	case "window":
	// 		fmt.Println("window")
	// 	}

	// 	switch osd := getOsName(); osd {
	// 	case "mac":
	// 		fmt.Println("mac")
	// 	default:
	// 		fmt.Println(osd)
	// 	}
	// 	log.Println("logging")
	// 	log.Printf("%T %v", "test", "test")
	// 	file, err := os.Open("khjh")
	// 	if err != nil {
	// 		log.Println("file not found")
	// 	}
	// 	defer file.Close()
	// 	data := make([]byte, 100)
	// 	count, err := file.Read(data)
	// 	if err != nil {
	// 		log.Println("Error")
	// 	} else {
	// 		fmt.Println(count, string(data))
	// 	}
	// 	if err = os.Chdir("test"); err != nil {
	// 		log.Println("error")
	// 	}
	// 	ll := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	// 	var min int
	// 	for i, v := range ll {
	// 		if i == 0 {
	// 			min = v
	// 		} else if v < min {
	// 			min = v
	// 		}
	// 	}
	// 	fmt.Println(min)

	// 	mmm := map[string]int{
	// 		"apple":  200,
	// 		"banana": 300,
	// 		"grapes": 150,
	// 		"orange": 80,
	// 		"papaya": 500,
	// 		"kiwi":   90,
	// 	}
	// 	totalCost := 0
	// 	for _, v := range mmm {
	// 		totalCost += v
	// 	}
	// 	fmt.Println(totalCost)

	// 	var p *int = new(int)
	// 	fmt.Println(p)
	// 	fmt.Println(*p)

	// 	v := Vertex{X: 1, Y: 2}

	// 	v2 := &Vertex{}
	// 	fmt.Println(v, v.X, v.Y)
	// 	fmt.Printf("%T %v\n", v2, v2)
	// 	v3 := &Vertex{1, 2, 3}
	// 	changeVertex(v3)
	// 	fmt.Println(*v3)
	// var mike Human = &Person{"Mike"}
	// var x Human = &Person{"x"}
	// DriveCar(mike)
	// DriveCar(x)

	// do(10)
	// do("Mike")
	// v := Vertex{3, 4}
	// fmt.Println(v.Plus())
	// v := Vertex{3, 4}
	// fmt.Println(v)
	// s := []int{1, 2, 3, 4}
	// c := make(chan int)
	// go goroutine1(s, c)
	// for i := range c {
	// 	fmt.Println(i)
	// }
	// var wg sync.WaitGroup
	// ch := make(chan int)

	// // producer
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go goroutine(ch, i)
	// }
	// // consumer
	// go consumer(ch, &wg)
	// wg.Wait()
	// close(ch)
	// first := make(chan int)
	// second := make(chan int)
	// third := make(chan int)

	// go producer(first)
	// go multi2(first, second)
	// go multi4(second, third)

	// for result := range third {
	// 	fmt.Println(result)
	// }
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

}

// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first <-chan int, second chan<- int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// func multi4(second <-chan int, third chan<- int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func goroutine(ch chan int, i int) {
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		fmt.Println("process", i*1000)
// 		wg.Done()
// 	}
// }

// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)

// }

// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) String() string {
// 	return fmt.Sprintf("X is %v! Y is %v", v.X, v.Y)
// }

// func (p Vertex) Plus() int {
// 	return p.X + p.Y
// }

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "!")
// 	default:
// 		fmt.Printf("i dont know %T\n", v)
// 	}
// }

// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println("RUN")
// 	} else {
// 		fmt.Println("Get Out")
// 	}
// }

// type Vertex struct {
// 	X    int
// 	Y, Z int
// }

// func changeVertex(v *Vertex) {
// 	v.X = 1000
// }
