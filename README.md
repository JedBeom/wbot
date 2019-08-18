# wbot

카카오톡 오픈빌더에 맞춰 새롭게 개발된 왕운봇입니다.

개발 언어: Go

현 기능:
 - 미세먼지
 - 급식
 - 일정
 - 페이스북 페이지 게시물 가져오기 (id가 하드코딩 되어있음)
 - 피드백 보내기
 - 주사위/무작위 선택/예아니요 답변
 - 학생부에 신고하기

## 실행

```
# 이 저장소를 클론한 뒤 
go build
```

```bash
cp config.json.example config.json
vi config.json # 설정 수정
vi wbot.service # 유저 이름과 디렉터리 수정
sudo cp wbot.service /etc/systemd/system/ # 서비스 디렉터리로 복사
sudo systemctl reload
sudo systemctl start wbot.service
```
