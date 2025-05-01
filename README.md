# Simple Wallet Management REST API

A simple RESTful backend service built with **Go**, **SQLite**, **GORM**, and **Docker** to manage user wallets and basic transactions between them.

## Features

-  Create a user
-  Create a wallet for the user
-  Check wallet balance
-  Transfer funds between wallets
-  List all transactions for a wallet
-  Input validation and error handling
-  Dockerized setup
-  Environment variable-based configuration

## Tech Stack 

   Technology           Description                   
     
 - Go (Golang)          Backend language               
 - net/http + Gin       HTTP framework (if used)       
 - SQLite               Lightweight SQL database       
 - GORM                 ORM for Go                     
 - Docker               Containerization      

 # Run the server
```
 go run main.go
```

The server will start at: http://localhost:8080

#  Run with Docker

## Build Docker image
```
docker build -t wallet-api .
```

## Run container (basic)
```
docker run -p 8080:8080 wallet-api
```

# API Endpoints 
### 👤 Users

- **POST** `/users`  
  ➤ Create a new user  
  **Request Body**:
  ```json
  {
    "name": "Ashish",
    "email": "ashish@example.com"
  }

### 👤 Wallets

- **POST** `/wallets`  
  ➤  Create a wallet for a user
  **Request Body**:
  ```json
  {
    "user_id": 1,
    "balance": 1000
  }

- **POST** `/transactions`  
  ➤ Transfer funds from one wallet to another
  **Request Body**:
  ```json
  {
    "from_wallet_id": 1,
    "to_wallet_id": 2,
    "amount": 100
  }

# Sample cURL Commands

## Create User
### Run the below command on terminal in the project directory 
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Ashish","email":"ashish@example.com"}'
```
## Create Wallet for the User 
### Run the below command on terminal in the project directory
```
curl -X POST http://localhost:8080/wallets \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "balance": 1000
  }'
```

## Check Wallet Balance
### Run the below command on terminal in the project directory
```
curl http://localhost:8080/wallets/1/balance
```

## Transfer Funds
### Run the below command on terminal in the project directory
```
curl -X POST http://localhost:8080/wallets/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "from_wallet_id": 1,
    "to_wallet_id": 2,
    "amount": 200
  }'
```

## List Transactions
### Run the below command on terminal in the project directory
```
curl http://localhost:8080/wallets/1/transactions
```

# 👤 Author
### Ashish Chaudhary