# 📧 API Email Confirmation

A simple Go API for sending confirmation emails with a unique hash link and verifying them using local JSON file storage.

## ⚙️ Features

- `POST /send`: Accepts an email, generates a random hash, and sends a confirmation link.
- `GET /verify/{hash}`: Checks whether the hash exists and deletes it if found.
- Local storage using a `data.json` file.
- No database required.
- Generic and reusable JSON storage system.

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/API-email-confirmation.git
cd API-email-confirmation
```

### 2. Configure .env file
```bash
EMAIL=youremail@gmail.com
PASSWORD=your_app_password
ADDRESS=smtp.gmail.com:587
PORT=8081
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the server
```bash
go run cmd/main.go
```

---

## 📬 Sending Email
### POST /send
```bash
POST http://localhost:8081/send
Content-Type: application/json

{
  "email": "user@example.com"
}
```
##### You will receive an email with a link like:
```bash
http://localhost:8081/verify/abc123hash
```

## ✅ Verifying
### GET /verify/{hash}
```bash
GET http://localhost:8081/verify/abc123hash
```
##### Returns true if the hash exists and was removed.
##### Returns false if the hash does not exist.
