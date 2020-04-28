package main
import "fmt"
func div(a, b float32){
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Println("Error: ", r)
		}
	}()
	if b==0 {
		panic("DIV BY ZERO")
	}
	fmt.Println( a/b)
}
func main()  {
div(10,5)
div(10,0)
div(10,2)


}