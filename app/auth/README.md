# Auth Service


**Table of Contents**
- [Important Notes](#important-notes)
- [Database PostgreSQL](#database-postgresql)
- [ORM ent](#orm-ent)
- [pgx](#pgx)
- [TODO](#todo)
---
<br><br>

## Important Notes
데이터베이스 서버를 처음 실행할 때는 `init_db.go`를 실행하여 'Schema 생성', 'role', 'root 계정' 초기화 필요


## Database PostgreSQL

### Naming
https://launchbylunch.com/posts/2014/Feb/16/sql-naming-conventions/

### Schema per Service

https://entgo.io/docs/feature-flags/#schema-config

### Tables

#### Role
- 추가, 변경 가능성이 있기에, enum type이 아닌 테이블 사용<br>


#### Account
- email
    - A length limit of email address (https://datatracker.ietf.org/doc/html/rfc3696#section-3)
- validator
    - 계정 생성 유효성 검사는 API서버에서 "복잡도", "형식" 같른 로직을 처리하고, DB에서는 무결성 검증




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

ent new <Schema> Account Role
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

### lib/pq
This package is currently in maintenance mode, which means:
- It generally does not accept new features.
- It does accept bug fixes and version compatability changes provided by the community.
- Maintainers usually do not resolve reported issues.
- Community members are encouraged to help each other with reported issues.
- For users that require new features or reliable resolution of reported bugs,
    we recommend using pgx which is under active development.

### pgx
- https://entgo.io/docs/sql-integration/#use-pgx-with-postgresql
- https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool

## TODO

### Database
- Valut에서 DB 정보, Secret key 정보
- Account 테이블
    - 사용자별 password 'salt' 고려
    - 사용자 삭제 후, 유예 기간으로 'is_deleted' 추가 고려


### Security
- uuid
    - https://tomharrisonjr.com/uuid-or-guid-as-primary-keys-be-careful-7b2aa3dcb439
    - https://www.percona.com/blog/store-uuid-optimized-way/?source=post_page-----43568d94878a--------------------------------