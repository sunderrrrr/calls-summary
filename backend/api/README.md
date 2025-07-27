# Calls-Summary REST API 🚀

---

## Стек технологий 🛠️

- **Язык:** Go 🐹
- **Фреймворк:** Gin 🍃
- **База данных:** PostgreSQL 🐘
- **Миграции:** migrate (golang-migrate) 🔄
- **Контейнеризация:** Docker 🐳

---

## Установка и запуск ⚙️

### 1. Запуск PostgreSQL в Docker 🐘

```bash
docker run --name=calls-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
```

Это создаст и запустит контейнер с PostgreSQL, доступным на порту `5436`. 🔥

---

### 2. Применение миграций 🔄

Убедитесь, что у вас установлен инструмент [migrate](https://github.com/golang-migrate/migrate). ✔️

Запустите миграции командой:

```bash
migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" up
```

---

### 3. Запуск сервера 🚀

После применения миграций запустите сервер командой:

```bash
go run main.go
```

---

## Описание API 📋

## Аутентификация

### Регистрация пользователя (Sign Up)

**URL:** `/api/v1/auth/sign-up`  
**Метод:** `POST`  
**Описание:** Создаёт нового пользователя с указанным именем, email и паролем.

**Тело запроса (JSON):**

```json
{
  "name": "string",       // Имя пользователя (обязательно)
  "email": "string",      // Email пользователя (обязательно, уникальный)
  "password": "string"    // Пароль пользователя (обязательно)
}
```

**Пример запроса:**

```json
{
  "name": "Иван Иванов",
  "email": "ivan@example.com",
  "password": "strongpassword123"
}
```

**Успешный ответ:**

- Код: `200 Created`
- Тело: 
```json
{
  "id": "0"
}
```

**Ошибки:**

- `400 Bad Request` — если отсутствуют обязательные поля или формат неправильный
- `500 Conflict` — если email уже зарегистрирован или ошибка на сервере

---

### Вход пользователя (Sign In)

**URL:** `/api/v1/auth/sign-in`  
**Метод:** `POST`  
**Описание:** Авторизует пользователя по email и паролю, возвращает токен bearer.

**Тело запроса (JSON):**

```json
{
  "email": "string",      // Email пользователя (обязательно)
  "password": "string"    // Пароль пользователя (обязательно)
}
```

**Пример запроса:**

```json
{
  "email": "ivan@example.com",
  "password": "strongpassword123"
}
```

**Успешный ответ:**

- Код: `200 OK`
- Тело:

```json
{
  "token": "jwt-token-string"
}
```

**Ошибки:**

- `400 Bad Request` — если отсутствуют обязательные поля
- `500 Unauthorized` — если email или пароль неверны или ошибка на сервере

---
