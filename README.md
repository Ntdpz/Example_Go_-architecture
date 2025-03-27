# Project Structure (Go)

## Overview
โปรเจกต์นี้ใช้ภาษา Go และมีโครงสร้างที่แบ่งแยกส่วนต่างๆ อย่างชัดเจน เพื่อให้ง่ายต่อการดูแลรักษาและพัฒนา

## Folder Structure

```
/project-root
│── /database
│   ├── db.go
│── /cmd
│   ├── main.go
│── /internal
│   ├── /repository
│   │   ├── user_repository.go
│   ├── /handlers
│   │   ├── user_handlers.go
│   ├── /service
│   │   ├── user_service.go
│── /config
│   ├── .env
│   ├── config.go
│── /routers
│   ├── router.go
│── /test
│   ├── test.go
│── docker-compose.yaml
│── Dockerfile
```

## Folder & File Explanation

### `/database`
- **db.go**: จัดการการเชื่อมต่อกับฐานข้อมูล เช่น การเปิดและปิด connection

### `/cmd`
- **main.go**: ไฟล์หลักของโปรเจกต์ที่ใช้ในการรันแอปพลิเคชัน โดยจะทำการโหลด configuration, สร้าง router และเริ่มต้น server

### `/internal`

#### `/repository`
- **user_repository.go**: ชั้นที่ใช้ในการเข้าถึงฐานข้อมูล (Data Access Layer) สำหรับ `User` เช่น คำสั่ง SQL, การ query ข้อมูล ฯลฯ

#### `/handlers`
- **user_handlers.go**: จัดการ HTTP request/response สำหรับ `User` เช่น API endpoints, รับ request, response JSON ฯลฯ

#### `/service`
- **user_service.go**: ชั้น business logic สำหรับ `User` เช่น การ validate ข้อมูล, ประมวลผลข้อมูล, เรียกใช้ repository

### `/config`
- **.env**: เก็บค่าคอนฟิกต่างๆ เช่น Database URL, API keys, Secrets
- **config.go**: โหลดค่า config จากไฟล์ `.env` และจัดการ environment variables

### `/routers`
- **router.go**: กำหนดเส้นทาง (Routes) และเชื่อมโยงไปยัง handlers ต่างๆ

### `/test`
- **test.go**: รวม unit tests หรือ integration tests สำหรับแอปพลิเคชัน

### Docker & Deployment
- **docker-compose.yaml**: ใช้สำหรับกำหนด container ต่างๆ เช่น Go app, Database, Redis ฯลฯ
- **Dockerfile**: ใช้สร้าง Docker image สำหรับแอปพลิเคชัน

---
## How to Run

### Run with Go
```sh
# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```

### Run with Docker
```sh
# Build and run with Docker
docker-compose up --build
```

---
## License
This project is licensed under the MIT License.

