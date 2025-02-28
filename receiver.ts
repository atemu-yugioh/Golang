// package receiver

// import (
// 	"fmt"
// )

// // Định nghĩa struct Vertex
// type Vertex struct {
// 	X, Y float64
// }

// // Phương thức Scale() sử dụng giá trị receiver (Vertex)
// func (v Vertex) Scale(factor float64) {
// 	v.X *= factor
// 	v.Y *= factor
// }

// func (v *Vertex) Scale(factor float64) {
// 	v.X *= factor
// 	v.Y *= factor
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println("Trước khi Scale():", v) // In ra: {3 4}

// 	v.Scale(2) // Gọi phương thức, nhưng nó làm việc trên bản sao của v
// 	fmt.Println("Sau khi Scale():", v)  // Vẫn là {3 4}, không thay đổi
// }