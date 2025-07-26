from typing import List
from pyannote.audio import Pipeline

def diarize_speakers(file_path: str) -> List[str]:
    '''
    Диаризация с использованием pyannote.audio.
    '''
    pipeline = Pipeline.from_pretrained('pyannote/speaker-diarization', use_auth_token='your_token_here')
    diarization = pipeline(file_path)

    segments = []
    for turn, _, speaker in diarization.itertracks(yield_label=True):
        segments.append(f"{speaker}: {turn.start:.2f}s - {turn.end:.2f}s")
    return segments
