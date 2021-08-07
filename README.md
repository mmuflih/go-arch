# Go Clean Architecture With Dependency Injection (DI)

## Tools
- Glide https://github.com/Masterminds/glide
- sql-migrate https://github.com/rubenv/sql-migrate
- fresh https://github.com/gravityblast/fresh

## How to Use
    git clone https://github.com/mmuflih/go-arch.git
    cd go-arch
    glide install
    cp env.json.example env.json
    cp dbconfig.yml.example dbconfig.yml
    sql-migrate up
    fresh
    
## Project Structure
    - Main_Project
        - config
        - container
        - context
        - domain
        - httphandler
        - lib
        - migrations
        - role
        
## Endpoints
- Register
        /api/v1/user/register
        
        [request body]

        {
            "email": "mmuflic@gmail.com",
            "pin": "123456"
        }

- Login
        /api/v1/user/login
        
        [request body]

        {
            "email": "mmuflic@gmail.com",
            "pin": "123456"
        }

- Get Me
        /api/v1/user