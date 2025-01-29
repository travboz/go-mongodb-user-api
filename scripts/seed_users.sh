#!/bin/bash

# Define the API endpoint
url="http://localhost:8080/users"

# Define the users to insert
users=(
  '{"name": "Alice Johnson", "email": "alice.johnson@example.com", "favourite_number": 42, "active": true}'
  '{"name": "David Smith", "email": "david.smith@example.com", "favourite_number": 17, "active": false}'
  '{"name": "Emma Wilson", "email": "emma.wilson@example.com", "favourite_number": 8, "active": true}'
  '{"name": "James Carter", "email": "james.carter@example.com", "favourite_number": 25, "active": false}'
  '{"name": "Sophia Brown", "email": "sophia.brown@example.com", "favourite_number": 13, "active": true}'
  '{"name": "Michael Green", "email": "michael.green@example.com", "favourite_number": 9, "active": false}'
  '{"name": "Olivia Lee", "email": "olivia.lee@example.com", "favourite_number": 33, "active": true}'
  '{"name": "Ethan Walker", "email": "ethan.walker@example.com", "favourite_number": 19, "active": true}'
  '{"name": "Mia Harris", "email": "mia.harris@example.com", "favourite_number": 5, "active": false}'
  '{"name": "Liam Clark", "email": "liam.clark@example.com", "favourite_number": 11, "active": true}'
)

# Loop through each user and insert them using curl
for user in "${users[@]}"; do
  curl -X POST "$url" \
    -H "Content-Type: application/json" \
    -d "$user"
  echo -e "\nUser inserted: $user\n"
done
