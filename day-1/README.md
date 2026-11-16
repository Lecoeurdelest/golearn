# Day 1 — Nền tảng: package, biến, kiểu, hàm

**Thời gian:** ~2–3 giờ · **Tour:** [welcome](https://go.dev/tour/welcome/1) + [basics/1 → 17](https://go.dev/tour/basics/1)

Cách làm mỗi bài: **đọc trang Tour → mở `main.go` → làm từng `TODO` → chạy → so với "Kết quả mong đợi"**.
Đừng copy từ Playground, hãy gõ tay. Kẹt hơn 20 phút thì mở phần *Gợi ý* (không phải lời giải), vẫn kẹt thì hỏi mình.

---

## Bước 0 — Khởi tạo module (một lần duy nhất)

Chạy ở **thư mục gốc repo** (`golearn/`), không phải trong `day-1/`:

```sh
go version                       # kiểm tra Go đã cài
go mod init golearn              # tạo go.mod
go get golang.org/x/tour@latest  # cần cho các bài Tour ở ngày sau
```

Sau đó kiểm tra bộ khung chạy được (chưa in gì là đúng):

```sh
go run ./day-1/hello
```

<details>
<summary>Gợi ý: vì sao phải có go.mod?</summary>

Từ Go 1.16, `go run` cần biết bạn đang ở trong module nào. `go.mod` ghi tên module (`golearn`) và các thư viện bạn dùng.
Một `go.mod` ở gốc repo là đủ cho cả 7 ngày; mỗi thư mục con (`day-1/hello`, …) là một *package main* riêng.

</details>

### Ba lệnh dùng sau mỗi bài

```sh
go run ./day-1/<tên-bài>   # chạy
gofmt -l day-1             # liệt kê file chưa đúng format (không in gì là tốt). Sửa: gofmt -w day-1
go vet ./day-1/...         # bắt lỗi hay gặp
```

---

## Bài 1 — `hello` · in "Hello, 世界"

**Đọc trước:** [basics/1](https://go.dev/tour/basics/1) (Packages) · [basics/2](https://go.dev/tour/basics/2) (Imports) · [basics/3](https://go.dev/tour/basics/3) (Exported names)
**File:** [hello/main.go](hello/main.go)

**Kết quả mong đợi:**

```
Hello, 世界
```

**Làm:** 3 TODO trong file. TODO 3 là thử phá code cho có lỗi rồi đọc thông báo lỗi. Đây là cách học nhanh nhất.

<details>
<summary>Gợi ý</summary>

- Chương trình chạy được cần `package main` và hàm `func main()`.
- `import "fmt"` hoặc dạng nhóm `import ( "fmt" )`.
- Hàm in nằm trong package `fmt`, tên bắt đầu bằng chữ hoa.
- Lỗi TODO 3 sẽ nói về "unexported name" hoặc "undefined". Tự trả lời: chữ hoa/thường ở đầu tên quyết định điều gì?

</details>

---

## Bài 2 — `vars` · biến, zero value, kiểu cơ bản

**Đọc trước:** [basics/8](https://go.dev/tour/basics/8) → [basics/12](https://go.dev/tour/basics/12) (Variables → Zero values) · [basics/11](https://go.dev/tour/basics/11) (Basic types)
**File:** [vars/main.go](vars/main.go)

**Kết quả mong đợi** (phần `TODO 2`):

```
int      0
float64  0
bool     false
string   ""
*int     <nil>
```

**Làm:** bạn sẽ dùng 3 cách khai báo: `var` ở cấp package, `var` trong hàm, và `:=`. Mục tiêu là *tự thấy* zero value, không phải đọc thuộc.

**Bảng verb của `fmt.Printf`** (dùng suốt khóa):

| Verb | In ra |
|---|---|
| `%v` | giá trị dạng mặc định |
| `%T` | kiểu của giá trị |
| `%q` | chuỗi có dấu ngoặc kép |
| `\n` | xuống dòng (`Printf` không tự thêm, `Println` thì có) |

<details>
<summary>Gợi ý</summary>

- Khai báo không gán: `var i int` → `i` tự có zero value.
- `fmt.Printf("%-8T %v\n", i, i)` căn lề cho cột kiểu. `%q` cho string để thấy chuỗi rỗng.
- `:=` chỉ dùng được **trong hàm**. Thử đặt nó ngoài hàm để đọc lỗi.
- Biến local khai báo mà không dùng thì lỗi biên dịch, còn biến cấp package thì không. Thử cả hai.

</details>

---

## Bài 3 — `swap` · hàm, nhiều giá trị trả về, named return

**Đọc trước:** [basics/4](https://go.dev/tour/basics/4) → [basics/7](https://go.dev/tour/basics/7)
**File:** [swap/main.go](swap/main.go)

**Kết quả mong đợi:**

```
world hello
7 10
```

**Làm:** viết `swap` (trả 2 giá trị) và `split` (named return + *naked return*), rồi gọi cả hai từ `main`.

<details>
<summary>Gợi ý</summary>

- Chữ ký: `func swap(x, y string) (string, string)`: kiểu trả về nằm **sau** danh sách tham số, nhiều kết quả thì đặt trong ngoặc.
- Nhận 2 giá trị: `a, b := swap("hello", "world")`.
- Named return: `func split(sum int) (x, y int)`. `x`, `y` đã được khai báo sẵn, gán trực tiếp rồi viết `return` trần.
- Công thức: `x = sum * 4 / 9`, `y = sum - x`. Với `17` ra `7 10`. Chú ý phép chia số nguyên.
- Tự hỏi: khi nào naked return làm code khó đọc?

</details>

---

## Bài 4 — `consts` · ép kiểu, type inference, constants

**Đọc trước:** [basics/13](https://go.dev/tour/basics/13) → [basics/16](https://go.dev/tour/basics/16)
**File:** [consts/main.go](consts/main.go)

**Kết quả mong đợi:**

```
3 4 5
21
0.2
1.2676506002282295e+29
```

**Làm:** cạnh huyền tam giác (cần ép `int` ↔ `float64` ↔ `uint`, Go **không** tự ép) và hằng số cực lớn `1 << 100`.

<details>
<summary>Gợi ý</summary>

- Ép kiểu viết như gọi hàm: `float64(x)`, `uint(f)`.
- `math.Sqrt` nhận và trả `float64`, nên phải ép `x*x + y*y` sang `float64` trước.
- Hằng số không có kiểu cố định cho tới khi dùng, nên `Big` (1 << 100) đưa vào hàm nhận `float64` được nhưng đưa vào `int` thì lỗi tràn. TODO 4 cho bạn tự thấy lỗi đó.
- `const` không dùng được với `:=`. Phải viết `const Big = 1 << 100`.

</details>

---

## Cuối ngày

Đánh dấu khi làm xong:

- [ ] `hello`, `vars`, `swap`, `consts` đều chạy đúng kết quả mong đợi
- [ ] `gofmt -l day-1` không in gì, `go vet ./day-1/...` không báo lỗi
- [ ] Đọc hết Tour `basics` đến trang *Congratulations!*

**Tự kiểm tra** (trả lời được mà không nhìn lại):

1. Vì sao `fmt.println` lỗi còn `fmt.Println` thì không?
2. `:=` dùng được ở đâu, không dùng được ở đâu?
3. Zero value của `int`, `string`, `bool`, `*int` là gì?
4. Với `x := 3` (kiểu `int`), vì sao `var f float64 = x` lỗi biên dịch còn `var f float64 = 3` thì được?
5. `var x, y int` và `x, y := 1, 2` khác nhau chỗ nào?

Rồi commit:

```sh
git add -A && git commit -m "day 1: basics"
```

Xong Day 1, quay lại [ROADMAP.md](../ROADMAP.md) mở Day 2.
