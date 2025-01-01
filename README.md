# API Documentation

(Use JSON)

## 1.	Authentication Endpoints
   
### a. Register

 Endpoint: POST /register
 
 Request Body:
 
    {
      "name": "Ahmad Fadhil Rizqi",
      "email": "fadhilrizqi@gmail.com",
      "password": "your password"
    }
  	
   Response:
   
    201 Created
    {
        "ID": 16,
        "CreatedAt": "2024-12-27T15:38:38.926+07:00",
        "UpdatedAt": "2024-12-27T15:38:38.926+07:00",
        "DeletedAt": null,
        "name": "Ahmad Fadhil Rizqi",
        "email":  fadhilrizqi@gmail.com",
        "password": “(Enscrypted)”
    }

### b. Login

  Endpoint: POST /login
    
  Request Body:
    
  Response
  
    200 OK
    {
        "token": "(JWT TOKEN)"
    }

### c. Token Validation (Optional)
 
  Endpoint: GET /validate
        
  Headers:
    
        Authorization: (Bearer your_jwt_token)
    
  Response:
  
    200 OK
    {
        "message": "Token is valid",
        "user_id": 16,
        "valid": true
    }


## 2.	User Endpoints
### a.	Create User

   Endpoint: POST /users	
   
   Request Body:
  	
      {
        "name": "Ahmad Fadhil Rizqi",
        "email": "fadhilrizqi@gmail.com",
        "password": "your password"
      }

  Response:
  
    201 Created
    {
      "id": 1,
      "name": "Ahmad Fadhil Rizqi",
      "email": "fadhilrizqi@gmail.com",
      "created_at": "2024-12-26T22:11:11.261+07:00”
    }

### b. Read User

  Endpoint: GET /users/:id
  
  Response:
  
    200 OK
    	{ 
       "ID": 9, 
       "CreatedAt": "2024-12-26T22:11:11.261+07:00", 
       "UpdatedAt": "2024-12-26T22:11:11.261+07:00", •	   
       "DeletedAt": null, 	      
       "name": "Ahmad Fadhil Rizqi", 	   
       "email": "fadhilrizqi2006@gmail.com",     
       "password": "hashed_password" 
    }

### c.	Update User

  Endpoint: PUT /users/:id
    
  Request Body:
    
    {
      "name": "Ahmad Fadhil Rizqi Update",
      "email": "fadhil_update@gmail.com",
      "password": "newpassword"
    }
    
  Response:
  
    200 OK
    {
        "ID": 9,
        "CreatedAt": "2024-12-26T22:11:11.261+07:00",
        "UpdatedAt": "2024-12-27T14:45:29.529+07:00",
        "DeletedAt": null,
        "name": "Ahmad Fadhil Rizqi Update",
        "email": "fadhil_update@gmail.com",
        "password": "newpassword"
    }

### d.	Delete User

   Endpoint: DELETE /users/:id
    
   Response:
    
     204 No Content


## 3.	Category Endpoints
   
   a.	Create category
  	
  Endpoint: POST /categories
  	
  Request Body:
  	
    {
      "category_name": "Electronics"
    }
  	
   Response:
   
    201 Created
    {
        "ID": 1,
        "CreatedAt": "2024-12-26T22:28:33.704+07:00",
        "UpdatedAt": "2024-12-26T22:28:33.704+07:00",
        "DeletedAt": null,
        "category_name": "Electronics"
    }

### b.	Read Category
    
  Endpoint: GET /categories/:id
    
  Response:
  
    200 OK
    {
        "ID": 1,
        "CreatedAt": "2024-12-26T22:28:33.704+07:00",
        "UpdatedAt": "2024-12-26T22:28:33.704+07:00",
        "DeletedAt": null,
        "category_name": "Electronics"
    }

### c.	Update Category
    
  Endpoint: PUT /categories/:id
    
  Request Body:
  
    {
      "category_name": "Updated Electronics"
    }
    
  Response:
    200 OK
    {
        "ID": 1,
        "CreatedAt": "2024-12-26T22:28:33.704+07:00",
        "UpdatedAt": "2024-12-27T15:04:58.117+07:00",
        "DeletedAt": null,
        "category_name": "Updated Electronics"
    }

### d.	Delete Category

   Endpoint: DELETE /categories/:id
   
  Response:
  
    204 No Content


## 4.	 Item Endpoints
   
### a.	Create Item:
    
  Endpoint: POST /items
   
  Request Body
   
    {
      "user_id": 1,
      "name": "Laptop",
      "description": "Laptop second hand",
      "category_id": 1,
      "price": 5000000,
      "stock": 4
    }
    
  Response:
  
    201 Created
    {
        "ID": 1,
        "CreatedAt": "2024-12-26T22:33:02.155+07:00",
        "UpdatedAt": "2024-12-26T22:33:02.155+07:00",
        "DeletedAt": null,
        "user_id": 1,
        "name": " Laptop ",
        "description": " Laptop second hand ",
        "category_id": 1,
        "price": 5000000,
        "stock": 4,
    }

### b.	Read Item

  Endpoint: GET /items/:id
  
  Response:
  
    200 OK
    {
        "ID": 1,
        "CreatedAt": "2024-12-26T22:33:02.155+07:00",
        "UpdatedAt": "2024-12-26T22:33:02.155+07:00",
        "DeletedAt": null,
        "user_id": 1,
        "name": " Laptop ",
        "description": " Laptop second hand ",
        "category_id": 1,
        "price": 5000000,
        "stock": 4,
    }

### c.	Update Item

  Endpoint: PUT /items/:id
  
  Request Body:
  
    {
      "name": "Laptop Updated",
      "description": "Updated description",
      "price": 4500000,
      "stock": 5
    }
    
   Response:
   
    200 OK
    {
        "ID": 2,
        "CreatedAt": "2024-12-26T22:33:02.155+07:00",
        "UpdatedAt": "2024-12-27T15:20:17.321+07:00",
        "DeletedAt": null,
        "user_id": 9,
        "name": "Laptop Updated",
        "description": "Updated description",
        "category_id": 1,
        "price": 4500000,
        "stock": 5,
    }

### d.	Delete Item
  Endpoint: DELETE /items/:id
  Response:
  
    204 No Content


## 5.	Transaction Endpoints
   
### a.	Create Transaction:
  	
   Endpoint: POST /transactions
  	
   Request Body:
  	
    {
      "item_id": 1,
      "buyer_id": 1,
      "status": "completed",
      "date": "2024-12-27"
    }
  	
  Response:
  
    201 Created
    {
        "ID": 1,
        "CreatedAt": "2024-12-27T15:24:14.401+07:00",
        "UpdatedAt": "2024-12-27T15:24:14.401+07:00",
        "DeletedAt": null,
        "item_id": 1,
        "buyer_id": 1,
        "status": "completed",
        "date": "2024-12-27"
    }

### b.	Read Transaction

  Endpoint: GET /transactions/:id
  
  Response:
  
    200 OK
    {
        "ID": 1,
        "CreatedAt": "2024-12-27T15:24:14.401+07:00",
        "UpdatedAt": "2024-12-27T15:24:14.401+07:00",
        "DeletedAt": null,
        "item_id": 1,
        "buyer_id": 1,
        "status": "completed",
        "date": "2024-12-27"
    }


### c.	Update Transaction

   Endpoint: PUT /transactions/:id
    
  Request Body:
    
    {
      "status": "pending"
    }
    
  Response:
  
    200 OK
    {
        "ID": 1,
        "CreatedAt": "2024-12-27T15:24:14.401+07:00",
        "UpdatedAt": "2024-12-27T15:30:33.936+07:00",
        "DeletedAt": null,
        "item_id": 1,
        "buyer_id": 1,
        "status": "pending",
        "date": "2024-12-26"
    }

### d.	Delete Transaction

   Endpoint: DELETE /transactions/:id
    
    Response:
    204 No Content



