# Create a professional README.md file for the Task-Tracker-API project.

readme_content = """# 🚀 Task Tracker API (Golang)

Ushbu loyiha **Go** tili va **Gin** freymvorki yordamida yaratilgan, vazifalarni boshqarish uchun mo'ljallangan yengil va samarali **RESTful API** hisoblanadi. Loyiha orqali foydalanuvchilarni ro'yxatdan o'tkazish, autentifikatsiya qilish va shaxsiy vazifalar (todo-list) ro'yxatini boshqarish mumkin.

---

## 🛠 Texnologiyalar
* **Language:** [Go (Golang)](https://go.dev/)
* **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
* **Data Persistence:** JSON (Mahalliy fayl tizimi) — *Bazaga ulanish shart bo'lmagan yengil saqlash usuli.*

## 📁 Loyiha Strukturasi
Loyiha modulli strukturaga ega bo'lib, quyidagi qismlardan iborat:

* 📂 `cmd/main.go` — Dasturning kirish nuqtasi, marshrutlarni (routes) sozlash.
* 📂 `internal/`
    * `model.go` — Ma'lumotlar modeli (`User`, `Task` struct'lari).
    * `handler.go` — HTTP so'rovlariga ishlov beruvchi mantiqiy qism.
    * `repository.go` — Ma'lumotlarni JSON fayllari bilan o'qish/yozish operatsiyalari.

---

## ⚙️ Imkoniyatlar
1.  **Foydalanuvchi tizimi:**
    * Yangi foydalanuvchi yaratish (`/register`).
    * Login qilish orqali identifikatsiya (`/login`).
2.  **Vazifalar boshqaruvi:**
    * Yangi vazifa qo'shish.
    * Vazifalarni `user_id` orqali filtrlab olish.
3.  **Persistence:**
    * Dastur to'xtatilsa ham, ma'lumotlar `.json` fayllarida xavfsiz saqlanadi.

---

## 🚀 O'rnatish va Ishga tushirish

Loyiha ustida ishlash uchun quyidagi qadamlarni bajaring:

1.  **Repozitoriyani klonlang:**
    ```bash
    git clone [https://github.com/7zet/Task-Tracker-API.git](https://github.com/7zet/Task-Tracker-API.git)
    cd Task-Tracker-API
    ```

2.  **Bog'liqliklarni yuklang:**
    ```bash
    go mod tidy
    ```

3.  **Serverni ishga tushiring:**
    ```bash
    go run cmd/main.go
    ```

---

## 📡 API Yo'llari (Endpoints)

| Metod | Yo'nalish | Tavsif |
| :--- | :--- | :--- |
| `POST` | `/register` | Yangi foydalanuvchi yaratish |
| `POST` | `/login` | Tizimga kirish va autentifikatsiya |
| `POST` | `/tasks` | Yangi vazifa yaratish |
| `GET` | `/tasks?user_id=N` | Ma'lum bir foydalanuvchining hamma vazifalari |

### 💡 So'rov namunasi (JSON)
**`POST /tasks`**
```json
{
  "user_id": 1,
  "title": "Go o'rganish",
  "status": "pending",
  "body": "GORM va PostgreSQL mavzusini tugatish"
}
