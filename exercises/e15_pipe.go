package exercises

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"shellings-ko/internal/exercise"
)

func init() { register(15, "15-pipe", "15★ 줄 수 세기", true, setupPipe, checkPipe) }

var firstNumber = regexp.MustCompile(`\d+`)

func setupPipe(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": til1014,
	})
}

func intIn(s string) int {
	m := firstNumber.FindString(s)
	if m == "" {
		return -1
	}
	n, err := strconv.Atoi(m)
	if err != nil {
		return -1
	}
	return n
}

// wc -l과 같은 규칙(줄바꿈 문자 수)으로 지금 파일들의 합계를 센다.
func totalLines(dir string) int {
	total := 0
	for _, name := range exercise.List(dir, "TIL") {
		if strings.HasSuffix(name, ".md") {
			body, _ := exercise.Read(dir, "TIL/"+name)
			total += strings.Count(body, "\n")
		}
	}
	return total
}

func checkPipe(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "linecount.txt")
	if !ok {
		return exercise.Missing(dir, "linecount.txt")
	}
	lines := exercise.NonEmptyLines(body)
	if len(lines) == 0 {
		return exercise.Fail("`linecount.txt`가 비어 있습니다")
	}
	want := totalLines(dir)
	first := intIn(lines[0])
	if first == want {
		return exercise.Pass(fmt.Sprintf("지금까지 모두 %d줄을 썼습니다", want))
	}
	for _, l := range lines {
		if strings.Contains(l, "total") && intIn(l) == want {
			return exercise.Fail("합계는 맞게 구했지만 파일마다의 줄 수까지 들어 있습니다. 숫자 하나만 남도록 `cat`과 파이프(`|`)로 이어 보세요")
		}
	}
	if first < 0 {
		return exercise.Fail("`linecount.txt`에 숫자가 없습니다")
	}
	return exercise.Fail(fmt.Sprintf("숫자가 다릅니다: %d. `TIL` 안 모든 `.md` 파일의 줄 수를 합쳐야 합니다", first))
}
