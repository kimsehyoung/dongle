import requests
import time
import os

url = "http://localhost:10001/text/synthesize"

token = os.environ.get("JWT", "")

headers = {
    "Content-Type": "application/json"
}
if token:
    headers["Authorization"] = f"Bearer {token}"

data = {
    "language": "en",
    "text": "Hello test",
}

start_time = time.time()
for _ in range(1):
    response = requests.post(url, json=data, headers=headers)
    print(response.status_code, response.text)
    # time.sleep(1)

end_time = time.time()
total_time = end_time - start_time
print(f"TTS elapsed time: {total_time:.2f} seconds")