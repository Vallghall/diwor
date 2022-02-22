# DIWOR project
## Дипломная работа (DIplomWORk)

## Auth API endpoints
### Here are listed the endpoints for authentication
- GET  /auth/login
- GET  /auth/register
- GET /auth/logout
- POST /auth/sign-up

Запрос должен иметь заголовок
Content-Type со значением application/json,
а тело запроса должно иметь следующий вид:
```
{
    "name"     : "your_name"
    "username" : "your_username"
    "password" : "your_password"
}
```
- POST /auth/sign-in

Запрос должен иметь заголовок
Content-Type со значением application/json,
а тело запроса должно иметь следующий вид:
```
{
    "username" : "your_username"
    "password" : "your_password"
}
```
После выполнения валидного запроса
сервер создаст Cookie в вашем браузере.

## Main API endpoints
### Here are listed the main endpoints of the application
### One needs to have a token cookie to call on them

- GET /api/profile/
- GET /api/experiment/
- GET /api/experiment/hashes
- GET /api/experiment/ciphers
- GET /api/experiment/hash-results
- GET /api/experiment/cipher-results
- POST /api/experiment/start-hash-experiment
```
{
    "algorithms" : ["SHA-512", "MD5", "RIPEMD-160"] # strings array with hash algorithms names
}
```
- POST /api/experiment/start-cipher-experiment
```
{
    "algorithms" : ["Кузнечик"] # strings array with cipher algorithms names
}
```