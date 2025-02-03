🚀 Project Description:
- This is a simple REST API application that implements interaction with the PostgreSQL database. I used JWT tokens to access the features for authorized users. The application also hashes passwords before entering the database.



🛠️ Technologies Used:
- Golang 
- PostgreSQL
- Gin web framework
- JWT tokens
- TOML for config

How to use:
- /register (username, email, password) to create a new user in the database
- /authorization (email, password) to receive a jwt token
- /delete (jwt) to delete a user in the database
- /update (jwt, username) to set a new username

Example (Used Postman):
- ![image](https://github.com/user-attachments/assets/b7e7f5ff-38d4-4f5f-8d69-90173672843a)
- ![image](https://github.com/user-attachments/assets/d8d379fd-617d-4835-9ec9-48bd876c3c2c)
- ![image](https://github.com/user-attachments/assets/fbfa5e3c-bc0c-4185-831d-c0d8733834a8)



