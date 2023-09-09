```mermaid
erDiagram
    account ||--o{ video : "API REFERENCE"
     
    video {
        int id PK
        int account_id
        varchar(32) title
        varchar(64) url UK
        date created_at
    }

    subtitle {
        int id PK
        int video_id FK
        varchar(16) language
        varchar(64) url UK
        date created_at
    }

    video ||--o{ subtitle : has

```
