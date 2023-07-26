### How to run server:

1. Run docker or install if you do not have one
2. Clone the repository
3. Go to server folder with command `cd server`
4. Add .env `SECRET_KEY=[anything]` with command `touch .env`
5. Create postgres container with `make postgresinit`
6. Create database inside container `make createdb`
7. Migrate database inside container `make dbup`
8. Run server with `go run main.go`

### How to run client:

1. Clone the repository
2. Go to client folder with command `cd client`
3. Add .env `VITE_PASSWORD=[anything]` with command `touch .env`
4. Start the client with command `npm run dev`

### How to test server:

1. Run docker or install if you do not have one
2. Clone the repository
3. Go to server folder with command `cd server`
4. Add .env `SECRET_KEY=[anything]` with command `touch .env`. If you have done this previously, **you may skip this one**
5. Create postgres container with `make postgresinit` or if you have done this previously, **you may skip this one**
6. Create test database inside container `make createdbtest`
7. Migrate test database inside container `make dbtestup`
8. Test with test runner or `go test ./test/ -v`
