```mermaid
erDiagram
    role  {
        integer id PK
        varchar(16) member_type
    }

    account {
        integer id PK
        integer role_id FK
        varchar(32) login_id UK
        varchar(64) hashed_password
        varchar(32) name
        varchar(320) email UK
        varchar(32) phone_number UK
        date created_at
    }

account }|--|| role : has
```
