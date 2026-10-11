# Windows 확인 목록 (강사용)

Windows Terminal + Git Bash 프로필 기준. 위에서부터 차례로 확인합니다. 안 되는 항목은 결과를 적어 둡니다.

## 반드시 (수업 진행 여부가 걸림)

- [ ] Windows Terminal의 새 탭 목록(`∨`)에 **Git Bash** 프로필이 있다. 없으면 어떻게 추가했는지 적는다
- [ ] Git Bash 탭에서 `./shellings.exe`가 실행된다(실행 권한·Defender·SmartScreen 반응 기록)
- [ ] `./shellings.exe`(watch) 화면이 그려지고 `j` `k` `↑` `↓` `s` `c` `h` `r` `q` 키가 모두 먹는다
- [ ] 화면 안 한글 제목·설명이 깨지지 않고, 왼쪽 목록과 오른쪽 칸의 테두리가 어긋나지 않는다
- [ ] 시작할 때 화면에 이상한 문자(`]11;rgb:...` 같은 것)가 찍히거나 몇 초 멈추지 않는다(배경색 감지)
- [ ] `s` 뒤 다른 탭에서 `cd ~/shellings-work/01-make-til`로 **같은 폴더**에 들어가진다(`ls`로 확인). 이 PC의 Git Bash에서 `echo $HOME`이 무엇인지 적는다
- [ ] 1번을 풀고 `c` → ✅, `q`로 나갔다 다시 켜도 ✅가 남아 있다
- [ ] 다른 탭이 `~/shellings-work/03-add-more` **안에 있는 채로** TUI에서 `r` → `y`: 오류 없이 되돌려지고, 그 탭에서 `ls`·`cat TIL/10-12.md`가 바로 된다
- [ ] 한글 입력 상태로 TUI에서 `c`를 누르면(ㅊ으로 들어감) 채점이 되거나 "한/영 키" 안내가 나온다

## 문제별

- [ ] 10번: Git Bash의 그냥 `ls mystery`가 폴더 뒤에 `/`를 붙여 보여 주는가? (그렇다면 `ls -l` 없이도 풀리므로 문제 설명을 고칠지 판단)
- [ ] 11번: `~/shellings-work/11-absolute-path/portfolio/README.md`로 쓴 파일이 통과한다
- [ ] 12번: 통과 메시지(Git 안내) 한글이 잘 보인다
- [ ] 13번: `grep -l 포트폴리오 TIL/*.md`가 한글 단어로 파일을 찾는다
- [ ] 15번: `cat TIL/*.md | wc -l > linecount.txt`가 통과한다

## 대체 경로

- [ ] `./shellings.exe list`, `start 1`, `check`, `hint`, `reset`이 Git Bash 탭에서 동작한다
