**Approval Endpoint**

**Purpose**: Confirm a user's registration by validating the registration JWT payload against the submitted email and, on success, set `users.approved_at`.

**Endpoint**: `POST /api/auth/approve`

- **Request Content-Type**: `application/json`

- **Request Body (JSON)**:

```json
{
  "token": "<registration_jwt>",
}
```

- **Behavior**:
  - The server validates the `token` (registration JWT) and extracts the payload.
  - If the token's `email` claim matches the submitted `email`, the server updates `users.approved_at` to the current time for that user.
  - If the token is invalid or the payload does not match the supplied email, the request is rejected.

- **Success Response (200)**:

```json
{
  "message": "approved"
}
```

- **Error Responses**:
  - `400 Bad Request` — invalid JSON, missing fields, token validation failure, or token payload/email mismatch.
  - Other error responses may be returned for unexpected server errors.

- **Example curl**:

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"token":"<registration_jwt>","email":"user@example.com"}' \
  http://localhost:8080/api/auth/approve
```

**Notes**:
- The registration token is produced by `JWTService.CreateRegistrationToken` during registration.
- The handler expects the `token` and `email` fields in the request body and returns a simple success message on approval.
