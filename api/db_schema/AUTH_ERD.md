```mermaid
erDiagram
    role  {
        int id PK
        varchar(16) type
    }

    account {
        int id PK
        int role_id FK
        varchar(320) email UK
        varchar(64) hashed_password
        varchar(32) name
        varchar(32) phone_number
        date created_at
    }

account }|--|| role : has
```
