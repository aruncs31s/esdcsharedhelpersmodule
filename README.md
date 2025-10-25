# ESDC Shared Helpers Module

> A lightweight, reusable Go module that provides common helper utilities for building REST APIs with Gin framework. Think of it as your toolkit for handling HTTP requests, validations, and error management consistently across all ESDC services.

- [What is This?](#what-is-this)
- [What's Inside](#whats-inside)
    - [Request Helpers](#request-helpers)
    - [Request Validators](#request-validators)
    - [Error Helpers](#error-helpers)
    - [HTTP Utilities](#http-utilities)
    - [Pre-built Error Types](#pre-built-error-types)
- [Quick Start](#quick-start)
    - [Installation](#installation)
    - [Basic Example](#basic-example)
- [Common Patterns](#common-patterns)
    - [Pattern 1: Protected List Endpoint](#pattern-1-protected-list-endpoint)
    - [Pattern 2: Get Resource by ID](#pattern-2-get-resource-by-id)
    - [Pattern 3: Create Resource](#pattern-3-create-resource)
    - [Pattern 4: Authorization Check](#pattern-4-authorization-check)
- [Architecture](#architecture)
    - [Interfaces (Type Contracts)](#interfaces-type-contracts)
    - [File Structure](#file-structure)
- [Contributing](#contributing)
- [License](#license)

## What is This?

This module ment to avoid repeating boilerplate code across services and handlers.

**Built for:** Gin-based REST APIs  
**Go Version:** 1.25.3+  
**Key Dependencies:** [Gin-gonic/Gin](https://github.com/gin-gonic/gin), [responsehelper](https://github.com/aruncs31s/responsehelper)

## What's Inside

### 1. **Request Helpers** - Make handling HTTP requests a breeze

```go
requestHelper := helper.NewRequestHelper()
```

#### Extract & Validate Username
```go
username, failed := requestHelper.GetAndValidateUsername(c, yourHandler)
if failed {
    // Already responded to client with error
    return
}
// Use username safely
```
- Extracts username from the request context
- Automatically validates it's not empty
- Responds with a friendly error message if invalid
- Perfect for authenticated endpoints

#### Parse Pagination Parameters
```go
limit, offset := requestHelper.GetLimitAndOffset(c)
// Use these to query your database
users := db.GetUsers(limit, offset)
```
- Extracts `page` and `per-page` query parameters
- Automatically calculates offset from page number
- Defaults: 10 items per page, starting at page 1
- Example: `?page=2&per-page=20` → limit=20, offset=20

#### Extract URL Parameters
```go
userID := requestHelper.GetURLParam(c, "userId")
```
- Safely retrieve URL path parameters
- Works with Gin's `:paramName` route syntax

#### Validate & Parse IDs
```go
id, failed := requestHelper.ValidateAndParseID(yourHandler, "id", c, "Invalid user ID format")
if failed {
    return // Error already sent to client
}
// Now use id safely as uint
```
- Validates ID is a valid positive integer
- Converts string to `uint` type
- Returns your custom error message if invalid
- Automatically responds to client on failure

---

### 2. **Request Validators** - Keep your data clean

```go
validator := &requestValidator{}
```

#### Validate Usernames
```go
if err := validator.ValidateUsername(username); err != nil {
    // Returns error if username is empty
    log.Println(err) // "invalid username"
}
```
- Ensures username is not empty
- Simple but effective

#### Validate & Parse IDs
```go
id, err := validator.ValidateIDAndParse(idString)
if err != nil {
    log.Println(err) // "invalid ID"
}
```
- Checks if string is a valid positive integer
- Rejects zero and negative numbers
- Converts to `uint` if valid

---

### 3. **Error Helpers** - Consistent error messages

```go
errorHelper := helper.NewErrorHelper()
```

#### Generate Authorization Errors
```go
err := errorHelper.GetRecordDoesNotBelongErrorMessage(123, "user1")
// Error message: "the record 123 does not belong to user1"
```
- Creates consistent error messages for access control
- Helpful when a user tries to access someone else's record
- Works with any ID type

---

### 4. **HTTP Utilities** - JSON handling made easy

#### Extract & Parse JSON Request Body
```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

data, failed := helper.GetJSONDataFromRequest[CreateUserRequest](c, responseHelper)
if failed {
    return // Error already sent to client
}
// Use data safely
```
- Parses JSON from request body
- Validates required fields automatically
- Returns helpful error messages
- Handles parsing errors gracefully

---

### 5. **Pre-built Error Types** - No more magic strings

```go
utils.ErrBadRequest      // "bad request"
utils.ErrNotFound        // "not found"
utils.ErrForbidden       // "forbidden"
utils.ErrInternal        // "internal server error"
utils.ErrInvalidUsername // "invalid username"
utils.ErrInvalidID       // "invalid ID"
```

#### User-Friendly Fix Messages
```go
utils.FixInvalidUsername // "Please provide a valid username."
utils.FixInvalidID       // "Please provide a valid ID."
```
- Use these in API responses to guide users
- Consistent messaging across all endpoints

---

## Quick Start

### Installation
```bash
go get github.com/aruncs31s/esdcsharedhelpersmodule
```

### Basic Example
```go
package main

import (
    "github.com/aruncs31s/esdcsharedhelpersmodule/helper"
    "github.com/gin-gonic/gin"
)

// Your handler that implements HasBothValidatorAndResponseHelper
type MyHandler struct {
    validator      helper.RequestValidator
    responseHelper // your response helper
}

func (h *MyHandler) GetValidator() helper.RequestValidator {
    return h.validator
}

// In your route handler
func getUsers(c *gin.Context) {
    username, failed := helper.GetAndValidateUsername(c, myHandler)
    if failed {
        return
    }
    
    limit, offset := helper.GetLimitAndOffset(c)
    
    // Fetch users from database
    users := db.GetUsers(limit, offset)
    c.JSON(200, users)
}
```

---

##  Common Patterns

### Pattern 1: Protected List Endpoint
```go
func (h *Handler) ListUsers(c *gin.Context) {
    // Get and validate current user
    username, failed := h.RequestHelper.GetAndValidateUsername(c, h)
    if failed {
        return
    }
    
    // Get pagination
    limit, offset := h.RequestHelper.GetLimitAndOffset(c)
    
    // Fetch and return
    users := h.DB.GetUsers(limit, offset)
    h.ResponseHelper.Success(c, users)
}
```

### Pattern 2: Get Resource by ID
```go
func (h *Handler) GetUser(c *gin.Context) {
    // Validate and parse ID
    userID, failed := h.RequestHelper.ValidateAndParseID(h, "id", c, utils.FixInvalidID)
    if failed {
        return
    }
    
    // Fetch and return
    user, err := h.DB.GetUserByID(userID)
    if err != nil {
        h.ResponseHelper.NotFound(c, "User not found")
        return
    }
    h.ResponseHelper.Success(c, user)
}
```

### Pattern 3: Create Resource
```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

func (h *Handler) CreateUser(c *gin.Context) {
    // Parse and validate JSON
    req, failed := helper.GetJSONDataFromRequest[CreateUserRequest](c, h.ResponseHelper)
    if failed {
        return
    }
    
    // Create user
    user := h.DB.CreateUser(req.Name, req.Email)
    h.ResponseHelper.Created(c, user)
}
```

### Pattern 4: Authorization Check
```go
func (h *Handler) DeleteUserPost(c *gin.Context) {
    postID, failed := h.RequestHelper.ValidateAndParseID(h, "postId", c, utils.FixInvalidID)
    if failed {
        return
    }
    
    username, _ := c.Get("username")
    post, err := h.DB.GetPost(postID)
    
    if post.Owner != username {
        err := h.ErrorHelper.GetRecordDoesNotBelongErrorMessage(postID, username)
        h.ResponseHelper.Forbidden(c, err.Error())
        return
    }
    
    h.DB.DeletePost(postID)
    h.ResponseHelper.Success(c, nil)
}
```

---

## Architecture

### Interfaces (Type Contracts)
The module uses Go interfaces for flexibility:

- **`RequestHelper`** - All request extraction & validation methods
- **`RequestValidator`** - Data validation rules
- **`ErrorHelper`** - Error message generation
- **`HasValidator`** - Provides access to validator
- **`HasResponseHelper`** - Provides access to response helper
- **`HasBothValidatorAndResponseHelper`** - Combined interface

This design means you can easily mock these for testing!

### File Structure
```
helper/                    # Main implementation
├── request_helper.go     # HTTP request utilities
├── request_validator.go  # Input validation logic
├── error_helper.go       # Error message generation
└── http_utils.go         # JSON parsing helpers

interface/helper/         # Type definitions (contracts)
├── request_helper.go
├── request_validator.go
├── error_helper.go
└── validator.go

utils/
└── errors.go            # Pre-defined error types
```

---


## 🤝 Contributing

Found a bug? Need a feature? 
- Create an issue in the repository
- Submit a pull request with your improvements
- Follow the existing code style and patterns

---

## �📄 License

Part of the ESDC shared modules ecosystem. Built for team productivity.
