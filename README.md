# About
an otp system using golang + fiber(web framework) + postgres + redis

# Flow
users send their number, an otp code will generate in backend and stored at redis for 2 minutes. afterward user sends number + code with a post request. code will check at the backend.
if the code was incorrect it returns "no no no" string :)))
if the code was correct it has 2 different way:
1- user already exist, in this case backend respond with a jwt.
2- user doesnt exist, in this case backend create user in database.

# How to run

### run locally
install golang first from golang website or your favorite package manager.
then clone the repo:
```git clone https://github.com/shimodii/dekamond-otp-challenge```
then enter the directory and install dependencies.
```
cd dekamond-otp-challenge
go mod tidy
```

congrats. you have installed the application dependencies.
now you need to install postgres and redis.
i recom the docker version of them, you can run them like this:

for postgress :
```
docker run --name my-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=dekamond -p 5432:5432 -d postgres:15
```

for redis:
```
docker run --name my-redis -p 6379:6379 -d redis:7
```

after that you can run the app with:
```
make run
```
or:
```
go run main.go
```

you can also build the code and run the built code like this:
```
go build main.go -o application
./application
```

### run with docker compose:
we have a docker compose file so you can do:
```
docker compose up --build
```
* golang website is filtered :DD, run you ghand shekan first

# Tech stack
we used golang and fiber
fiber for easier and faster web development.

### why we used postgres?
mysql is moreeeee lightweight than oracle or ms-sql and its alternatives.
so mysql is better for our need, but we have postgres :)))
postgres is more efficient and faster in cuncurrent read and write.
so postgres is better.

### why we used redis?
we had to store the otp codes, they was not percistant data so they should handled in memory.
we could save them in a vector or some datastructures like that. but that was messy.
so we used redis. an in-memmory database with some functionalities like expire time :)))
