# 🚀 Task Tracker API (Golang)

Ushbu loyiha Go tili va Gin freymvorki yordamida yaratilgan vazifalarni boshqarish uchun mo'ljallangan RESTful API hisoblanadi. Loyiha o'z ichiga foydalanuvchilarni ro'yxatdan o'tkazish, autentifikatsiya va shaxsiy vazifalar ro'yxatini boshqarish funksiyalarini oladi.

## 🛠 Texnologiyalar
- **Language:** Go (Golang)
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Data Persistence:** JSON (Local File Storage)

## 📁 Loyiha Strukturasi
- `cmd/main.go`: Dasturning kirish nuqtasi va router konfiguratsiyasi.
- `internal/model.go`: User va Task ma'lumotlarining strukturasi (structs).
- `internal/handler.go`: HTTP so'rovlarni qabul qilish va mantiqiy ishlov berish.
- `internal/repository.go`: Ma'lumotlarni JSON fayllari bilan sinxronizatsiya qilish.

## ⚙️ Imkoniyatlar
1. **Foydalanuvchi tizimi:**
   - Registratsiya (`POST /register`)
   - Login (`POST /login`)
2. **Vazifalar boshqaruvi:**
   - Yangi vazifa yaratish (`POST /tasks`)
   - Foydalanuvchiga tegishli vazifalarni `user_id` orqali filtrlab ko'rish (`GET /tasks?user_id=1`)
3. **Ma'lumotlar xavfsizligi:**
   - Dastur restart bo'lganda ham ma'lumotlar `json` fayllarida saqlanib qoladi.

## 🚀 O'rnatish va Ishga tushirish

1. Repozitoriyani klon qiling:
```bash
git clone [https://github.com/7zet/Task-Tracker-API.git](https://github.com/7zet/Task-Tracker-API.git)
cd Task-Tracker-API



Bog'liqliklarni yuklang:Bashgo mod tidy
Serverni ishga tushiring:Bashgo run cmd/main.go
📡 API Yo'llari (Endpoints)MetodYo'nalishTavsifPOST/registerYangi foydalanuvchi yaratishPOST/loginTizimga kirish va tekshirishPOST/tasksYangi vazifa yaratishGET/tasks?user_id=NN-idli foydalanuvchining hamma vazifalari💡 Namuna (Request Example)Create Task:
{
  "user_id": 1,
  "title": "Go organish",
  "status": "pending",
  "body": "GORM va PostgreSQL mavzusini tugatish"
}
📝 Roadmap[ ] Malumotlarni PostgreSQL malumotlar bazasiga kochirish.[ ] JWT (JSON Web Token) orqali xavfsizlikni ta'minlash.[ ] Vazifalarni tahrirlash (Update) va ochirish (Delete) funksiyalari.
Author: 7zet
---

### Push qilish uchun kichik eslatma:

GitHub-ga push qilishdan oldin, loyihang papkasida `.json` fayllari (agar ular avtomatik yaratilsa) va kesh fayllar git-ga chiqib ketmasligi uchun `.gitignore` faylini yaratishingni maslahat beraman. Ichiga shunchaki buni yoz:

```text
# Binary files
main
*.exe

# JSON storage
*.json
