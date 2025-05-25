# Go E-Commerce API

A RESTful e-commerce API built with Go, featuring user authentication, product management, and order processing capabilities.

## 🚀 Features

- **User Management**
  - User registration and authentication
  - JWT-based authorization
  - Secure password hashing

- **Product Management**
  - Create, read, and manage products
  - Product inventory tracking
  - Image URL support

- **Order Processing**
  - Shopping cart functionality
  - Order checkout and management
  - Inventory validation during checkout

- **API Documentation**
  - Swagger/OpenAPI documentation
  - Interactive API explorer

## 🛠️ Tech Stack

- **Language**: Go 1.20+
- **Web Framework**: Gorilla Mux
- **Database**: MySQL with GORM ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Documentation**: Swagger/OpenAPI
- **Password Hashing**: bcrypt

## 📁 Project Structure

```
go-ecom/
├── bin/                    # Compiled binaries
├── cmd/
│   ├── app/               # Application initialization
│   ├── config/            # Environment configuration
│   ├── controllers/       # HTTP request handlers
│   ├── db/                # Database connection
│   ├── docs/              # Swagger documentation
│   ├── middlewares/       # HTTP middlewares
│   ├── models/            # Database models
│   ├── routes/            # Route definitions
│   ├── types/             # Request/response types
│   ├── utils/             # Utility functions
│   ├── validations/       # Input validation
│   └── main.go            # Application entry point
├── example.env            # Environment variables template
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── Makefile              # Build and run commands
└── run.sh                # Run script
```

## 🚦 Getting Started

### Prerequisites

- Go 1.20 or higher
- MySQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/ZUBERKHAN034/go-ecom.git
   cd go-ecom
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp example.env .env
   ```
   
   Edit `.env` file with your configuration:
   ```env
   # Server
   JWT_SECRET=your_jwt_secret_key_here
   PORT=8080
   BASE_URL=http://localhost:8080

   # Database
   DB_PUBLIC_HOST=localhost:3306
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_CA_CERT=your_db_ca_certificate
   ```

4. **Set up MySQL database**
   - Create a MySQL database
   - Update the database configuration in your `.env` file
   - The application will auto-migrate the required tables

### Running the Application

**Using Makefile:**
```bash
# Build and run
make run

# Build only
make build

# Run tests
make test

# Run with hot reload (requires air)
make run-dev

# Generate Swagger docs
make run-swagger
```

**Using run script:**
```bash
./run.sh
```

**Manual run:**
```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## 📖 API Documentation

### Interactive Documentation
Visit `http://localhost:8080/swagger/` for interactive API documentation.

### API Endpoints

#### Authentication
- `POST /user/register` - Register a new user
- `POST /user/login` - Login user

#### Products
- `GET /products` - Get all products
- `GET /product/{id}` - Get product by ID
- `POST /product` - Create new product (requires authentication)

#### Orders
- `POST /order/checkout` - Checkout order (requires authentication)

### Authentication

Most endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

### Request/Response Examples

#### User Registration
```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword123",
    "address": "123, Street Name, City, Country"
  }'
```

#### User Login
```bash
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword123"
  }'
```

#### Create Product
```bash
curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "name": "Wireless Headphones",
    "description": "High-quality wireless headphones with noise cancellation.",
    "price": 99.99,
    "quantity": 50,
    "image": "https://example.com/images/headphones.jpg"
  }'
```

#### Checkout Order
```bash
curl -X POST http://localhost:8080/order/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{
    "orderItems": [
      {
        "productId": 1,
        "quantity": 2
      }
    ]
  }'
```

## 🗄️ Database Schema

The application uses the following main entities:

- **Users**: User account information
- **Products**: Product catalog with inventory
- **Orders**: Order records
- **OrderItems**: Individual items within orders

## 🔧 Configuration

Environment variables can be configured in the `.env` file:

| Variable | Description | Default |
|----------|-------------|---------|
| `JWT_SECRET` | Secret key for JWT token signing | Required |
| `PORT` | Server port | 8080 |
| `BASE_URL` | Base URL for the API | http://localhost:8080 |
| `DB_PUBLIC_HOST` | Database host | Required |
| `DB_USER` | Database username | Required |
| `DB_PASSWORD` | Database password | Required |
| `DB_NAME` | Database name | Required |
| `DB_CA_CERT` | Database CA certificate (optional) | - |

## 🧪 Testing

Run the test suite:
```bash
make test
# or
go test ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Zuber Khan**
- GitHub: [@ZUBERKHAN034](https://github.com/ZUBERKHAN034)

## 🙏 Acknowledgments

- Built with Go and the amazing Go ecosystem
- Swagger for API documentation
- GORM for database operations
- Gorilla Mux for HTTP routing

---

⭐ If you found this project helpful, please give it a star!
