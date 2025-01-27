# Zombie Process Demo

이 프로젝트는 Go로 작성된 좀비 프로세스 데모 애플리케이션입니다. Docker 컨테이너 환경에서 좀비 프로세스가 어떻게 생성되고 관리되는지 보여주는 예제입니다.

## 기능

- 좀비 프로세스 생성 API
- 현재 실행 중인 프로세스 목록 조회 API
- Docker 환경에서의 좀비 프로세스 관리 데모

## 기술 스택

- Go 1.23.5
- Docker
- Tini (컨테이너 init 시스템)

## API 엔드포인트

- `POST /api/zombie-process`: 새로운 좀비 프로세스 생성
- `GET /api/processes`: 현재 실행 중인 모든 프로세스 목록 조회

## 실행 방법

### 일반 실행

```
go build -o main ./cmd/main.go
./main
```

### Docker로 실행 (좀비 프로세스 발생 가능)

```
docker build -f Dockerfile -t zombie-process-demo .
docker run -p 8080:8080 zombie-process-demo
```

### Docker + Tini로 실행 (좀비 프로세스 방지)

```
docker build -f Dockerfile_tini -t zombie-process-demo-tini .
docker run -p 8080:8080 zombie-process-demo-tini
```

## 좀비 프로세스 테스트

1. 서버 실행
2. `curl -X POST http://localhost:8080/api/zombie-process`로 좀비 프로세스 생성
3. `curl http://localhost:8080/api/processes`로 프로세스 목록 확인

## Docker 환경에서의 좀비 프로세스 관리

이 프로젝트는 두 가지 Docker 설정을 제공합니다:

1. 기본 Docker 설정 (\`Dockerfile\`): 좀비 프로세스가 발생할 수 있는 환경
2. Tini 기반 Docker 설정 (\`Dockerfile_tini\`): init 시스템을 통해 좀비 프로세스를 자동으로 정리

## 라이선스

MIT
