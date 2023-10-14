```mermaid
erDiagram
    account ||--o{ original_video : "API REFERENCE"

    original_video {
        int id PK
        int account_id
        varchar(32) title
        varchar(64) url UK
        date created_at
    }

    subtitled_video {
        int id PK
        int original_video_id FK
        varchar(64) url UK
        date created_at
    }

    original_video ||--o{ subtitled_video : has

```
