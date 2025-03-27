# Transaction API

This project is a simple CRUD API for managing transactions using Go and PostgreSQL. It allows users to create, read, update, and delete transactions with various attributes.

## Project Structure

```
transaction-api
├── cmd
│   └── main.go                  # Entry point of the application
├── config
│   └── config.go                # Configuration settings
├── internal
│   ├── controllers
│   │   └── transaction_controller.go  # HTTP request handlers for transactions
│   ├── models
│   │   └── transaction.go        # Transaction model definition
│   ├── repositories
│   │   └── transaction_repository.go # Database interaction methods
│   ├── routes
│   │   └── routes.go            # API route definitions
│   ├── services
│   │   └── transaction_service.go # Business logic for transactions
├── migrations
│   └── 0001_create_transactions_table.sql # SQL script for creating transactions table
├── go.mod                        # Module initialization file
├── go.sum                        # Dependency checksums
└── README.md                     # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd transaction-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the PostgreSQL database:**
   - Create a PostgreSQL database and user.
   - Update the database connection settings in `config/config.go`.

4. **Run migrations:**
   - Execute the SQL script in `migrations/0001_create_transactions_table.sql` to create the transactions table.

5. **Run the application:**
   ```go run cmd/main.go```

## API Usage

- **Create Transaction:** `POST /transactions`
- **Get Transaction:** `GET /transactions/{id}`
- **Update Transaction:** `PUT /transactions/{id}`
- **Delete Transaction:** `DELETE /transactions/{id}`

## Transaction Object Format

```json
{
  "id": "string",
  "justification": "string",
  "type": "investment | income | outcome",
  "category": "string",
  "start_date": "epoch_day",
  "end_date": "epoch_day",
  "frequency": "string",
  "amount": "float"
}
```

## License

This project is licensed under the MIT License.