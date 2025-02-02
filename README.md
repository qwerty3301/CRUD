🚀 Project Description:
- This is a simple REST API application that implements interaction with the PostgreSQL database. I used JWT tokens to access the features for authorized users. The application also hashes passwords before entering the database.



🛠️ Technologies Used:
- Golang 
- PostgreSQL
- Gin web framework
- JWT tokens
- TOML for config

  This is a learning project and to run it, you need a postgreSQL database with a Users table (id, username, email, password). The data for connecting to the database can be set in the config file.

How to use:
- /register (username, email, password) to create a new user in the database
- /authorization (email, password) to receive a jwt token
- /delete (jwt) to delete a user in the database
- /update (jwt, username) to set a new username
