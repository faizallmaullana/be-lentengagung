# Register

### Endpoint

		POST /register

### Request Body
```json
{
	"nik": "string",
	"email": "string",
	"phone": "string",
	"password": "string"
}
```

### Success Response
- Status: 201 Created
```json
{
	"id": "string",
	"nik": "string",
	"email": "string",
	"phone": "string",
	"created_at": "2025-12-27T00:00:00Z"
}
```

### Error Response
- Status: 400 Bad Request
```json
{
	"error": "error message"
}
```
