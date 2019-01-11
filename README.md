# wbot_new

카카오톡 오픈빌더에 맞춰 새롭게 개발된 왕운봇입니다.

개발 언어: Go

현 기능:
 - 미세먼지
 - 급식
 - 일정

## 실행

`wbot.serivce`를 `/etc/systemd/system/` 등의 서비스 디렉터리로 이동하고
`WorkingDirectory`를 `wbot_new` 실행 파일이 위치한 디렉터리로 바꾸어줍니다.
`User` 역시 올바르게 교체합니다.

```
go build
```

```bash
sudo systemctl reload
sudo systemctl start wbot.service
```
