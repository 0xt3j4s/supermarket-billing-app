# Supermarket Billing App

This is a simple Supermarket Billing App implemented using Go (server-side) and MongoDB (database). The application allows you to perform CRUD (Create, Read, Update, Delete) operations on bills, which represent customer purchases at a supermarket.



## Table of Contents
- [Pre-requisites](#pre-requisites)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
    - [Create a new bill](#create-a-new-bill)
    - [Retrieve a specific bill by ID](#retrieve-a-specific-bill-by-id)
    - [Retrieve all bills](#retrieve-all-bills)
    - [Update a specific bill by ID](#update-a-specific-bill-by-id)
    - [Delete a specific bill by ID](#delete-a-specific-bill-by-id)
- [Contributing](#contributing)




## Pre-requisites
Before running the application, ensure that you have the following installed:
- [Go](https://go.dev/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)




## Getting Started

1. Clone the repository:
    ```bash 
    git clone https://github.com/0xt3j4s/supermarket-billing-app.git
    ```
2. Navigate to the project directory:
    ```bash
    cd supermarket-billing-app
    ```
3. Initialize Go Modules:
    ```bash
    go mod init
    ```
4. Download the project dependencies:
    ```bash
    go mod tidy
    ```
5. Set up the MongoDB connection by updating the `databaseURL` variable in the `main.go` file with your MongoDB connection string.
6. Start the application:
    ```bash
    go run main.go
    ```

7. The application will start running on `http://localhost:8080`.



## API Endpoints
The following are the API endpoints are available:
- `POST /bills`: Create a new bill.
- `GET /bills/:id`: Retrieve a specific bill by ID.
- `GET /bills`: Retrieve all bills.

- `PUT /bills/:id`: Update a specific bill by ID.
- `DELETE /bills/:id`: Delete a specific bill by ID.

## Usage
Initial bill entries in the database:
![Initial Bill Entries](/output/initial_bills.png)


### Create a new bill
Send a POST request to `/bills`:
```shell
curl -X POST -H "Content-Type: application/json" -d '{
  "id": 3,
  "user_name": "William",
  "items": [
    {
      "id": 1,
      "name": "Item 1",
      "quantity": 2,
      "price": 10,
      "added_at": "2023-05-16"
    },
    {
      "id": 2,
      "name": "Item 2",
      "quantity": 3,
      "price": 15,
      "added_at": "2023-05-16"
    }
  ],
  "created_at": "2023-05-16"
}' http://localhost:8080/bills

```
![New Bill Entry](/output/create_bill.png)

### Retrieve a specific bill by ID
Send a GET request to `/bills/:id`:
```shell
curl -X GET http://localhost:8080/bills/3
```
![Retrieve Bill](/output/get_bill.png)


### Retrieve all bills
Send a GET request to `/bills`:
```shell
curl -X GET http://localhost:8080/bills
```
![Retrieve All Bills](/output/get_all_bills.png)

### Update a specific bill by ID
Send a PUT request to `/bills/:id`:
```shell
curl -X PUT -H "Content-Type: application/json" -d '{
  "id": 1,
  "user_name": "William Smith",
  "items": [
    {
      "id": 1,
      "name": "Updated Item 1",
      "quantity": 5,
      "price": 20,
      "added_at": "2023-05-16"
    },
    {
      "id": 2,
      "name": "Updated Item 2",
      "quantity": 2,
      "price": 10,
      "added_at": "2023-05-16"
    }
  ],
  "created_at": "2023-05-16"
}' http://localhost:8080/bills/3
```
![Update Bill](/output/update_bill.png)

### Delete a specific bill by ID
Send a DELETE request to `/bills/:id`:
```shell
curl -X DELETE http://localhost:8080/bills/3
```

![Delete Bill](/output/deleted_bill.png)
    

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please create an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for details.
