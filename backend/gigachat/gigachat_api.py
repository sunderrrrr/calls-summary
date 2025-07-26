import requests

def generate_summary(text: str) -> str:
    url = "https://gigachat.example.com/api/v1/summarize"
    headers = {
        "Authorization": "Bearer YOUR_GIGACHAT_TOKEN",
        "Content-Type": "application/json"
    }
    data = {
        "text": text,
        "lang": "ru"
    }
    response = requests.post(url, json=data, headers=headers, timeout=60)
    response.raise_for_status()
    return response.json().get("summary", "")