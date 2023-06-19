# backend-example
a go backend for an imaginary social media   

*I was planning on building a frontend, but turns out `neon db` can only handle like 10 requets per minute which is not even nearly enough I dont have money to pay for a real db*

## Env
The env files needs to contain 2 variables  
  - DATABASE_URL
  - JWT_SECRET_KEY
 
## Database
By default, this project uses the postgresql api for GORM, but it can be changed with very minor changes
