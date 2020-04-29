package main
 
import "fmt"
 
type point struct {
    x, y int
}
 
func main() {
    p := point{1, 2}
 
    fmt.Printf("%v\n", p)  //in các giá trị của một biến struct
    fmt.Printf("%+v\n", p) //in các giá trị kèm với tên trường của biến struct
    fmt.Printf("%#v\n", p) // giống %+v kèm theo tên kiểu dữ liệu của struct và tên hàm đã gọi nó
    fmt.Printf("%T\n", p) // in tên struct và tên hàm đã gọi nó
 
    fmt.Printf("%t\n", true) //in giá trị boolean
 
    fmt.Printf("%d\n", 123) //in số nguyên (hệ 10)
    fmt.Printf("%b\n", 14)  //in số nguyên dưới dạng số nhị phân (hệ 2)
    fmt.Printf("%c\n", 33) // in kí tự dựa theo mã ASCII
    fmt.Printf("%x\n", 456) // in số nguyên dưới dạng số thập lục phân (hệ 16) hoặc chuyển một chuỗi thành số thập lục phân
 
    fmt.Printf("%f\n", 78.9) //in số thập phân
    fmt.Printf("%e\n", 123400000.0) // in số thập phân dưới dạng số mũ
    fmt.Printf("%E\n", 123400000.0) // in số thập phân dưới dạng số mũ
 
    fmt.Printf("%s\n", "\"string\"") // in một chuỗi
    fmt.Printf("%q\n", "\"string\"") //in một chuỗi có 2 cặp dấu nháy kép “”
 
    fmt.Printf("%x\n", "hex this") // 
    fmt.Printf("%p\n", &p) //
    fmt.Printf("|%6d|%6d|", 12, 345) // in một số nguyên, nếu số đó không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //in số thập phân, làm tròn đến 2 chữ số thập phân, sau đó nếu phần thập phân và phần nguyên cùng với dấu chấm không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
    fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45) // t tương tự với %6.2f nhưng các khoảng trống được thêm vào bên phải
    fmt.Printf("|%6s|%6s|\n", "foo", "b") // in một chuỗi, nếu chuỗi không đủ 6 kí tự thì thêm các khoảng trống vào bên trái cho đủ
    fmt.Printf("|%-6s|%-6s|", "foo", "b") // tương tự %6s nhưng thêm các khoảng trống vào bên phải
  
    s := fmt.Sprintf("a %s", "string") //
    fmt.Println(s)
}
