package main

import "C"
import (
	"fmt"
	"reflect"
)

func main() {
	//args := os.Args
	//
	//if args == nil || l(args) < 3 {
	//	Usage()
	//	return
	//}
	//
	//switch args[1] {
	//case "add":
	//	if l(args) < 4 {
	//		fmt.Println("Usage: calc add <int> <int>")
	//		return
	//	}
	//	v1, err1 := strconv.Atoi(args[2])
	//	v2, err2 := strconv.Atoi(args[3])
	//	if err1 != nil || err2 != nil {
	//		fmt.Println("Usage: calc add <int> <int>")
	//		return
	//	}
	//	ret := simplemath.Add(v1, v2)
	//	fmt.Println("Result: ", ret)
	//case "sqrt":
	//	if l(args) < 3 {
	//		fmt.Println("Usage: calc sqrt <int>")
	//		return
	//	}
	//	v, err := strconv.Atoi(args[2])
	//	if err != nil {
	//		fmt.Println("Usage: calc sqrt <int>")
	//		return
	//	}
	//	ret := simplemath.Sqrt(v)
	//	fmt.Println("Result: ", ret)
	//default:
	//	Usage()
	//}
	/*
	   	var i1 int
	   	var s1 string
	   	var b1 bool
	   	var a1 [5]int
	   	var a2 []int
	   	var st1 struct {
	   		f float64
	   		s string
	   	}
	   	var p1 *int
	   	var mp1 map[string]string
	   	var fun1 func(a int) string

	   	i1 = 111
	   	s1 = "string s1"
	   	b1 = true
	   	a1 = [5]int{1, 2, 3, 4, 5}
	   	a2 = []int{9, 8, 7}
	   	st1 = struct {
	   		f float64
	   		s string
	   	}{f: 5.67, s: "aox"}
	   	p1 = &i1
	   	mp1 = map[string]string{
	   		"key1": "value1",
	   		"key2": "value2",
	   	}
	   	fun1 = func(a int) string {
	   		return strconv.Itoa(a)
	   	}

	   	fmt.Printf("i1:%d\ns1:%s\nb1:%t\np1:%d\nfun1(99):%s\n", i1, s1, b1, *p1, fun1(99))
	   	fmt.Println("a1:", a1)
	   	fmt.Println("a2:", a2)
	   	fmt.Println("st1:", st1)
	   	fmt.Println("mp1:", mp1)

	   	const (
	   		c0 = iota
	   		c1 = iota + c0
	   		c2
	   		c3
	   	)
	   	fmt.Println(c1, c2, c3)

	   	const f1 float64 = 9.88
	   	f2 := f1
	   	i2 := int(f2)
	   	fmt.Println(i2)

	   	a3 := [3]int{1, 2, 3}
	   	a4 := a3[0:2]
	   	fmt.Println(reflect.TypeOf(a3))
	   	fmt.Println(reflect.TypeOf(a4) == reflect.TypeOf(a3))
	   	fmt.Println(a4)

	   	//for i := 0; i < l(a1); i++ {
	   	//	fmt.Println(a1[i])
	   	//}
	   	//
	   	//a1[1] = 99
	   	//for i, v := range a1 {
	   	//	fmt.Println(i, v)
	   	//}

	   	//for i, v := range s1 {
	   	//	fmt.Println(i, string(v))
	   	//}
	   	var carr []byte = []byte(s1)
	   	carr[1] = 'p'
	   	s1 = string(carr)
	   	fmt.Println(s1)

	   	a2 = append(a2, 6)
	   	fmt.Println(len(a2), cap(a2), a2, a4)
	   	copy(a4, a2)
	   	fmt.Println(len(a2), cap(a2), a2, a4)
	   	fmt.Println(len(a2), cap(a2), a2, a4)

	   	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	   	//slice4 := slice3[:copy(slice3, slice3[3:])] // 删除开头前三个元素
	   	l := copy(slice3, slice3[3:])
	   	fmt.Println(slice3)
	   	slice4 := slice3[:l]
	   	fmt.Println("l:", l, slice4)

	   	testMap := map[string]int{
	   		"one":   1,
	   		"two":   2,
	   		"three": 3,
	   	}
	   	//testMap := make(map[string]int)
	   	testMap["four"] = 4
	   	fmt.Println(testMap)
	   	str := "four1"
	   	_, _ = testMap[str]
	   	delete(testMap, "two")

	   	//for key, _ := range testMap {
	   	//	fmt.Println(key)
	   	//}

	   	//for m, n := 1, 11; m < 10; m, n = m+1, n+1 {
	   	//	fmt.Println(m, n)
	   	//}

	   	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	   ITERATOR1:
	   	for i := 0; i < 3; i++ {
	   	ITERATOR2:
	   		for j := 0; j < 3; j++ {
	   			num := arr[i][j]
	   			if j > 0 {
	   				break ITERATOR1
	   			} else {
	   				break ITERATOR2
	   			}
	   			fmt.Println(num)
	   		}
	   	}

	   	type Student struct {
	   		id    int
	   		name  string
	   		grade string
	   		male  bool
	   	}

	   	var pStu *Student = new(Student)
	   	//var stu []Student = make([]Student, 2)

	   	fmt.Println(pStu)
	*/

	//arrStu := make([]Student, 0)
	//
	//fmt.Println(arrStu)
	//arrStu = append(arrStu, Student{1, "Tom"})
	//fmt.Println(arrStu)

	//stu := Student{1, "Tom"}
	//
	//fmt.Println(stu.getName())

	//person := make([]People, 0)
	//person = append(person, new(Student))
	//person = append(person, new(Teacher))
	//
	//for _, v := range person {
	//	v.Speak()
	//}

	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//resultChan := make(chan int, 2)
	//go sum(arr[:5], resultChan)
	//go sum(arr[5:], resultChan)
	//sum1, sum2 := <-resultChan, <-resultChan
	//fmt.Println(sum1, sum2)

	stu := &Student{1, false, 10, "Tom"}
	fmt.Println(stu, reflect.TypeOf(stu))
	s := reflect.ValueOf(stu).Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}

	//cStr := C.CString("Hello, World")
	////fmt.Println(cstr)
	//C.printf("Hello world")
	//C.free(unsafe.Pointer(cStr))
}

type People interface {
	Speak()
}

type Student struct {
	Id   int
	Male bool
	Age  int
	Name string
}

func (p *Student) Speak() {
	fmt.Println("I'm a student")
}

type Teacher struct {
}

func (p *Teacher) Speak() {
	fmt.Println("I'm a teacher")
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	resultChan <- sum
}

//type Student struct {
//	id   int
//	name string
//}
//
//func (p *Student) getName() string {
//	//return p.name
//	return "test name"
//}

func Usage() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}
