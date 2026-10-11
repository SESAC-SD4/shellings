# shellings — 터미널 연습 문제

터미널에서 **진짜 명령**을 쳐서 15문제를 풉니다. 결과로 남은 파일과 폴더를 프로그램이 자동으로 채점합니다.

## 준비: 터미널 두 개

| | 터미널 1 — 문제 화면 | 터미널 2 — 명령 치는 곳 |
| --- | --- | --- |
| Windows | Windows Terminal에서 **Git Bash** 탭 | 같은 창에 **Git Bash** 탭 하나 더 (`Ctrl+Shift+T` 또는 탭 옆 `∨`에서 Git Bash) |
| mac | iTerm 또는 터미널 창 | 새 탭(`Cmd+T`) |

> Windows에서는 PowerShell 탭이 아니라 **Git Bash** 탭이어야 합니다. 프롬프트에 `MINGW64`가 보이면 맞습니다.

## 시작하기

터미널 1에서:

```bash
cd ~
git clone <강사가 알려 준 주소> shellings
cd shellings
./shellings.exe     # Windows
./shellings         # mac
```

왼쪽에 문제 목록, 오른쪽에 문제 설명이 보입니다.

| 키 | 하는 일 |
| --- | --- |
| `j` `k` (또는 `↓` `↑`) | 문제 고르기 |
| `s` | 연습 폴더 만들기 → 오른쪽 아래 `cd ...`를 터미널 2에서 칩니다 |
| `c` | 채점 |
| `h` | 힌트 (누를 때마다 한 단계씩, 3단계까지) |
| `r` | 처음 상태로 되돌리기 |
| `q` | 끝내기 |

## 한 문제 푸는 흐름

1. 터미널 1: 문제를 고르고 `s`
2. 터미널 2: 화면에 나온 `cd ~/shellings-work/01-make-til` 같은 명령을 친다
3. 터미널 2: 문제를 푼다 (`mkdir`, `echo`, `cp` ...)
4. 터미널 1: `c` — 통과하면 ✅, 아니면 무엇이 아직인지 알려 줍니다

연습 폴더는 모두 `~/shellings-work/` 안에 생깁니다. `~/sd4/`는 건드리지 않습니다.

## 화면이 안 뜰 때 (명령으로 풀기)

```bash
./shellings.exe list       # 목록 (mac은 ./shellings)
./shellings.exe start 1    # 1번 연습 폴더 만들기 + 문제 보기
./shellings.exe check      # 채점
./shellings.exe hint       # 힌트
./shellings.exe reset      # 처음 상태로
```

## 참고한 것

- 분할 화면 구성은 [dockerlings](https://github.com/furkan/dockerlings)(MIT)를 참고했습니다
- 명령 체계는 같은 수업의 `gitlings-ko`와 같습니다

라이선스: MIT
