# Windows 확인 목록 (강사용)

Windows Terminal + Git Bash 프로필 기준. 위에서부터 차례로 확인합니다. 안 되는 항목은 결과를 적어 둡니다.

## 반드시 (수업 진행 여부가 걸림)

- [ ] Windows Terminal의 새 탭 목록(`∨`)에 **Git Bash** 프로필이 있다. 없으면 어떻게 추가했는지 적는다
- [ ] Git Bash 탭에서 `./shellings.exe`가 실행된다(실행 권한·Defender·SmartScreen 반응 기록)
- [ ] 켜자마자 `~/shellings-work/` 안에 `01-make-til`부터 `22-diff`까지 22개 폴더가 생긴다
- [ ] 화면이 그려지고 `↑` `↓` Enter Tab Ctrl+R Esc가 모두 먹는다
- [ ] **한글 입력 상태**에서도 Enter(채점)·Tab(힌트)·Ctrl+R·Esc·화살표가 먹는다. 한글 상태에서 `c`를 누르면 무엇이 일어나는지 적는다
- [ ] 화면 안 한글 제목·설명이 깨지지 않고, 왼쪽 목록과 오른쪽 칸의 테두리가 어긋나지 않는다
- [ ] 시작할 때 화면에 이상한 문자(`]11;rgb:...` 같은 것)가 찍히거나 몇 초 멈추지 않는다(배경색 감지)
- [ ] 다른 탭에서 `cd ~/shellings-work/01-make-til`로 **같은 폴더**에 들어가진다(`ls`로 확인). 이 PC의 Git Bash에서 `echo $HOME`이 무엇인지 적는다
- [ ] 1번을 풀고 Enter → ✅, Esc로 나갔다 다시 켜도 ✅와 풀던 내용이 남아 있다
- [ ] 다른 탭이 `~/shellings-work/04-add-more` **안에 있는 채로** TUI에서 Ctrl+R → Enter: 오류 없이 되돌려지고, 그 탭에서 `ls`·`cat TIL/10-12.md`가 바로 된다

## 문제별

- [ ] 12번: Git Bash의 그냥 `ls mystery`가 폴더 뒤에 `/`를 붙여 보여 주는가? (그렇다면 `ls -l` 없이도 풀리므로 문제 설명을 고칠지 판단)
- [ ] 13번: `week1` 안에서 `pwd`를 `~/shellings-work/13-absolute-path/portfolio/where.txt`로 보내면 통과한다(Git Bash의 `pwd`는 `/c/Users/...` 꼴)
- [ ] 15번: `grep -rl 포트폴리오 TIL`이 한글 단어로 파일을 찾는다
- [ ] 16번: `ls TIL | grep 10-` 결과를 `october.txt`로 보내면 통과한다
- [ ] 18번: 통과 메시지(Git 안내) 한글이 잘 보인다
- [ ] 19★번: `mv *.png images/`가 동작한다
- [ ] 20★번: Git Bash에서 `find`가 Windows의 `find.exe`가 아니라 리눅스 `find`로 실행된다(`which find` → `/usr/bin/find`)
- [ ] 21★번: `wc -l`의 합계 줄이 `total`로 나온다(한국어로 번역되어 나오면 적는다)
- [ ] 22★번: Git Bash에 `diff`가 있다(`which diff`), 한글 줄이 깨지지 않는다

## 대체 경로

- [ ] `./shellings.exe list`, `start 1`, `check`, `hint`, `reset`이 Git Bash 탭에서 동작한다
