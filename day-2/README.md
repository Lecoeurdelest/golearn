# Day 2 — Luồng điều khiển

**Thời gian:** ~2–3 giờ · **Tour:** [flowcontrol/1 → 14](https://go.dev/tour/flowcontrol/1)

Cách làm giống Day 1: **đọc trang Tour, mở `main.go`, làm từng `TODO`, chạy, so với "Kết quả mong đợi"**.
Gõ tay, không copy từ Playground. Kẹt hơn 20 phút thì mở phần *Gợi ý* (không phải lời giải), vẫn kẹt thì hỏi mình.

---

## Bước 0 — Kiểm tra module

Day 2 dùng chung `go.mod` ở **thư mục gốc repo** (`golearn/`) với Day 1, không tạo mới.
Kiểm tra bằng cách chạy từ thư mục gốc:

```sh
go run ./day-2/sqrt
```

Chưa in gì và không có lỗi là đúng. Nếu báo `go.mod file not found`, bạn cần chuyển `go.mod` và `go.sum` từ `day-1/` lên thư mục gốc:

```sh
git mv day-1/go.mod day-1/go.sum .
```

### Ba lệnh dùng sau mỗi bài

```sh
go run ./day-2/<tên-bài>   # chạy
gofmt -l day-2             # liệt kê file chưa đúng format (không in gì là tốt). Sửa: gofmt -w day-2
go vet ./day-2/...         # bắt lỗi hay gặp
```

---

## Bài 1 — `sqrt` · vòng lặp và hàm (Newton)

**Đọc trước:** [flowcontrol/1](https://go.dev/tour/flowcontrol/1) → [flowcontrol/8](https://go.dev/tour/flowcontrol/8)
**File:** [sqrt/main.go](sqrt/main.go)

**Kết quả mong đợi:** `Sqrt(2)` gần bằng `math.Sqrt(2)` (khoảng `1.41421356237`), chênh lệch cỡ `1e-16`.

**Làm:** tự cài căn bậc hai bằng phương pháp Newton, rồi so với `math.Sqrt`. Sau đó nâng cấp từ "lặp đúng 10 lần" sang "lặp đến khi đủ chính xác".

<details>
<summary>Gợi ý</summary>

- Vòng `for` đủ 3 phần: `for i := 0; i < 10; i++ { ... }`.
- Vòng "như while": bỏ phần khởi tạo và phần sau, chỉ giữ điều kiện: `for điều_kiện { ... }`. Go không có từ khóa `while`.
- Cập nhật Newton: `z -= (z*z - x) / (2 * z)`. Với `x = 2`, chỉ cần vài lần lặp là hội tụ.
- Để so "thay đổi nhỏ", lưu giá trị cũ vào biến, trừ cho giá trị mới rồi lấy trị tuyệt đối bằng `math.Abs`.
- Giá trị in ra có thể khác `math.Sqrt` ở chữ số cuối. So sánh số thực bằng `==` là thói quen xấu, hãy so hiệu với một ngưỡng nhỏ.

</details>

---

## Bài 2 — `fizzbuzz` · switch không điều kiện

**Đọc trước:** [flowcontrol/9](https://go.dev/tour/flowcontrol/9) → [flowcontrol/11](https://go.dev/tour/flowcontrol/11) (và [flowcontrol/5](https://go.dev/tour/flowcontrol/5) → [7](https://go.dev/tour/flowcontrol/7) cho `if`)
**File:** [fizzbuzz/main.go](fizzbuzz/main.go)

**Kết quả mong đợi** (1 đến 15):

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

**Làm:** in 1 đến 15 theo luật FizzBuzz, bằng `switch` không có điều kiện. TODO 2 là thử đảo thứ tự `case` để thấy `switch` dừng ở nhánh đúng đầu tiên.

<details>
<summary>Gợi ý</summary>

- `switch` không điều kiện viết gọn thay cho chuỗi `if / else if`: `switch { case i%15 == 0: ... }`.
- Go **không** cần `break` ở cuối mỗi `case`, nó tự dừng sau nhánh đã chạy.
- Phải kiểm tra chia hết cho 15 **trước** 3 và 5, nếu không `15` sẽ rơi vào nhánh `Fizz`.
- `if` có thể kèm câu lệnh khởi tạo: `if r := i % 3; r == 0 { ... }`. Biến `r` chỉ tồn tại trong khối if/else đó.

</details>

---

## Bài 3 — `defer` · thứ tự chạy LIFO

**Đọc trước:** [flowcontrol/12](https://go.dev/tour/flowcontrol/12) → [flowcontrol/13](https://go.dev/tour/flowcontrol/13)
**File:** [defer/main.go](defer/main.go)

**Kết quả mong đợi** (TODO 1 và 2):

```
start
end loop
4
3
2
1
0
```

**Làm:** dùng `defer` trong vòng lặp, **đoán trước** thứ tự output rồi mới chạy. Sau đó thử đoạn code ở TODO 3 để hiểu `defer` đánh giá tham số lúc nào.

<details>
<summary>Gợi ý</summary>

- Lệnh `defer` hoãn việc **gọi hàm** đến khi hàm bao ngoài (ở đây là `main`) kết thúc.
- Nhiều `defer` được xếp vào một ngăn xếp: cái đến sau chạy trước (LIFO).
- **Tham số** của hàm bị hoãn được tính ngay khi gặp `defer`, không chờ đến lúc chạy. Đó là lý do TODO 3 in `deferred x = 1` dù `x` sau đó đổi thành `100`.
- Thực tế `defer` dùng nhiều nhất để dọn dẹp: đóng file, giải phóng khóa (mutex), đóng kết nối.

</details>

---

## Cuối ngày

Đánh dấu khi làm xong:

- [ ] `sqrt`, `fizzbuzz`, `defer` đều chạy đúng kết quả mong đợi
- [ ] `gofmt -l day-2` không in gì, `go vet ./day-2/...` không báo lỗi
- [ ] Đọc hết Tour `flowcontrol` đến trang *Congratulations!*

**Tự kiểm tra** (trả lời được mà không nhìn lại):

1. Go không có `while`. Viết lại một vòng `while` bằng `for`.
2. Vòng `for` có 3 dạng nào? Dạng nào là vòng lặp vô hạn?
3. Biến khai báo trong câu lệnh khởi tạo của `if` sống đến đâu?
4. Vì sao `switch` trong Go không cần `break`? Khi nào bạn muốn chạy tiếp sang case sau (từ khóa `fallthrough`)?
5. `defer` đánh giá **tham số** lúc nào và chạy **hàm** lúc nào?
6. Ba lệnh `defer` liên tiếp chạy theo thứ tự nào?

Rồi commit:

```sh
git add -A && git commit -m "day 2: control flow"
```

Xong Day 2, quay lại [ROADMAP.md](../ROADMAP.md) mở Day 3.
