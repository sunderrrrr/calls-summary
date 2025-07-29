# Transcriber Backend

Сервис для автоматической транскрипции, диаризации и генерации отчётов по аудиофайлам с использованием FastAPI, Whisper, pyannote.audio и GigaChat.

## 📚 Стек технологий

- **Python 3.9+**
- **FastAPI** — REST API
- **Whisper** — автоматическая транскрипция аудио
- **pyannote.audio** — диаризация (выделение спикеров)
- **GigaChat API** — генерация краткого содержания
- **python-docx** — генерация DOCX отчётов
- **fpdf** — генерация PDF отчётов
- **requests** — HTTP-запросы

## 📂 Структура проекта

```
backend/gigachat/
├── fast_service.py         # Основной API FastAPI
├── whisper_service.py      # Транскрипция аудио через Whisper
├── diarization.py          # Диаризация через pyannote.audio
├── gigachat_api.py         # Генерация резюме через GigaChat
├── report_generator.py     # Генерация отчётов (txt, pdf, docx)
└── README.md               # Документация
```

## ⚙️ Установка

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/yourusername/transcriber-backend.git
    cd transcriber-backend/backend/gigachat
    ```

2. Установите зависимости:

    ```sh
    pip install -r requirements.txt
    ```

    Пример содержимого `requirements.txt`:
    ```
    fastapi
    uvicorn
    whisper
    pyannote.audio
    python-docx
    fpdf
    requests
    ```

3. Установите [ffmpeg](https://ffmpeg.org/download.html) и убедитесь, что он доступен в PATH.

## Настройка

- В файле [`diarization.py`](diarization.py) замените `'your_token_here'` на ваш токен HuggingFace.
- В файле [`gigachat_api.py`](gigachat_api.py) замените `'YOUR_GIGACHAT_TOKEN'` на ваш токен GigaChat.

## Запуск

Запустите сервер FastAPI:

```sh
uvicorn fast_service:app --reload
```

## Использование

### 1. Анализ аудиофайла

`POST /analyze`

- Формат запроса: `multipart/form-data`
- Параметры:
    - `file`: аудиофайл
    - `model`: (опционально) модель для транскрипции

Ответ:
```json
{
  "session_id": "uuid",
  "summary": "Краткое содержание"
}
```

### 2. Генерация отчёта

`POST /report`

- Формат запроса: JSON
- Параметры:
    - `session_id`: идентификатор сессии
    - `format`: `txt`, `pdf` или `docx`

Ответ: файл отчёта.

## Пример запроса через curl

```sh
curl -F "file=@audio.wav" http://localhost:8000/analyze
```

```sh
curl -X POST -H "Content-Type: application/json" \
     -d '{"session_id": "ваш_id", "format": "pdf"}' \
     http://localhost:8000/report --output report.pdf
```
