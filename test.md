curl -X POST -d '{ "first_name": "Parham", "id": "98243032", "last_name": "Alvani" }' -H 'Content-Type: application/json' 127.0.0.1:1373/create_user

curl localhost:1373/create_endpoint -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDExNzk2NzcsImZpcnN0X25hbWUiOiJQYXJoYW0iLCJpZCI6Ijk4MjQzMDMyIiwibGFzdF9uYW1lIjoiQWx2YW5pIn0.GXtkaGfFNO7-qrbviMluW3FXybyG8wlBYagyrShcvHs" -X POST -d '{ "url": "1325http://googles.com", "interval": 2, "threshold": 2 }' -H 'Content-Type: application/json'
