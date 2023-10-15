package main
import "fmt"

func main() {
	s:= make([]int, 3)
	fmt.Println("emp:",s)

	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("set:",s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, 4)
	s = append(s, 5, 6)
	fmt.Println("apd:", s)
	//fmt.Println(type(s))
	var i int


	for  i = 0;i<len(s);i++{
		append(s, i)
		fmt.Println("apd:",s)
	}
	fmt.Println()



}