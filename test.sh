#just some testing requests

curl -H "Content-Type: application/json" \
--cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDg2ODQwMTIsInVzZXJuYW1lIjoiVG9tIn0.aCOf0p2q61xhidOwEFVAC94bbqtcxfuu_KNV3lj3qPk" \
-X GET http://localhost:9090/

curl -H "Content-Type: application/json" -d '{"User":"Tom","Password":"12345"}' -X POST http://localhost:9090/register 

curl -H "Content-Type: application/json" -d '{"User":"Tom","Password":"12345"}' -X POST http://localhost:9090/login 

curl \
-H "Content-Type: application/json" \
--cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDg5MTE3MzQsInVzZXJuYW1lIjoiVG9tIn0.0VBAnOWdbOQMMpFabWEJd1RDcJRyiW6QZ484KCR19ds" \
-X POST http://localhost:9090/logout 

curl -X POST http://localhost:9090/logout 