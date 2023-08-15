# Speech Service

## Features

### Speech-To-Text
- Self-hosted Whisper API 사용
- REST
- No auth

### Text-To-Speech
- Google Cloud API 사용
- gRPC
- Service Account Key


## Todo
- Whisper Server 토큰 기반 인증
- google 서비스 계정 키: 'Valut'를 사용하여 키 순환 및 중앙 인증 관리


## Service Structure
```
.
|-- README.md
|-- google-credentials.json
|-- main.go
`-- server
    |-- google
    |   |-- google.go
    |   `-- service_info.go
    |-- server.go
    |-- speech.go
    `-- whisper
        |-- service_info.go
        `-- whisper.go
```