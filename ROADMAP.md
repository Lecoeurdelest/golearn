# Lộ trình học Go trong 7 ngày

Nguồn chính: [A Tour of Go](https://go.dev/tour/welcome/1) (~100 trang, 7 module).
Nguyên tắc: **đọc Tour để hiểu → gõ lại code trong repo này để nhớ**. Không chỉ chạy trên Playground.

- Thời gian: ~2–3 giờ/ngày (ngày 5 và 7 nặng hơn, dự phòng 3–4 giờ).
- Mỗi ngày 1 thư mục `day-N/`, mỗi bài tập 1 thư mục con có `main.go`.
- Mỗi ngày kết thúc bằng 1 commit (`git commit -m "day N: ..."`).

## Chuẩn bị (15 phút, làm ở Day 1)

```sh
go version                      # đã có go1.27.1
cd golearn && go mod init golearn
go get golang.org/x/tour@latest # các exercise Pic, WordCount, Readers, Binary Trees import gói này
mkdir -p day-1/hello && cd day-1/hello
# tạo main.go, rồi:
go run .
gofmt -l . ; go vet ./...       # chạy sau mỗi bài
```

Cấu trúc đích:

```
golearn/
├── go.mod
├── ROADMAP.md
├── day-1/ hello/  vars/  ...
├── day-2/ sqrt/   fizzbuzz/ ...
└── day-7/ linkcheck/
```

Muốn chạy Tour offline: `go install golang.org/x/website/tour@latest && $(go env GOPATH)/bin/tour`.

> **Lưu ý:** chạy local, `pic.Show` (Exercise Pic, Images) không hiện ảnh mà in một dòng `IMAGE:<base64>`.
> Muốn xem ảnh thì dán code lên [Playground](https://go.dev/tour/moretypes/18); còn không thì chỉ cần code chạy không lỗi.

---

## Day 1 — Nền tảng: package, biến, kiểu, hàm

**Tour:** `welcome` (cả module) + `basics/1 → 17` (~20 trang)

| Học | Trang |
|---|---|
| Package, import, exported name (viết hoa) | basics/1–3 |
| Hàm, multiple results, named returns | basics/4–7 |
| `var`, `:=`, zero value, ép kiểu, type inference | basics/8–14 |
| Constants | basics/15–16 |

**Làm trong repo:**
1. `day-1/hello` — in "Hello, 世界" bằng `fmt`.
2. `day-1/vars` — khai báo mọi kiểu cơ bản, in ra zero value của từng kiểu bằng `%v` và `%T`.
3. `day-1/swap` — hàm `swap(a, b string) (string, string)` và 1 hàm dùng named return.
4. `day-1/consts` — ép kiểu (cạnh huyền 3-4-5) và hằng số `1 << 100`.

Hướng dẫn chi tiết từng bước và bộ khung code: [day-1/README.md](day-1/README.md).

**Tự kiểm tra (trả lời được mà không nhìn lại):**
- Vì sao `fmt.println` lỗi còn `fmt.Println` thì không?
- `:=` dùng được ở đâu, không dùng được ở đâu?
- Zero value của `int`, `string`, `bool`, `*int` là gì?

---

## Day 2 — Luồng điều khiển

**Tour:** `flowcontrol/1 → 14` (14 trang)

| Học | Trang |
|---|---|
| `for` (3 dạng: đủ, như `while`, vô hạn) | 1–4 |
| `if` với câu lệnh khởi tạo, `else` | 5–7 |
| **Exercise: Loops and Functions** (Newton's sqrt) | 8 |
| `switch` (không cần `break`, switch không điều kiện) | 9–11 |
| `defer`, thứ tự LIFO | 12–13 |

**Làm trong repo:**
1. `day-2/sqrt` — Exercise của Tour: tự cài `Sqrt(x)` bằng Newton, so với `math.Sqrt`.
2. `day-2/fizzbuzz` — dùng `switch` không điều kiện.
3. `day-2/defer` — in 5 dòng bằng `defer` trong vòng lặp, đoán thứ tự trước khi chạy.

**Tự kiểm tra:**
- Go không có `while` — viết lại vòng `while` bằng `for`.
- `defer` đánh giá **tham số** lúc nào, chạy hàm lúc nào?

---

## Day 3 — Pointer, struct, array, slice

**Tour:** `moretypes/1 → 18` (18 trang)

| Học | Trang |
|---|---|
| Pointer (không có pointer arithmetic) | 1 |
| Struct, field, pointer-to-struct, literal | 2–5 |
| Array vs slice, slice là *view* của array | 6–8 |
| Slice literal, len/cap, nil slice, `make`, `append` | 9–15 |
| `range` | 16–17 |
| **Exercise: Slices** (`Pic`) | 18 |

**Làm trong repo:**
1. `day-3/pointers` — hàm nhận `int` vs `*int`, chứng minh khác biệt.
2. `day-3/slices` — tạo 1 slice, cắt `s[1:3]`, sửa phần tử, quan sát slice gốc; in `len`/`cap` sau mỗi lần `append`.
3. `day-3/pic` — Exercise `Pic`.

**Tự kiểm tra (đây là chỗ người mới hay vấp nhất):**
- Hai slice cùng trỏ vào một array: sửa 1 cái, cái kia có đổi không?
- Khi nào `append` cấp phát array mới?
- Vì sao `for i, v := range s` sửa `v` không đổi `s`?

---

## Day 4 — Map, closure, và method

**Tour:** `moretypes/19 → 27` + `methods/1 → 8` (~17 trang)

| Học | Trang |
|---|---|
| Map, literal, thêm/xóa/kiểm tra tồn tại (`v, ok := m[k]`) | moretypes/19–22 |
| **Exercise: Maps** (`WordCount`) | moretypes/23 |
| Function là giá trị, closure | moretypes/24–25 |
| **Exercise: Fibonacci closure** | moretypes/26 |
| Method trên struct, pointer receiver vs value receiver | methods/1–8 |

**Làm trong repo:**
1. `day-4/wordcount` — Exercise `WordCount`.
2. `day-4/fib` — Exercise Fibonacci closure.
3. `day-4/shapes` — struct `Rect`, `Circle` với method `Area()`; thêm 1 method đổi kích thước dùng pointer receiver.

**Tự kiểm tra:**
- Đọc key không tồn tại từ map trả về gì? Ghi vào map `nil` thì sao?
- Khi nào bắt buộc dùng pointer receiver?

---

## Day 5 — Interface, error, Reader (ngày nặng nhất)

**Tour:** `methods/9 → 26` (18 trang, nhiều exercise)

| Học | Trang |
|---|---|
| Interface, implement ngầm, interface value, nil interface | 9–13 |
| `interface{}` / `any`, type assertion, type switch | 14–16 |
| `Stringer` + **Exercise** | 17–18 |
| `error` + **Exercise: Errors** | 19–20 |
| `io.Reader` + **Exercise: Readers**, **rot13Reader** | 21–23 |
| `image` + **Exercise: Images** (tùy chọn) | 24–25 |

**Làm trong repo:**
1. `day-5/stringer` — `IPAddr` với `String()`.
2. `day-5/errors` — `Sqrt` trả `ErrNegativeSqrt`; thêm 1 hàm dùng `errors.Is` / `errors.As`.
3. `day-5/readers` — `MyReader` (stream ký tự `'A'`) và `rot13Reader`.
4. Bỏ qua Images nếu quá giờ — không ảnh hưởng ngày sau.

**Tự kiểm tra:**
- Một interface chứa con trỏ nil có `== nil` không? Vì sao đây là bẫy?
- Quy tắc idiom Go: "accept interfaces, return structs" nghĩa là gì?

---

## Day 6 — Generics và Concurrency cơ bản

**Tour:** `generics/1 → 3` + `concurrency/1 → 8` (11 trang)

| Học | Trang |
|---|---|
| Type parameter, generic type | generics/1–2 |
| Goroutine | concurrency/1 |
| Channel, buffered channel, `range` + `close` | 2–4 |
| `select`, default case | 5–6 |
| **Exercise: Equivalent Binary Trees** (chỉ cần goroutine + channel) | 7–8 |

**Làm trong repo:**
1. `day-6/generics` — `Map[T, U any]`, `Filter[T any]`, `Sum` với constraint `~int | ~float64`.
2. `day-6/sum` — chia slice làm 2, 2 goroutine tính tổng, gửi qua channel (đúng ví dụ Tour, tự gõ lại).
3. `day-6/fanin` — 2 producer + `select` để nhận từ cả hai kèm timeout (`time.After`).
4. `day-6/trees` — Exercise Binary Trees (`Walk`, `Same`).

**Tự kiểm tra:**
- Ai nên là bên `close` channel — sender hay receiver?
- Gửi vào channel không có ai nhận thì chuyện gì xảy ra?
- Chạy với `go run -race .` để thấy data race là gì (cố ý tạo một cái).

---

## Day 7 — Concurrency nâng cao + dự án nhỏ + ôn tập

**Tour:** `concurrency/9 → 11` (3 trang)

| Học | Trang |
|---|---|
| `sync.Mutex` | 9 |
| **Exercise: Web Crawler** | 10 |
| Where to go from here | 11 |

**Làm trong repo (buổi sáng — hoàn tất Tour):**
1. `day-7/mutex` — bộ đếm an toàn với `sync.Mutex` (ví dụ `SafeCounter` của Tour); chạy thử có và không có mutex bằng `go run -race .`.
2. `day-7/crawler` — Exercise Web Crawler dùng `sync.Mutex` hoặc channel. Đây là bài khó nhất Tour, dành riêng thời gian.

**Dự án tốt nghiệp (buổi chiều, ~2 giờ): `day-7/linkcheck`**
CLI nhận danh sách URL (flag hoặc file), kiểm tra song song bằng `net/http`, in bảng `URL | status | thời gian`.
Dùng lại: struct, slice, map, error, goroutine, channel, mutex/WaitGroup, `defer resp.Body.Close()`.
Thêm 1 file `linkcheck_test.go` với `httptest.NewServer` — đây là lần đầu bạn chạm `go test`.

**Ôn tập cuối tuần:** đọc lại các mục "Tự kiểm tra" của 7 ngày, đánh dấu mục nào chưa trả lời được và quay lại đúng trang Tour đó.

---

## Sau tuần này (không bắt buộc)

| Mục tiêu | Tài liệu |
|---|---|
| Cú pháp và idiom | [Effective Go](https://go.dev/doc/effective_go), [Go by Example](https://gobyexample.com) |
| Test | `go test`, table-driven tests, `testing` package |
| Thư viện chuẩn hay dùng | `net/http`, `encoding/json`, `context`, `flag`, `os`, `bufio` |
| Dự án tiếp theo | REST API nhỏ (todo) dùng `net/http` + `encoding/json` |

## Mẹo giữ nhịp

- Gõ lại code, đừng copy từ Playground.
- Mỗi bài chạy `gofmt` và `go vet` — hình thành thói quen sớm.
- Kẹt > 20 phút ở 1 exercise: xem lời giải chính thức tại [golang/website/tour/solutions](https://github.com/golang/website/tree/master/_content/tour/solutions), nhưng tự viết lại sau khi hiểu.
- Trễ 1 ngày không sao: Day 4 và Day 5 là hai ngày hay dồn nhất, dùng ngày 7 buổi sáng làm bù.
