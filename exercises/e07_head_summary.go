package exercises

import (
	"strings"

	"shellings/internal/exercise"
)

func init() {
	register(7, "07-head-summary", "07 앞부분만 요약", false, setupHeadSummary, checkHeadSummary)
}

func setupHeadSummary(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": long1012,
		"portfolio/":   "",
	})
}

func linesOf(s string, from, to int) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if from < 0 {
		from = len(lines) + from
	}
	if to > len(lines) {
		to = len(lines)
	}
	return strings.Join(lines[from:to], "\n")
}

func checkHeadSummary(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "portfolio/summary.md")
	if !ok {
		return exercise.Missing(dir, "portfolio/summary.md")
	}
	switch {
	case exercise.Same(body, linesOf(long1012, 0, 3)):
		if orig, _ := exercise.Read(dir, "TIL/10-12.md"); !exercise.Same(orig, long1012) {
			return exercise.Fail("요약은 맞지만 원본 `TIL/10-12.md`가 바뀌었습니다. 원본은 그대로 두어야 합니다")
		}
		return exercise.Pass("앞 3줄만 골라 요약 파일을 만들었습니다")
	case exercise.Same(body, long1012):
		return exercise.Fail("`summary.md`에 파일 전체가 들어 있습니다. 앞 3줄만 있어야 합니다")
	case exercise.Same(body, linesOf(long1012, -3, 1<<30)):
		return exercise.Fail("`summary.md`에 뒷부분 3줄이 들어 있습니다. 앞부분 3줄이어야 합니다")
	}
	return exercise.Fail("`summary.md`의 내용이 원본의 앞 3줄과 다릅니다")
}
