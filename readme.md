# Simple Token-Bucket Rate Limiter

### Run the Application

```bash
$ go run main.go

```
### Test it, from postman
Post:localhost:9000/users/signup

json body:{
    "first_name": "my_name",
    "last_name": "my_surname",
    "password": "moh12345",
    "email": "myemail@gmail.com",
    "phone": "000000111",
    "user_type": "ADMIN"
}


Post:localhost:9000/users/login

json body:{
    "email": "myemail@gmail.com",
    "password": "moh12345"
}
