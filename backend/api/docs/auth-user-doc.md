### 🧾 Общие заголовки:
```
Authorization: Bearer <token>  // Только для защищённых маршрутов
Content-Type: application/json
```


---

## 🔐 Auth

### 🔸 `POST /api/v1/auth/sign-up`

**Назначение:** Регистрация нового пользователя.

**Тело запроса:**
```json
{
  "name": "John",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Успешный ответ:**
```json
{
  "id": 1
}
```

**Ошибки:**
- `400 Bad Request` — `field validation fail` (некорректные данные)
- `400 Bad Request` — `user already exists` (пользователь с таким email уже зарегистрирован)

---

### 🔸 `POST /api/v1/auth/sign-in`

**Назначение:** Авторизация пользователя.

**Тело запроса:**
```json
{
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Успешный ответ:**
```json
{
  "token": "jwt-token-string"
}
```

**Ошибки:**
- `400 Bad Request` — `field validation fail` (ошибка валидации тела запроса)
- `400 Bad Request` — `user don't exist` (неверный email или пароль)

---

### 🔸 `POST /api/v1/auth/forgot`

**Назначение:** Запрос на сброс пароля (отправка токена на почту или другой канал).

**Тело запроса:**
```json
{
  "login": "user@example.com"
}
```

**Ответ:**
```json
{
  "message": "reset link sent to your email"
}
```

**Ошибки:**
- `400 Bad Request` — ошибка валидации

---

### 🔸 `POST /api/v1/auth/reset`

**Назначение:** Сброс пароля по полученному токену.

**Тело запроса:**
```json
{
  "token": "abcdef123456",
  "new_password": "NewSecurePass123"
}
```

**Ответ:**
```json
{
  "message": "password reset successfully"
}
```

**Ошибки:**
- `400 Bad Request` — ошибка валидации или недействительный токен
