package main
import (
    "fmt"
     "time"
)

func f(n int) {
    for i:=1;i<=10;i++ {
        fmt.Println(n, " : ", i)
    }

}
func main(){
    for n:=0; n<=5; n++ {
        go f(n)
    }
    time.Sleep(time.Second)
    // fmt.Println("done")
}
