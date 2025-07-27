### 🧾 Общие заголовки:
```
Authorization: Bearer <token> 
Content-Type: application/json
```

---

## 📊 Analysis

### 🔸 `POST /api/v1/analysis`

**Назначение:** Загрузка файла для анализа.

**Тело запроса:**
- Формат: `multipart/form-data`
- Поле: `file` — файл для анализа.

**Успешный ответ:**
```json
{
  "id": "analysis-id"
}
```

**Ошибки:**
- `400 Bad Request` — `file not found` (файл не передан)
- `401 Unauthorized` — `unauthorized` (неавторизованный пользователь)
- `500 Internal Server Error` — `failed to process report` (ошибка обработки файла)

---

### 🔸 `GET /api/v1/analysis/:analysisId/chat`

**Назначение:** Получение истории чата для анализа.

**Успешный ответ:**
```json
{
  "messages": [
    {
      "id": "message-id",
      "analysis_id": "analysis-id",
      "sender": "user",
      "message": "Когда были установлены дедлайны задачи?",
      "created_at": "2023-01-01T12:00:00Z"
    },
    {
      "id": "message-id",
      "analysis_id": "analysis-id",
      "sender": "bot",
      "message": "До 28 июля",
      "created_at": "2023-01-01T12:01:00Z"
    }
  ]
}
```

**Ошибки:**
- `401 Unauthorized` — `unauthorized` (неавторизованный пользователь)
- `500 Internal Server Error` — `failed to get chat history` (ошибка получения истории чата)

---

### 🔸 `POST /api/v1/analysis/:analysisId/chat`

**Назначение:** Отправка сообщения в чат.

**Тело запроса:**
```json
{
  "sender": "user",
  "message": "Расскажи о целях встречи"
}
```

**Успешный ответ:**
```json
{
  "message": "message sent successfully"
}
```

**Ошибки:**
- `400 Bad Request` — `field validation fail` (ошибка валидации тела запроса)
- `401 Unauthorized` — `unauthorized` (неавторизованный пользователь)
- `500 Internal Server Error` — `failed to send message` (ошибка отправки сообщения)
