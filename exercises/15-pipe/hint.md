`cat`은 파일 여러 개를 이어서 보여 줄 수 있습니다: `cat TIL/*.md`
---
`cat TIL/*.md | wc -l`이 합친 줄 수를 숫자 하나로 보여 줍니다. 이 출력을 `>`로 파일에 보냅니다.
---
`cat TIL/*.md | wc -l > linecount.txt`를 치고 `cat linecount.txt`로 확인하세요.
