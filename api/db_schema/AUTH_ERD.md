```mermaid
erDiagram
    role  {
        integer id PK
        varchar(16) type
    }

    account {
        integer id PK
        integer role FK
        varchar(320) email UK
        varchar(64) hashed_password
        varchar(32) name
        varchar(32) phone_number UK
        date created_at
    }

account }|--|| role : has
```
