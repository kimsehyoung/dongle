



## Todo
- JWT secret: 추후 Env -> hashicorp 'Valut'(MPL 2.0) or AWS 'secret manager'로 변경
- graceful shutdown, rate limit
- refresh toekn rotation으로 보안 강화 -> But, 정상 사용자가 refresh token을 사용하기 전에, 공격자가 탈취하는 경우에 대한 처리가 필요.