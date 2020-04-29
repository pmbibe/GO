package main

import "fmt"
import "math"
type hChuNhat struct{
	chieuDai float32
	chieuRong float32
}
type hTron struct{
	banKinh float32
}

type Shape interface {
	chuvi() float32
}
func main()  {
	hCN := hChuNhat{10,2}
	fmt.Println(totalArea(&hCN))
}


func chuviHCN(c *hChuNhat) float32  {
	return (c.chieuDai + c.chieuRong)*2
}
func chuviHT(c *hTron) float32{
	return math.Pi * c.banKinh
}

func dientichHCN(c *hChuNhat) float32{
	return c.chieuDai * c.chieuRong
}

func dientichHT(c *hTron)  float32{
	return math.Pi * c.banKinh
	
}

func (c *hChuNhat) chuvi() float32 {
	return (c.chieuDai + c.chieuRong)*2
}
func (c *hTron) chuvi() float32 {
	return math.Pi * c.banKinh
}
func totalArea(shapes Shape) float32 {
    return shapes.chuvi()
}
