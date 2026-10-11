명령의 출력을 파일로 보내는 방법은 `echo`에서 배운 `>`와 같습니다.
---
`find . -name "*.md"`의 출력을 `>`로 `filelist.txt`에 보냅니다.
---
`find . -name "*.md" > filelist.txt`를 치고 `cat filelist.txt`로 확인하세요.
