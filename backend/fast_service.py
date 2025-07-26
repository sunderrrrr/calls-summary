from fastapi import FastAPI, UploadFile, Form
from fastapi.responses import FileResponse, JSONResponse
from pydantic import BaseModel
import os
import uuid
import tempfile
import requests

from app.whisper_service import transcribe_audio
from app.diarization import diarize_speakers
from app.gigachat_api import generate_summary
from app.report_generator import save_report

app = FastAPI()


class ReportRequest(BaseModel):
    session_id: str
    format: str  # 'txt', 'pdf', 'docx'


@app.post("/analyze")
async def analyze_audio(
    file: UploadFile,
    model: str = Form(default="gpt-4o-transcribe")
):
    """
    Эндпоинт для анализа аудио, совместимый с форматом OpenAI (multipart/form-data).
    """
    tmp_dir = tempfile.mkdtemp()
    file_path = os.path.join(tmp_dir, file.filename)

    with open(file_path, "wb") as f:
        f.write(await file.read())

    session_id = str(uuid.uuid4())

    # Шаг 1: ASR (распознавание речи)
    transcription = transcribe_audio(file_path)

    # Шаг 2: Диаризация
    diarized_segments = diarize_speakers(file_path)

    # Шаг 3: Резюме через GigaChat
    summary = generate_summary(transcription)

    # Сохраняем промежуточные результаты для генерации отчёта
    os.makedirs(f"/tmp/{session_id}", exist_ok=True)
    with open(f"/tmp/{session_id}/transcript.txt", "w", encoding="utf-8") as f:
        f.write(transcription)
    with open(f"/tmp/{session_id}/summary.txt", "w", encoding="utf-8") as f:
        f.write(summary)
    with open(f"/tmp/{session_id}/diarization.txt", "w", encoding="utf-8") as f:
        for d in diarized_segments:
            f.write(f"{d}\n")

    return JSONResponse({"session_id": session_id, "summary": summary})


@app.post("/report")
def generate_report(request: ReportRequest):
    path = f"/tmp/{request.session_id}"
    if not os.path.exists(path):
        return JSONResponse(status_code=404, content={"error": "Сессия не найдена"})

    with open(os.path.join(path, "summary.txt"), encoding="utf-8") as f:
        summary = f.read()

    # Проверяем поддерживаемые форматы
    allowed_formats = {"txt": "text/plain", "pdf": "application/pdf", "docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
    if request.format not in allowed_formats:
        return JSONResponse(status_code=400, content={"error": "Неподдерживаемый формат отчёта"})

    report_path = os.path.join(path, f"report.{request.format}")
    save_report(summary, report_path, request.format)

    return FileResponse(report_path, media_type=allowed_formats[request.format], filename=f"report.{request.format}")