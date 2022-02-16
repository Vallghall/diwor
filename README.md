# DIWOR project
##Дипломная работа (DIplomWORk)

## Auth API endpoints
###Here are listed the endpoints for authentication
- GET  /auth/login
- GET  /auth/register
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

