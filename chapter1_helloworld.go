package main

/*const PI = 3.14

var name = "xyd"
//var b bool = true

var isActive bool

var enable,disable = true,true*/

/*func test() {
	var avaliable bool
	valid := false
	avaliable = true
}*/

/*var a = "谢远东"
var b string = "张淑贞"

type newType int

type gopher struct {}

type golang interface{}*/

/*
var (
	a int
	b bool
)

var c,d int = 1,2
var e,f = 123 , "hello"

func main() {

    fmt.Println(a,b,c,d,f)

    var g int = 8
    var h int = 10
    var i int
    i = g + h
    var ptr *int
    i <<= 2
    ptr = &i
    fmt.Printf("<<=结果为%d\n",*ptr)
}

//私有的方法
func getId() {}

//公有的方法
func Printf(){}*/

/*func main() {

var count int
var flag bool

count = 1

for count < 100 {
	count ++
	flag = true

	for tmp := 2; tmp < count; tmp++ {
		if count%tmp == 0 {
			flag = false
		}
	}
	if flag == true {
		fmt.Println(count, "素数")
	} else {
		continue
	}

}*/

/*func main() {
var i interface{}

switch j := i.(type) {
case nil:
	fmt.Println("i的类型是", j)
case int:
	fmt.Printf("i的类型是int")
}*/

/*func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d  ", j, i, j*i)
		}
		fmt.Println(" ")
	}
}*/

/*func max(num1,num2 int) int {
	var max int

	if (num1 > num2) {
		max = num1
	} else if num2 > num1 {
		max = num2
	}else {
		max = num1
	}
	return max
}

func swap(x,y string) (string,string) {
	var z string
	z = x
	x = y
	y = z
	return x,y

}

func main() {
	max := max(2,3)
	fmt.Println(max)

	a,b := swap("kendric","larmer")

	fmt.Printf(a,b)
}*/

/*func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)
}*/

/*
func getAverage(arr []int,size int) float32{
	var average float32
	var sum int

	for i:= 0;i<size;i++ {
		sum += arr[i]
	}
	average = float32(sum) / float32(size)
	return average
}

func main() {

	var arr = []int {1000, 2, 3, 17, 50}

	average := getAverage(arr,5)
	fmt.Println(average)
}*/

/*func main() {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a的变量地址是: %d\n",&a)

	fmt.Printf("a变量存储的指针的地址是: %x\n",ip)

	fmt.Printf("ip变量的值是: %d\n",*ip)


}*/

/*
func main() {
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	numbers1 := numbers[2:5]
	printSlice(numbers1)

	numbers2 := make([]int,3,4)
	printSlice(numbers2)

	numbers3 := append(numbers, 9)
	printSlice(numbers3)
}

func printSlice(x []int)  {
	fmt.Printf("len = %d , cap = %d , slcie = %v \n",len(x),cap(x),x)
}*/

/*func main() {
	for i,c := range "go" {
		fmt.Println(i,c)
	}

	countryMap := make(map[string]string)

	countryMap["china"] = "北京"
	delete(countryMap,"china")
	for country := range countryMap {
		fmt.Println(country,"首都是",countryMap[country])
	}
}*/

/*func Factorial(n uint64)(result uint64) {
	if (n>0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int{
	if(n<2) {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}


func main() {
	var n uint64 = 15
	var m int = 15
	result := Factorial(n)
	fmt.Printf("%d的阶乘是%d",n,result)
	fmt.Printf("%d的斐波那契数列值为%d",n,fibonacci(m))
}*/

/*type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat,de.dividee)
}

func Divide(varDividee int,varDivider int) (result int,errorMsg string) {
	if varDivider == 0  {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}

		errorMsg = dData.Error()
		return
	}else {
		return varDividee / varDivider,""
	}
}

func main() {
	if result,errorMsg := Divide(100,10);errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}*/
/*

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var s,seq string

	for i := 1;i < len(os.Args);i++ {
		s += os.Args[i] + seq
		seq = " "
	}
	fmt.Println(s)

	s1,seq1 := "",""
	for  _,arg := range os.Args[1:] {
		s1 += seq1 + arg
		seq1 = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:]," "))
}*/

/*
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}*/

/*import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts{
		if n > 1{
		fmt.Printf("%d\t%s\n", n, line)
	}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}*/

/*import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}*/

/*import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s :%v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}*/

/*import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s :%v\n", url, err)
			os.Exit(1)
		}
	}
}*/

/*import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s :%v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}*/

/*import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("http-get : %v\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s :%v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}*/

/*import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//	"strings"
	"time"
)

func main() {

	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //从channel中接收
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	//if !strings.HasPrefix(url, "http://") {
	//	url = "http://" + url
	//}

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //给channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}*/

/*import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	//	"strings"
	"time"
)

func main() {

	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //从channel中接收
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	//if !strings.HasPrefix(url, "http://") {
	//	url = "http://" + url
	//}

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //给channel
		return
	}

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch : reading %s :%v\n", url, err)
		os.Exit(1)
	}

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("  %.2f   %s", secs, url)
}*/

/*import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}*/

/*import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}*/

/*import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
	}

	fmt.Fprintf(w, "%q \n", r.Host)
	fmt.Fprintf(w, "%q \n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q \n", k, v)
	}

}*/

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.RGBA{255, 255, 255, math.MaxUint8}, color.Black, color.RGBA{0, 0, 128, math.MaxUint8}, color.RGBA{30, 144, 255, math.MaxUint8}, color.RGBA{72, 61, 139, math.MaxUint8}, color.RGBA{255, 0, 255, math.MaxUint8}, color.RGBA{124, 252, 0, math.MaxUint8}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2 //新定义的颜色
)

func lissajous(out http.ResponseWriter, cycles float64) {
	const (
		//cycles  = 5     //全x振子转速
		res    = 0.001 //角分辨率
		size   = 1000  //图像覆盖像素
		nframe = 64    //动画帧数
		delay  = 8     //10 ms的帧之间的延迟
	)

	freq := rand.Float64() * 3.0 // y振子转动频率
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0 //相位差
	for i, c := 0, 1; i < nframe; i++ {
		rect := image.Rect(0, 0, 0.2*size+1, 0.2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
		}
		c++ //修改颜色

		if c > len(palette) {
			c = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Print(err)
	}

	for k, v := range r.Form {
		if k == "cycles" {

			//返回多个值，不需要的用_接收
			cycles, _ := strconv.ParseFloat(v[0], 64)

			//加判断，防止占用过多的内存
			if cycles < 50 {
				lissajous(w, cycles)
			}
		} else {
			fmt.Fprintf(w, "URL PATH:%q \n", r.URL.Path)
			fmt.Fprintf(w, "%q:%q", k, v)
		}
	}
}
