# Lộ trình học Golang — Dành cho Frontend Dev đã biết Express.js

> **Thời gian:** 8 tuần | **Cường độ:** ~1-1.5 tiếng/ngày, 5 ngày/tuần
> **Phương pháp:** Học đến đâu, code ngay đến đó — không học chay lý thuyết

---

## Nguyên tắc học

- [ ] Mắc kẹt quá 20 phút → hỏi AI / StackOverflow, đừng ngồi tự tra mãi
- [ ] Không copy-paste code — gõ tay từng dòng để nhớ
- [ ] Commit code lên GitHub từ ngày đầu (tạo repo `golang-learning`)
- [ ] Cuối mỗi tuần: viết lại những gì đã học bằng lời của bạn (vào file `notes/week-X.md`)

---

## Chuẩn bị môi trường (Ngày 0)

```bash
# Cài Go
brew install go          # macOS
# hoặc tải tại: https://go.dev/dl/

# Kiểm tra
go version               # go1.22.x

# Cài extension VS Code
# → Go (by Google)
# → REST Client (test API không cần Postman)

# Tạo repo học tập
mkdir golang-learning && cd golang-learning
git init
git remote add origin <your-github-repo>
```

---

## TUẦN 1 — Go Syntax (Nền tảng bắt buộc)

> **Mục tiêu:** Đọc và viết được Go cơ bản, hiểu sự khác biệt với JavaScript

### Ngày 1 — Hello Go

**Học:**

- Cấu trúc một file `.go`
- `package main`, `func main()`
- `fmt.Println`, `fmt.Printf`
- Cách chạy: `go run main.go` và `go build`

**Bài tập:**

```
Tạo file main.go
In ra: "Xin chào, tôi đang học Golang!"
In ra ngày hiện tại (dùng package time)
```

**Checkpoint:** Chạy được `go run main.go` không lỗi ✓

---

### Ngày 2 — Variables & Types

**Học:**

- Khai báo biến: `var`, `:=`
- Các kiểu cơ bản: `string`, `int`, `float64`, `bool`
- Zero values (Go không có `undefined` hay `null` với primitive)
- Type conversion: `int(x)`, `string(x)`

**So sánh với JS:**

```go
// JavaScript
let name = "An"
const age = 25

// Go
name := "An"        // short declaration
var age int = 25    // explicit
var salary float64  // zero value = 0.0
```

**Bài tập:**

```
Viết chương trình nhận vào nhiệt độ Celsius (hardcode)
→ Tính và in ra Fahrenheit và Kelvin
→ Dùng đúng kiểu float64
```

---

### Ngày 3 — Control Flow

**Học:**

- `if / else if / else` (không cần dấu ngoặc `()`)
- `for` — Go chỉ có `for`, không có `while`
- `switch` (không cần `break`)

**So sánh với JS:**

```go
// JS: while (i < 10) { i++ }
// Go:
for i < 10 {
    i++
}

// JS: for (let i = 0; i < 10; i++)
// Go:
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**Bài tập:**

```
1. FizzBuzz (1-100)
2. Tính tổng các số chẵn từ 1 đến 100
3. In bảng cửu chương từ 2 đến 9
```

---

### Ngày 4 — Functions

**Học:**

- Khai báo function
- Multiple return values ← khác JS hoàn toàn
- Named return values
- Variadic functions (`...args`)

```go
// Multiple return — rất phổ biến trong Go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("không thể chia cho 0")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Println("Lỗi:", err)
}
```

**Bài tập:**

```
Viết các function:
1. Tính diện tích hình chữ nhật (trả về diện tích, chu vi)
2. Tìm min và max trong danh sách số (trả về cả 2 giá trị)
3. Kiểm tra số nguyên tố (trả về bool, error)
```

---

### Ngày 5 — Arrays, Slices, Maps

**Học:**

- Array: kích thước cố định (ít dùng)
- Slice: mảng động ← dùng nhiều nhất
- Map: giống Object/Map trong JS
- `append`, `len`, `range`

```go
// Slice
fruits := []string{"táo", "cam", "xoài"}
fruits = append(fruits, "nho")

for index, value := range fruits {
    fmt.Printf("%d: %s\n", index, value)
}

// Map
scores := map[string]int{
    "An": 90,
    "Bình": 85,
}
scores["Châu"] = 95
```

**Bài tập:**

```
1. Tính điểm trung bình của danh sách điểm (dùng slice)
2. Đếm số lần xuất hiện của từng chữ cái trong một câu (dùng map)
3. Xóa phần tử trùng lặp trong slice
```

---

### Ngày 6-7 — Structs & Methods

**Học:**

- Struct (thay thế cho Object/Class trong JS)
- Methods (function gắn với struct)
- Pointer receiver vs Value receiver
- Constructor pattern

```go
type Product struct {
    ID    int
    Name  string
    Price float64
}

// Method
func (p Product) PriceWithTax(taxRate float64) float64 {
    return p.Price * (1 + taxRate)
}

// Pointer receiver — có thể thay đổi giá trị gốc
func (p *Product) ApplyDiscount(percent float64) {
    p.Price = p.Price * (1 - percent/100)
}

// Constructor
func NewProduct(id int, name string, price float64) Product {
    return Product{ID: id, Name: name, Price: price}
}
```

**Bài tập:**

```
Tạo struct Student:
- Thuộc tính: ID, Name, Scores ([]float64)
- Method: Average() float64
- Method: Grade() string (A/B/C/D/F)
- Method: String() string (in thông tin đẹp)
```

---

### Mini Project Tuần 1

```
CLI Bill Splitter
- Input: tổng tiền, số người, tip %
- Output: mỗi người trả bao nhiêu
- Bonus: nhập từng món ăn và giá, tự tính tổng
```

```bash
$ go run main.go
Tổng tiền: 850000
Số người: 4
Tip (%): 10
---
Tổng + tip: 935000
Mỗi người trả: 233750 VNĐ
```

---

## TUẦN 2 — Error Handling & Packages

> **Mục tiêu:** Viết Go "đúng cách" — xử lý lỗi như người Go thực sự

### Ngày 8-9 — Error Handling

**Học:**

- `error` interface
- `fmt.Errorf`, `errors.New`
- Custom error types
- Không dùng `try/catch` — Go dùng `if err != nil`

```go
// Custom error
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "phải >= 0"}
    }
    if age > 150 {
        return &ValidationError{Field: "age", Message: "không hợp lệ"}
    }
    return nil
}
```

**Bài tập:**

```
Refactor lại tất cả bài tập tuần 1:
- Mọi function có thể lỗi phải trả về error
- Tạo ít nhất 1 custom error type
```

---

### Ngày 10 — Interfaces

**Học:**

- Interface trong Go (implicit implementation)
- Tại sao interface quan trọng
- Empty interface `interface{}` và `any`

```go
// Bất kỳ struct nào có method Area() đều thỏa mãn Shape
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct { Width, Height float64 }
type Circle struct { Radius float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func printShapeInfo(s Shape) {
    fmt.Printf("Diện tích: %.2f, Chu vi: %.2f\n", s.Area(), s.Perimeter())
}
```

---

### Ngày 11-12 — Packages & Modules

**Học:**

- `go mod init`, `go.mod`, `go.sum`
- Tạo package riêng
- Import package nội bộ
- Cài thư viện ngoài: `go get`

```
golang-learning/
├── go.mod
├── main.go
├── models/
│   └── product.go      ← package models
├── utils/
│   └── calculator.go   ← package utils
└── handlers/
    └── product.go      ← package handlers
```

---

### Ngày 13-14 — Pointers

**Học:**

- Pointer là gì, `*` và `&`
- Khi nào dùng pointer
- Không cần hiểu quá sâu ở giai đoạn này

```go
x := 10
ptr := &x       // ptr giữ địa chỉ của x
*ptr = 20       // thay đổi giá trị tại địa chỉ đó
fmt.Println(x)  // 20
```

---

## TUẦN 3 — Fiber & REST API

> **Mục tiêu:** Viết được REST API với Fiber — thứ bạn đã biết từ Express

### Setup Project

```bash
mkdir shop-api && cd shop-api
go mod init github.com/yourname/shop-api
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/fiber/v2/middleware/logger
go get github.com/gofiber/fiber/v2/middleware/cors
```

```
shop-api/
├── go.mod
├── main.go
├── models/
│   └── product.go
├── handlers/
│   └── product.go
├── routes/
│   └── routes.go
└── middleware/
    └── auth.go
```

---

### Ngày 15-16 — Fiber Basics

**So sánh Express vs Fiber:**

```javascript
// Express
const app = express();
app.use(express.json());
app.use(cors());

app.get("/products", (req, res) => {
  res.json({ success: true, data: products });
});

app.post("/products", (req, res) => {
  const { name, price } = req.body;
  res.status(201).json({ success: true, data: newProduct });
});

app.listen(3000, () => console.log("Server running"));
```

```go
// Fiber — gần như y chang
app := fiber.New()
app.Use(logger.New())
app.Use(cors.New())

app.Get("/products", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"success": true, "data": products})
})

app.Post("/products", func(c *fiber.Ctx) error {
    var body struct {
        Name  string  `json:"name"`
        Price float64 `json:"price"`
    }
    if err := c.BodyParser(&body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(201).JSON(fiber.Map{"success": true, "data": body})
})

app.Listen(":3000")
```

**Bài tập:**

```
CRUD API hoàn chỉnh cho Product (dùng slice in-memory trước):
GET    /api/products        → danh sách
GET    /api/products/:id    → chi tiết
POST   /api/products        → tạo mới
PUT    /api/products/:id    → cập nhật
DELETE /api/products/:id    → xóa
```

---

### Ngày 17 — Middleware

```go
// Custom middleware — giống Express
func LoggerMiddleware(c *fiber.Ctx) error {
    start := time.Now()
    err := c.Next()                          // gọi handler tiếp theo
    duration := time.Since(start)
    fmt.Printf("%s %s - %dms\n", c.Method(), c.Path(), duration.Milliseconds())
    return err
}

app.Use(LoggerMiddleware)
```

**Bài tập:**

```
Viết middleware:
1. Log mọi request (method, path, status, duration)
2. Rate limiter đơn giản (tối đa 10 req/phút per IP)
```

---

### Ngày 18-19 — Validation & Error Handling

```bash
go get github.com/go-playground/validator/v10
```

```go
type CreateProductRequest struct {
    Name     string  `json:"name" validate:"required,min=2,max=100"`
    Price    float64 `json:"price" validate:"required,gt=0"`
    Stock    int     `json:"stock" validate:"gte=0"`
    Category string  `json:"category" validate:"required,oneof=electronics clothing food"`
}

var validate = validator.New()

func CreateProduct(c *fiber.Ctx) error {
    var req CreateProductRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
    }
    if err := validate.Struct(req); err != nil {
        return c.Status(422).JSON(fiber.Map{"error": err.Error()})
    }
    // ... xử lý logic
}
```

---

## TUẦN 4 — Database với PostgreSQL

> **Mục tiêu:** Kết nối database thật, viết CRUD thực sự

### Setup

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

```bash
# .env
DATABASE_URL=postgresql://user:password@localhost:5432/shopdb
PORT=3000
```

---

### Ngày 20-21 — Kết nối Database & GORM Basics

```go
// database/db.go
package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

var DB *gorm.DB

func Connect() {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Không kết nối được database: " + err.Error())
    }
    DB = db
}
```

```go
// models/product.go
type Product struct {
    gorm.Model                          // ID, CreatedAt, UpdatedAt, DeletedAt
    Name        string  `gorm:"not null;size:100"`
    Price       float64 `gorm:"not null"`
    Stock       int     `gorm:"default:0"`
    Category    string
}
```

```go
// Auto migrate
database.DB.AutoMigrate(&models.Product{})
```

---

### Ngày 22-23 — CRUD với GORM

```go
// Tạo
product := models.Product{Name: "iPhone", Price: 25000000}
database.DB.Create(&product)

// Đọc
var products []models.Product
database.DB.Find(&products)
database.DB.Where("price > ?", 1000000).Find(&products)

// Cập nhật
database.DB.Model(&product).Update("price", 26000000)

// Xóa (soft delete nếu có DeletedAt)
database.DB.Delete(&product)
```

**Bài tập:**

```
Refactor CRUD API từ tuần 3:
- Thay slice in-memory bằng PostgreSQL thật
- Thêm tìm kiếm: GET /products?category=electronics&min_price=100
- Thêm phân trang: GET /products?page=1&limit=10
```

---

### Ngày 24-25 — Relationships

```go
type Category struct {
    gorm.Model
    Name     string    `gorm:"unique;not null"`
    Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
    gorm.Model
    Name       string
    Price      float64
    CategoryID uint
    Category   Category
}

// Query với preload
var products []Product
database.DB.Preload("Category").Find(&products)
```

---

## TUẦN 5 — Authentication

> **Mục tiêu:** JWT Auth hoàn chỉnh — giống những gì bạn đã gọi từ frontend

### Ngày 26-27 — User & Password Hashing

```bash
go get golang.org/x/crypto
go get github.com/golang-jwt/jwt/v5
```

```go
type User struct {
    gorm.Model
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Name     string
}

// Hash password trước khi lưu
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

---

### Ngày 28-29 — JWT

```go
// Tạo token
func GenerateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Middleware xác thực
func AuthRequired(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(401).JSON(fiber.Map{"error": "missing token"})
    }
    tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
    // ... verify token
    return c.Next()
}
```

```go
// Routes
auth := app.Group("/api/auth")
auth.Post("/register", handlers.Register)
auth.Post("/login", handlers.Login)

api := app.Group("/api", middleware.AuthRequired)
api.Get("/products", handlers.GetProducts)
api.Post("/products", handlers.CreateProduct)
```

---

### Ngày 30 — Hoàn thiện Auth

**Bài tập:**

```
Hoàn thiện hệ thống auth:
- POST /api/auth/register
- POST /api/auth/login → trả về JWT
- GET  /api/auth/me → thông tin user hiện tại
- Các route /products chỉ cho người đã đăng nhập
```

---

## TUẦN 6 — Concurrency (Điểm mạnh của Go)

> **Mục tiêu:** Hiểu goroutines và channels — thứ làm Go khác biệt

### Ngày 31-32 — Goroutines

```go
// Goroutine = lightweight thread, dùng từ khóa "go"
go func() {
    fmt.Println("Tôi chạy trong goroutine khác")
}()

// WaitGroup để đợi goroutines hoàn thành
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // ... làm việc
    }(i)
}
wg.Wait()
```

**Bài tập:**

```
Gọi 5 API (JSONPlaceholder) song song, tổng hợp kết quả
So sánh thời gian: tuần tự vs song song
```

---

### Ngày 33-34 — Channels

```go
// Channel = pipe để goroutines giao tiếp
ch := make(chan string)

go func() {
    ch <- "kết quả từ goroutine"  // gửi
}()

result := <-ch  // nhận
fmt.Println(result)

// Buffered channel
ch := make(chan int, 10)  // buffer 10 items
```

**Bài tập:**

```
Worker pool:
- Có 100 tasks cần xử lý
- Chạy tối đa 5 workers song song
- Dùng channels để phân phối và thu kết quả
```

---

### Ngày 35 — Context & Timeout

```go
// Hủy request nếu quá 5 giây
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Dùng trong Fiber
func SlowHandler(c *fiber.Ctx) error {
    ctx := c.Context()

    resultCh := make(chan string, 1)
    go func() {
        // ... xử lý chậm
        resultCh <- "done"
    }()

    select {
    case result := <-resultCh:
        return c.JSON(fiber.Map{"data": result})
    case <-ctx.Done():
        return c.Status(408).JSON(fiber.Map{"error": "timeout"})
    }
}
```

---

## TUẦN 7 — Testing & Clean Code

### Ngày 36-37 — Unit Testing

```go
// handlers/product_test.go
func TestCreateProduct(t *testing.T) {
    app := setupTestApp()

    body := `{"name": "Test Product", "price": 100}`
    req := httptest.NewRequest("POST", "/api/products", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    resp, _ := app.Test(req)

    assert.Equal(t, 201, resp.StatusCode)
}
```

```bash
go test ./...              # chạy tất cả test
go test ./... -v           # verbose
go test ./... -cover       # coverage
```

---

### Ngày 38-39 — Project Structure & Clean Code

**Cấu trúc chuẩn cho production:**

```
shop-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/         ← database layer
│   ├── service/            ← business logic
│   └── routes/
├── pkg/
│   ├── database/
│   ├── jwt/
│   └── validator/
├── config/
│   └── config.go
├── .env
├── .env.example
├── go.mod
└── Dockerfile
```

---

### Ngày 40 — Deploy

```dockerfile
# Dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]
```

```bash
# Build và chạy
docker build -t shop-api .
docker run -p 3000:3000 --env-file .env shop-api
```

---

## TUẦN 8 — Project Tổng Hợp

> **Mục tiêu:** Tự làm một project hoàn chỉnh từ 0

### Chọn một trong ba:

#### Option A — E-commerce API

```
Entities: User, Product, Category, Order, OrderItem
Features:
- Auth (register/login/profile)
- Product CRUD với hình ảnh (upload file)
- Giỏ hàng
- Đặt hàng, cập nhật trạng thái
- Phân trang, tìm kiếm, lọc
```

#### Option B — Blog API

```
Entities: User, Post, Comment, Tag
Features:
- Auth với role (admin/author/reader)
- CRUD bài viết, markdown support
- Comment system
- Tag và tìm kiếm
- View count
```

#### Option C — Task Management API

```
Entities: User, Workspace, Project, Task, Comment
Features:
- Multi-workspace
- Assign task, deadline
- Status tracking
- File attachments
- Activity log
```

---

## Checklist Hoàn Thành

### Tuần 1-2 (Go Basics)

- [ ] Viết được function với multiple return values
- [ ] Xử lý error đúng cách (không panic)
- [ ] Dùng được struct, method, interface
- [ ] Hiểu sự khác biệt giữa pointer và value receiver

### Tuần 3 (Fiber)

- [ ] Viết CRUD API hoàn chỉnh
- [ ] Dùng được middleware
- [ ] Validate request body
- [ ] Trả về đúng HTTP status code

### Tuần 4 (Database)

- [ ] Kết nối PostgreSQL thành công
- [ ] CRUD với GORM
- [ ] Query có filter và phân trang
- [ ] Hiểu relationship (has many, belongs to)

### Tuần 5 (Auth)

- [ ] Register/Login với bcrypt
- [ ] Tạo và verify JWT
- [ ] Protected routes với middleware

### Tuần 6 (Concurrency)

- [ ] Hiểu goroutine khác thread ở điểm nào
- [ ] Dùng WaitGroup
- [ ] Dùng channel cơ bản
- [ ] Xử lý timeout với context

### Tuần 7-8

- [ ] Viết unit test cho ít nhất 3 handlers
- [ ] Cấu trúc project đúng chuẩn
- [ ] Build Docker image thành công
- [ ] Project tổng hợp chạy được end-to-end

---

## Tài nguyên

| Tài nguyên    | Link                      | Dùng khi nào  |
| ------------- | ------------------------- | ------------- |
| Go Tour       | `tour.golang.org`         | Tuần 1-2      |
| Go by Example | `gobyexample.com`         | Tra cứu nhanh |
| Fiber Docs    | `docs.gofiber.io`         | Tuần 3+       |
| GORM Docs     | `gorm.io/docs`            | Tuần 4+       |
| Effective Go  | `go.dev/doc/effective_go` | Sau tuần 4    |

---

## Ghi chú cá nhân

> Dùng phần này để ghi lại những điểm khó hoặc "aha moment" khi học

- **Tuần 1:**
- **Tuần 2:**
- **Tuần 3:**
- **Tuần 4:**
- **Tuần 5:**
- **Tuần 6:**
- **Tuần 7:**
- **Tuần 8:**

---

_Tạo ngày: 2026-05-18 | Cập nhật theo tiến độ học_
