import threading
import requests
import json

URL = "http://localhost:8080/workouts/11"
TOKEN = "NZU2NF2PO7VZHPXC3FDA2YPXLUQM4PUSCMFPGJJT5CPSUCHP5YRQ"
VERSION = 4

headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Content-Type": "application/json"
}

base_payload = {
    "title": "",
    "description": "Test for concurrent modification conflict",
    "duration_minutes": 45,
    "calories_burned": 250,
    "version": VERSION,
    "entries": [
        {
            "exercise_name": "Walking",
            "sets": 1,
            "duration_seconds": 2700,
            "weight": 0,
            "notes": "Keep a steady pace",
            "order_index": 1
        }
    ]
}


def update_workout(title):
    payload = base_payload.copy()
    payload["title"] = title
    response = requests.put(URL, headers=headers, json=payload)
    print(f"[{title}] Status: {response.status_code}")
    print(f"[{title}] Body: {response.text}\n")


t1 = threading.Thread(target=update_workout, args=("First Update",))
t2 = threading.Thread(target=update_workout, args=("Second Update",))

t1.start()
t2.start()

t1.join()
t2.join()
