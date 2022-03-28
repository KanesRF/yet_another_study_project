curl \
--cookie "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDg1MDYxMjYsInVzZXJuYW1lIjoiQWxleCJ9.jxpnZmRczSY9WXIrtw9O4EPcEVw-tFmBEzsTpSr8_Q4" \
-H "Content-Type: application/json" -X GET http://localhost:9090/

curl -H "Content-Type: application/json" -d '{"User":"Tom","Password":"12345"}' -X POST http://localhost:9090/register 