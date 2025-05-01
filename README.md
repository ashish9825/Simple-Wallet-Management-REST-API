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

#  API Endpoints

 - User
   POST /users
   Create a new user
   Body:
   {
        "name": "Ashish",
        "email": "ashish@example.com"
   }
