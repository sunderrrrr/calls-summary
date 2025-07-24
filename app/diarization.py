from typing import List

def diarize_speakers(file_path: str) -> List[str]:
    '''
    Шаблон функции диаризации с использованием pyannote.audio.
    Здесь будет определение, кто когда говорил.
    '''
    # TODO: Подключить pyannote.audio и HuggingFace токен
    # from pyannote.audio import Pipeline
    # pipeline = Pipeline.from_pretrained('pyannote/speaker-diarization', use_auth_token='your_token_here')
    # diarization = pipeline(file_path)

    # Пример возвращаемого списка:
    return [
        "Speaker 1: 0.00s - 10.25s",
        "Speaker 2: 10.25s - 20.70s",
        "Speaker 1: 20.70s - 32.00s"
    ]
