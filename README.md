# IDN - Case Test

## Description
Case Study : Sudoku Solver

## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to do before running this apps :

1. Go Binary >= 1.16.x
2. MySQL (optional)

After install all required depedencies, follow this instruction :

1. Make sure you can do run "go mod download" or "go mod tidy" inside root of the project for downloading the depedencies lib
2. Configure your mysql credential inside .env file (optional)
3. Create database and please ensure that your database name in accordance with the .env file (optional)

### Running

Just run command "go run main.go" from root directory, then you can see your console output for monitoring purposes.

- Using cURL: (Note: make sure your input match with this example)
curl --location --request POST 'localhost:8080/sudoku/solver' \
--header 'Content-Type: application/json' \
--data-raw '[
    [0,3,0,0,0,0,8,0,0],
    [0,0,6,3,0,0,0,4,2],
    [2,0,8,6,7,0,3,0,5],
    [8,5,0,0,1,0,6,2,0],
    [0,0,7,0,0,0,9,0,0],
    [0,4,9,0,5,0,0,1,8],
    [9,0,5,0,4,7,2,0,6],
    [3,7,0,0,0,6,4,0,0],
    [0,0,1,0,0,0,0,7,0]
]'

- Or, or you can do import postman collection using file <b>IDN - Test Case (Backend).postman_collection.json</b> in the root directory.