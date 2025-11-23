# Live Coding Project

یک پروژه REST API ساده برای مدیریت کاربران با استفاده از معماری Clean Architecture در زبان Go.

## 📋 فهرست مطالب

- [ویژگی‌ها](#ویژگی‌ها)
- [تکنولوژی‌های استفاده شده](#تکنولوژی‌های-استفاده-شده)
- [ساختار پروژه](#ساختار-پروژه)
- [نصب و راه‌اندازی](#نصب-و-راه‌اندازی)
- [پیکربندی](#پیکربندی)
- [اجرای پروژه](#اجرای-پروژه)
- [API Endpoints](#api-endpoints)
- [مایگریشن دیتابیس](#مایگریشن-دیتابیس)

## ✨ ویژگی‌ها

- معماری Clean Architecture
- RESTful API با استفاده از Echo Framework
- پایگاه داده PostgreSQL با GORM
- مدیریت کاربران و آدرس‌ها
- CLI Commands با استفاده از Cobra
- پشتیبانی از Migration با dbmate

## 🛠 تکنولوژی‌های استفاده شده

- **Go 1.25.3** - زبان برنامه‌نویسی
- **Echo v4** - فریمورک HTTP
- **GORM** - ORM برای پایگاه داده
- **PostgreSQL** - پایگاه داده رابطه‌ای
- **Cobra** - کتابخانه CLI
- **dbmate** - مدیریت مایگریشن
- **godotenv** - مدیریت متغیرهای محیطی

## 📁 ساختار پروژه

```
live-coding/
├── cmd/                    # نقطه ورود برنامه
│   ├── main.go
│   └── command/           # دستورات CLI
│       ├── http.go        # دستور اجرای سرور HTTP
│       ├── migration.go   # دستورات مایگریشن
│       └── root.go        # تنظیمات root command
├── config/                 # پیکربندی پروژه
│   └── config.go
├── internal/              # کدهای داخلی پروژه
│   ├── adapter/           # لایه اتصال به دنیای بیرون
│   │   ├── migrations/    # فایل‌های مایگریشن SQL
│   │   └── repository/    # پیاده‌سازی repository
│   ├── domain/            # لایه دامنه
│   │   └── entity/        # موجودیت‌های دامنه
│   ├── entrypoint/        # لایه ورودی
│   │   ├── handler/       # HTTP handlers
│   │   └── http.go        # تنظیمات HTTP
│   └── usecase/           # لایه منطق کسب و کار
│       └── user.go
├── scripts/               # اسکریپت‌های کمکی
├── go.mod                 # وابستگی‌های Go
└── README.md
```

## 🚀 نصب و راه‌اندازی

### پیش‌نیازها

- Go 1.25.3 یا بالاتر
- PostgreSQL

### نصب وابستگی‌ها

```bash
go mod download
```

## ⚙️ پیکربندی

یک فایل `.env` در ریشه پروژه ایجاد کنید:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_WRITE_TIMEOUT=10s
SERVER_READ_TIMEOUT=10s
SERVER_DEBUG=false

# Database Configuration
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=your_password
DATABASE_NAME=livecoding
DATABASE_SSLMODE=disable
DATABASE_MAX_OPEN_CONNECTION=100

# Service Configuration
SERVICE_NAME=livecoding
```

## 🏃 اجرای پروژه

### اجرای مایگریشن

ابتدا باید جداول دیتابیس را ایجاد کنید:

```bash
go run cmd/main.go migration up --env-file .env
```

### اجرای سرور HTTP

```bash
go run cmd/main.go http --env-file .env
```

سرور روی پورت `8080` اجرا می‌شود.

## 📡 API Endpoints

### دریافت کاربر بر اساس ID

```http
GET /users/:id
```

**مثال:**
```bash
curl http://localhost:8080/users/1
```

**پاسخ موفق (200 OK):**
```json
{
  "id": "1",
  "name": "John Doe",
  "email": "john@example.com",
  "phone_number": "+1234567890",
  "addresses": [
    {
      "street": "123 Main St",
      "city": "New York",
      "state": "NY",
      "zip_code": "10001",
      "country": "USA"
    }
  ]
}
```

**خطاها:**
- `400 Bad Request` - ID نامعتبر
- `404 Not Found` - کاربر یافت نشد
- `500 Internal Server Error` - خطای سرور

## 🗄 مایگریشن دیتابیس

پروژه از `dbmate` برای مدیریت مایگریشن استفاده می‌کند.

### دستورات مایگریشن

```bash
# اجرای مایگریشن‌ها
go run cmd/main.go migration up --env-file .env

# بازگردانی آخرین مایگریشن
go run cmd/main.go migration down --env-file .env

# نمایش وضعیت مایگریشن‌ها
go run cmd/main.go migration status --env-file .env
```

## 📝 موجودیت‌ها

### User
- `ID`: شناسه یکتا
- `Name`: نام کاربر
- `Email`: ایمیل
- `PhoneNumber`: شماره تلفن
- `Addresses`: لیست آدرس‌ها

### Address
- `Street`: نام خیابان
- `City`: شهر
- `State`: استان/ایالت
- `ZipCode`: کد پستی
- `Country`: کشور
