# 📘 Task Management API Documentation

Base URL: `http://localhost:8080`

## 📌 User Endpoints
This API uses JWT-based authentication. All protected routes require the Authorization header in this format:
- Authorization: Bearer <your_token_here>

 ## 📌 User Roles
- Admin: Can create, update, delete tasks, and promote users.
- Regular User: Can view tasks only.

### Register User
 **POST /register**

**Request Body:**

```
{
  "username": "johndoe",
  "password": "securepassword123"
}
```

- 📌 If it's the first user, they will be assigned the admin role automatically.

**Responses:**

- 201 Created: User registered successfully.
- 400 Bad Request: Username or Email already exists

### 🔹 Login
**POST /login**

**Request Body:**
```
{
  "username": "johndoe",
  "password": "securepassword123"
}
```

**Response:**
```
{
  "token": "jwt_token_here"
}
```
- 401 Unauthorized: Invalid credentials.

### 🔹 Promote User to Admin
**PUT /promote**
- 🔒 Protected — Requires Admin Role

**Request:**
```
  {
  "user_id": "687e2a42d40a9b07d466f534",
  "new_role": "admin"
}
```

**Response:**

- 200 OK: User promoted successfully.
- 403 Forbidden: admin access required.
- 404 Not Found: User not found.




## 📌 Task management Endpoints
🔐 Authorization Required:
- All task endpoints require a valid JWT.
- Only Admins can create, update, delete.
- All authenticated users can view tasks.

---

### 🔹 Create Task

**POST** `/tasks`

**Request Body:**
```
{
  "id": "1",
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "due_date": "2025-07-25T10:00:00Z",
  "status": "pending"
}
```
**Responses**
- 201 Created: Task created successfully
- 400 Internal Server Error

### 🔹 Get All Tasks

**GET** `/tasks`

**Responses**

- 200 OK: List of tasks
- 200 OK + message: "No Task in DB" if list is empty
- 500 Internal Server Error

### 🔹 Get Task by ID

**GET** `/tasks/:id`

**Responses**

- 200 OK: Task details
- 500: Internal Server Error

### 🔹 Update Task
**PUT** `/tasks/:id`

**Request Body**

```
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-08-01T08:00:00Z",
  "status": "completed"
}
```

**Responses**

- 200 OK: Task updated
- 400 Bad Request: Invalid JSON
- 404 Not Found: Task not found

### 🔹 Delete Task

**DELETE** `/tasks/:id`

**Responses**

- 200 OK: "Task deleted Successfully"
- 404 Not Found: Task not found

### 🧪 Testing Tools
**You can test the endpoints using:**
- Postman (recommended)
- VS Code REST Client (.http or .rest file)
- curl / Invoke-RestMethod



