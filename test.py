# test.py
import requests
from dotenv import load_dotenv
import os

load_dotenv()

URI = os.getenv('URI', 'http://localhost:31020')
if URI[-1] == '/':
    URI = URI[:-1]

r=requests.post(
    f'{URI}/exec',
    json={"command": "pwd"}
)

print(r.text)
