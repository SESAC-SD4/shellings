package exercises

import "shellings-ko/internal/exercise"

func init() {
	register(16, "16-grep-pipe", "16 10월만 골라내기", false, setupGrepPipe, checkGrepPipe)
}

var (
	octoberWant = []string{"10-12.md", "10-13.md", "10-14.md"}
	septemberNo = []string{"09-28.md", "09-30.md"}
)

func setupGrepPipe(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/09-28.md": til0928,
		"TIL/09-30.md": til0930,
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": til1014,
	})
}

func checkGrepPipe(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "october.txt")
	if !ok {
		return exercise.Missing(dir, "october.txt")
	}
	got := names(body)
	for _, n := range septemberNo {
		if got[n] {
			return exercise.Fail("9월 파일이 섞여 있습니다: `" + n + "`. 10월 파일만 남겨야 합니다")
		}
	}
	for _, n := range octoberWant {
		if !got[n] {
			return exercise.Fail("빠진 10월 파일이 있습니다: `" + n + "`")
		}
	}
	return exercise.Pass("파이프로 10월 기록만 골라냈습니다")
}
