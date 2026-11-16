# golearn

Repo học Go trong 7 ngày, theo [A Tour of Go](https://go.dev/tour/welcome/1).

Nguyên tắc: đọc Tour để hiểu, rồi gõ lại code trong repo này để nhớ. Không chỉ chạy trên Playground.

## Lộ trình

| Ngày | Chủ đề |
|---|---|
| Day 1 | Nền tảng: package, biến, kiểu, hàm |
| Day 2 | Luồng điều khiển |
| Day 3 | Pointer, struct, array, slice |
| Day 4 | Map, closure, method |
| Day 5 | Interface, error, Reader |
| Day 6 | Generics và concurrency cơ bản |
| Day 7 | Concurrency nâng cao, dự án nhỏ, ôn tập |

Chi tiết từng ngày xem [ROADMAP.md](ROADMAP.md).

## Cấu trúc

```
golearn/
  go.mod
  go.sum
  README.md
  ROADMAP.md
  day-1/
    README.md      hướng dẫn và checklist của ngày
    hello/main.go
    vars/main.go
    swap/main.go
    consts/main.go
  day-2/ ...
```

Mỗi ngày có một thư mục `day-N/`. Mỗi bài tập là một thư mục con có `main.go` riêng, là một chương trình độc lập.

## Yêu cầu

- Go 1.27 trở lên (kiểm tra bằng `go version`)

## Cách chạy

Chạy từ thư mục gốc repo:

```sh
go run ./day-1/hello
go run ./day-1/vars
```

Lưu ý phải có `./` ở đầu. Nếu viết `go run day-1/hello`, Go hiểu đó là đường dẫn import chứ không phải thư mục.

Kiểm tra định dạng và lỗi thường gặp sau mỗi bài:

```sh
gofmt -l .
go vet ./...
```

## Cách học mỗi bài

1. Đọc trang Tour tương ứng.
2. Mở `main.go` của bài, làm lần lượt từng `TODO`.
3. Chạy và so sánh với phần "Kết quả mong đợi" ở đầu file.
4. Cuối ngày commit: `git commit -m "day N: ..."`.
