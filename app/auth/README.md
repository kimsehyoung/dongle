# Auth Service


**Table of Contents**
- [Database PostgreSQL](#database-postgresql)
- [ORM ent](#orm-ent)
- [pgx](#pgx)
- [TODO](#todo)
---
<br><br>


## Database PostgreSQL

### Schema

### Role
- 추가, 변경 가능성이 있기에, enum type이 아닌 테이블 사용<br>


#### Account
- email
    - A length limit of email address (https://datatracker.ietf.org/doc/html/rfc3696#section-3)





## ORM ent


### Overview
Apache-2.0 license

- 공식 문서가 잘 되어 있음.
    - https://entgo.io/docs/getting-started  
- Facebook에서 내부적으로 사용하다가 오픈소스화 되었으며, 프로젝트에서 사용 중.
    - https://github.com/ent/ent#about-the-project
    - https://developers.facebook.com/blog/post/2021/04/26/eli5-ent-schema-as-code-go/?locale=ko_KR
- 커뮤니티, 지속적인 Release.
    - https://github.com/ent/ent
- type 기반 schema 설계 후, 코드 생성


### Usage
```bash
go install entgo.io/ent/cmd/ent@v0.12

ent new <Schema>
# Afet writing Schema
go generate ./ent

```


### Syntax

**`Fields`**<br>
https://entgo.io/docs/schema-fields

- Primary Key
    - `id field` is builtin in the schema and does not need declaration (https://entgo.io/docs/schema-fields#id-field)
    - `field.ID` annotation is used for a composite primary-key of the two edge-fields.<br>
    Therefore, the ID field will not be generated. (https://entgo.io/docs/schema-edges/#user-likes-example)



## pgx

- https://entgo.io/docs/sql-integration/#use-pgx-with-postgresql
- https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool

## TODO

### Database
- Account 테이블
    - 사용자별 password 'salt' 고려
    - 사용자 삭제 후, 유예 기간으로 'is_deleted' 추가 고려
