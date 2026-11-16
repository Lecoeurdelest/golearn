// Bài 1: Hello, 世界
// Đọc trước: https://go.dev/tour/basics/1  →  basics/3
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi:
//   Hello, 世界
//
// Chạy (từ thư mục gốc repo):  go run ./day-1/hello

package main

import "fmt"

// TODO 1: import package "fmt" để dùng hàm in ra màn hình.

func main() {
	// TODO 2: in ra đúng dòng:  Hello, 世界
	//         (Println tự xuống dòng)
	fmt.Println("Hello, 世界")

	// TODO 3 (thử phá): đổi tên hàm in thành chữ thường (ví dụ fmt.println),
	//         chạy lại và ĐỌC lỗi. Sau đó hoàn tác.
	//         Câu hỏi: vì sao chữ hoa/thường ở đầu tên lại quan trọng? (basics/3)
	// fmt.println("Hello, 世界")
}
