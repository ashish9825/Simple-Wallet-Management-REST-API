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

 go run main.go

 The server will start at: http://localhost:8080

#  Run with Docker

## Build Docker image
docker build -t wallet-api .

## Run container (basic)
docker run -p 8080:8080 wallet-api

# Sample cURL Commands

## Create User
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Ashish","email":"ashish@example.com"}'

## Create Wallet
curl -X POST http://localhost:8080/wallets \
  -H "Content-Type: application/json" \
  -d '{"user_id":1}'

## Check Wallet Balance
curl http://localhost:8080/wallets/1/balance

## Transfer Funds
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{"from_wallet_id":1,"to_wallet_id":2,"amount":50}'

## List Transactions
curl http://localhost:8080/wallets/1/transactions

# 👤 Author
### Ashish Chaudhary