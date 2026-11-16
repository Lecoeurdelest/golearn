// Bài 3: Hàm, nhiều giá trị trả về, named return
// Đọc trước: https://go.dev/tour/basics/4  →  basics/7
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi:
//   world hello
//   7 10
//
// Chạy (từ thư mục gốc repo):  go run ./day-1/swap

package main

// TODO 0: import "fmt"
import "fmt"

// TODO 1: viết hàm swap nhận 2 string và trả về 2 string theo thứ tự ĐẢO NGƯỢC.
//
//	Chữ ký: func swap(x, y string) (string, string)   (basics/6)
func swap(x, y string) (string, string) {
	fmt.Printf("HEHE")
	return y, x
}

// TODO 2: viết hàm split(sum int) (x, y int) dùng NAMED RETURN (basics/7):
//
//	  x = sum * 4 / 9
//	  y = sum - x
//	và kết thúc bằng `return` trần (naked return).
func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return 0, y
}

func main() {
	// TODO 3: gọi swap("hello", "world"), nhận 2 giá trị về 2 biến, in ra: world hello
	fmt.Println(swap("Hello", "world"))
	// TODO 4: gọi split(17), nhận 2 giá trị, in ra: 7 10
	fmt.Println(split(17))
	// TODO 5 (tự hỏi): viết lại split KHÔNG dùng named return.
	//         Bản nào dễ đọc hơn? Khi nào naked return là ý tồi?

}
