/*package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const {
	AbsoluteZeroC Celsius = -273.15 //绝对零度(摄氏度)
	FreezingC Celsius = 0.0
	BoilingC Celsius = 100.0
}

func (c Celsius) String() string { return fmt.Sprintf("%g°C\n"),c}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F\n",f)}


package tempconv


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}

func FToC(f Fahrenheit) Celsius { return Celsius((f -32) * 5 /9}



package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/xieydd/Go_Zh/tempconv"
)



func main() {
	_, flag := range os.Args[1:] {
		if flag == "" {
			fmt.Println("You need a Celsius or Fahrenheit")
		}

		t, err := strconv.ParseInt(flag,64)
		if err != nil {
			fmt.Fprint(os.Stderr,"parse error : %s \n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %S\n",
			f,tempconv.FToC(f),c,tempconv.CToF(c))
	}
}*/

/*
位运算符
A 60 二进制 0011 1100
B 13 二进制 0000 1101

& 二进位相与 0000 1100 一起为1为1
| 二进为相或 0011 1101 为1就填1
^ 二进位相异或 0011 0001 不同为1
<< n 二进位左移2位 A 左移变成 1111 0000
>> n 二进制右移2位 B右移变成会出错
*/

package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	a := PopCount(80000)
	fmt.Println(a)
}

/*var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i%1)
	}
	return
}()

func PopCount(x uint64) int {
 for i := 0,i<8,i++ {
 	pc += pc[byte(x>>(i*8))]
 }
}*/
