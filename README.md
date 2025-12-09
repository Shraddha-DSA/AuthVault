# 🔐 AuthVault – JWT Authentication in Go

AuthVault is a simple and clean **authentication system** built using pure Go (`net/http`).  
It includes user registration, login, JWT tokens, protected routes, admin access, logout, and refresh tokens — all without any database.

---

## 🚀 Features
- User registration with bcrypt hashing  
- Login with **Access + Refresh** tokens  
- JWT-based protected routes  
- Admin-only route  
- Token refresh  
- Logout with blacklist  
- Clean folder structure  

---

## 📁 Project Structure
authvault/
│── main.go
│── handlers/ → Register, Login, Profile, Admin, Refresh, Logout
│── middleware/ → Auth + Admin middleware
│── models/ → User struct
│── storage/ → In-memory users + blacklist
└── utils/ → JWT + Hash utilities

Main APIs
Register
-POST /register

Login
-POST /login

Profile (Protected)
-GET /profile

Header: Authorization: Bearer <access_token>

Admin (Protected)
-GET /admin

Refresh Token
-POST /refresh

Logout
-GET /logout

Future Enhancements
-Move storage from map → SQL database
-Add email OTP verification
-Add forgot password
-Add rate limiting middleware
-Add logging middleware
-Add unit tests
-Add Docker support
