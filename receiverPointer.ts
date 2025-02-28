// package receiverPointer

// import (
// 	"fmt"
// )

// // Định nghĩa struct Vertex
// type Vertex struct {
// 	X, Y float64
// }

// // Phương thức Scale() sử dụng con trỏ receiver (*Vertex)
// func (v *Vertex) Scale(factor float64) {
// 	v.X *= factor
// 	v.Y *= factor
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println("Trước khi Scale():", v) // In ra: {3 4}

// 	v.Scale(2) // Gọi phương thức, giá trị gốc sẽ bị thay đổi
// 	fmt.Println("Sau khi Scale():", v)  // Kết quả: {6 8}
// }
