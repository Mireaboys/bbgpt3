## GPT3

### Install
```bash
git clone <this repo>
cd <this repo>
go run main.go
```

### Docker
```bash
git clone <this repo>
cd <this repo>
docker build -t bbgpt3:latest .
docker run -d -p 9001:9001 --env-file .\.env bbgpt3:latest
```