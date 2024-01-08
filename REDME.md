```bash:
$ docker-compose up
$ make migrateup

$ go run main.go

$ grpcurl -plaintext 0.0.0.0:8080 list
$ grpcurl -plaintext 0.0.0.0:8080 list user.UserService
$ grpcurl -plaintext -d '{"username": "test_user", "email": "test@example.com", "password": "password"}' 0.0.0.0:8080 user.UserService.Signup
$ grpcurl -plaintext -d @ 0.0.0.0:8080 user.UserService.Signup
  {"username": "test_user", "email": "test@example.com", "password": "password"}
  Ctrl+Z (UnixはCtrl+D)
$ grpcurl -plaintext -d @ 0.0.0.0:8080 user.UserService.Login
  {"email": "test@example.com", "password": "password"}
  Ctrl+Z (UnixはCtrl+D)

```
