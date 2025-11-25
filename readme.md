# Go Blog API

A RESTful blog API built with Go and the Gin framework, backed by MongoDB.

## 🚀 Features

- **RESTful API** with clean endpoints
- **MongoDB Integration** with MongoDB Atlas
- **Pagination Support** for blog listings
- **Modular Architecture** with separated concerns
- **Environment Configuration** with .env support
- **Error Handling** and proper HTTP status codes
- **CORS Ready** with trusted proxy configuration

## 📁 Project Structure

```
go-blog/
├── main.go              # Application entry point
├── config/
│   └── config.go        # Environment variables and configuration
├── database/
│   └── database.go      # MongoDB connection management
├── handlers/
│   └── handlers.go      # HTTP request handlers
├── middleware/
│   └── middleware.go    # Custom middleware
├── routes/
│   └── routes.go        # Route configuration
├── static/
│   └── style.css        # Static CSS assets
├── templates/           # HTML templates
│   ├── index.html       # Home page template
│   └── 404.html         # 404 error page
├── .env                 # Environment variables (not in git)
├── go.mod               # Go module dependencies
└── readme.md            # This file
```

## 🛠️ Tech Stack

- **Backend**: Go 1.16
- **Web Framework**: Gin v1.8.1
- **Database**: MongoDB (MongoDB Atlas)
- **Environment**: godotenv for .env management
- **Driver**: mongo-driver for MongoDB

## 📋 API Endpoints

### Blog Management
- `GET /api/v1/blogs` - Get all blogs with pagination
- `GET /api/v1/blog/:id` - Get a single blog by ID
- `POST /api/v1/seed` - Seed sample blog data (for testing)

### System
- `GET /` - Home page (HTML)
- `GET /ping` - Health check endpoint
- `GET /api/v1/movies` - Get movies from MongoDB

### Query Parameters

#### Get Blogs (`GET /api/v1/blogs`)
- `page` (int, default: 1) - Page number for pagination
- `limit` (int, default: 10) - Number of items per page

**Example Response:**
```json
{
  "blogs": [
    {
      "id": "6801b459916ffa9fa5ff06fd",
      "title": "Note",
      "content": "Blog content here...",
      "author": "John Doe",
      "created_at": "2025-04-18T02:09:29.475Z",
      "updated_at": "0001-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 119,
    "pages": 12
  }
}
```

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher
- MongoDB database (local or MongoDB Atlas)
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-blog
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your MongoDB connection string
   ```

4. **Configure MongoDB**
   Create a `.env` file with your MongoDB connection:
   ```env
   MONGODB_URL=mongodb+srv://username:password@cluster.mongodb.net/database_name?retryWrites=true&w=majority
   ```
   
   For local MongoDB:
   ```env
   MONGODB_URL=mongodb://localhost:27017
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 🧪 Testing

### Health Check
```bash
curl http://localhost:8080/ping
```

### Get Blogs
```bash
# Get first page with default limit
curl http://localhost:8080/api/v1/blogs

# Get specific page with custom limit
curl "http://localhost:8080/api/v1/blogs?page=2&limit=5"
```

### Get Single Blog
```bash
curl http://localhost:8080/api/v1/blog/{blog-id}
```

### Seed Sample Data
```bash
curl -X POST http://localhost:8080/api/v1/seed
```

## 📊 Database Schema

### Blog Collection
```go
type Blog struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title     string             `json:"title" bson:"title"`
    Content   string             `json:"content" bson:"content"`
    Author    string             `json:"author" bson:"author"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Description | Required |
|----------|-------------|----------|
| `MONGODB_URL` | MongoDB connection string | Yes |

## 🚦 Development

### Project Status
- [x] ✅ MongoDB connection
- [x] ✅ Blog struct definition
- [x] ✅ Get list of blogs (with pagination)
- [x] ✅ Get single blog by ID
- [x] ✅ Modular architecture
- [x] ✅ Error handling
- [x] ✅ Environment configuration
- [ ] 🔄 Create blog endpoint
- [ ] 🔄 Update blog endpoint
- [ ] 🔄 Delete blog endpoint
- [ ] 🔄 User authentication
- [ ] 🔄 Input validation
- [ ] 🔄 Unit tests
- [ ] 🔄 Docker support

### Running in Development Mode
```bash
# For development with debug output
export GIN_MODE=debug
go run main.go
```

### Production Deployment
```bash
# Build for production
go build -o blog-app main.go

# Run production binary
./blog-app
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📞 Contact

[Your Name] - [Your Email] - [Your GitHub Profile]

---

**Note**: This is a personal project for learning Go web development with Gin and MongoDB.