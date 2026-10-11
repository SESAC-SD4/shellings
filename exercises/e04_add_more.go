package exercises

import "shellings-ko/internal/exercise"

func init() { register(4, "04-add-more", "04 내용 덧붙이기", false, setupAddMore, checkAddMore) }

func setupAddMore(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{"TIL/10-12.md": "## 터미널\n"})
}

func checkAddMore(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "TIL/10-12.md")
	if !ok {
		return exercise.Missing(dir, "TIL/10-12.md")
	}
	lines := exercise.NonEmptyLines(body)
	hasHead := false
	for _, l := range lines {
		if l == "## 터미널" {
			hasHead = true
		}
	}
	if !hasHead {
		return exercise.Fail("처음 줄 `## 터미널`이 사라졌습니다. `>` 하나는 덮어씁니다. 처음 상태로 되돌린 뒤(TUI는 `r`) `>>`로 다시 해 보세요")
	}
	if len(lines) < 2 {
		return exercise.Fail("아직 `## 터미널` 한 줄뿐입니다. 파일 끝에 내용을 덧붙여 보세요")
	}
	return exercise.Pass("처음 줄을 지키면서 내용을 덧붙였습니다")
}
