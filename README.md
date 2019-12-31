# Go Clean Architecture With Dependency Injection (DI)

## Tools
- Glide https://github.com/Masterminds/glide
- sql-migrate https://github.com/rubenv/sql-migrate
- fresh https://github.com/gravityblast/fresh

## How to Use
    git clone https://github.com/mmuflih/go-di-arch.git
    cd go-di-arch
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
        
