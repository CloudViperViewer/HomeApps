
tags: [[Go Database API]] [[API's]] [[endpoints]]



# 🧩 API Endpoint: `/api/select`

tag: [[/api/select]]

**Method**: `POST`  
**Content-Type**: `application/json`  
**Description**: Dynamically queries a database table and returns filtered, paginated results.

---

## 📨 Request Body

```json
{
  "table": "Bank",               // Required: Table name (e.g., "Bank", "Account")
  "fields": ["BankName"],        // Optional: List of field names to include in result
  "logicalExpression": {         // Optional: Filters and conditions (AND/OR)
    "operator": "AND",
    "filters": [
      { "operator": "=", "field": "BankID", "value": [1] }
    ],
    "logicalExpressions": []
  },
  "pagingInfo": {
    "startIndex": 1,             // Required: 1-based offset
    "batchSize": 10              // Required: -1 for all, > 0 for limit
  }
}
```

---

## ✅ Example Request

```http
POST /api/select HTTP/1.1
Content-Type: application/json

{
  "table": "Account",
  "fields": ["AccountName", "Balance"],
  "logicalExpression": {
    "operator": "AND",
    "filters": [
      { "operator": ">", "field": "Balance", "value": [1000] }
    ],
    "logicalExpressions": []
  },
  "pagingInfo": {
    "startIndex": 1,
    "batchSize": 5
  }
}
```

---

## 🔁 Response

### ✅ Success (200 OK)

```json
{
  "success": true,
  "data": [
    {
      "accountName": "MyAccount",
      "balance": 3500.0
    },
    ...
  ]
}
```

### ❌ Error (400/404/500)

```json
{
  "error": "bad_request",
  "message": "table missing"
}
```

---

## 🧠 Internal Logic

### Core Functions

- `dbQuerySelect`: Main handler, validates input and returns results.
- `confirmData`: Checks required fields (`table`, `startIndex`, `batchSize`).
- `queryDb`: Prepares query using reflection and executes it.

### Logical Expression Support

Nested `AND`/`OR` expressions with filters:
- `=`, `>`, `<`, `>=`, `<=`
- `includes` (LIKE)
- `is null`, `is not null`
- `in`, `not in`

---

## 🛡️ Validation & Safety

- **Max body size**: 1MB
- **Content-Type**: Must be `application/json`
- **SQL Injection**: Prevented via parameterized queries and field name reflection
- **Error Handling**:
  - `bad_request`: Invalid/missing input
  - `not_found`: Empty results
  - `server_error`: Internal execution failure