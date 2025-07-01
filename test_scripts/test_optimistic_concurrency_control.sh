#!/bin/bash

URL="http://localhost:8080/workouts/11"
TOKEN="NZU2NF2PO7VZHPXC3FDA2YPXLUQM4PUSCMFPGJJT5CPSUCHP5YRQ"
VERSION=2

DATA_TEMPLATE='{
  "title": "Concurrent Test TITLE",
  "description": "Test for concurrent modification conflict",
  "duration_minutes": 45,
  "calories_burned": 250,
  "version": '$VERSION',
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
}'

make_request() {
  local title=$1
  local data="${DATA_TEMPLATE//TITLE/$title}"
  echo "Sending request with title: $title"

  curl -s -o /tmp/response_$title.txt -w "%{http_code}" -X PUT "$URL" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d "$data" &
}

make_request "First"
make_request "Second"

wait

echo -e "\n--- Responses ---"
for file in /tmp/response_*.txt; do
  echo "$file:"
  cat "$file"
  echo -e "\n-----------------\n"
done
