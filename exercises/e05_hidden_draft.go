package exercises

import "shellings-ko/internal/exercise"

func init() {
	register(5, "05-hidden-draft", "05 숨은 초안", false, setupHiddenDraft, checkHiddenDraft)
}

func setupHiddenDraft(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/.draft":   draft1013,
	})
}

func checkHiddenDraft(dir string) exercise.Result {
	if exercise.Exists(dir, "TIL/.draft") {
		if exercise.IsFile(dir, "TIL/10-13.md") {
			return exercise.Fail("`TIL/10-13.md`는 생겼지만 숨김 파일 `.draft`가 아직 남아 있습니다. 새로 만들지 말고 `.draft`의 이름을 바꾸세요")
		}
		return exercise.Missing(dir, "TIL/10-13.md")
	}
	body, ok := exercise.Read(dir, "TIL/10-13.md")
	if !ok {
		return exercise.Fail("숨김 파일 `.draft`가 사라졌는데 `TIL/10-13.md`가 없습니다. 지우지 말고 이름을 바꿔야 합니다. 처음 상태로 되돌려 다시 해 보세요")
	}
	if !exercise.Same(body, draft1013) {
		return exercise.Fail("`TIL/10-13.md`의 내용이 초안과 다릅니다. 처음 상태로 되돌려 `.draft`의 이름만 바꿔 보세요")
	}
	return exercise.Pass("숨어 있던 초안을 찾아 제 이름을 붙였습니다")
}
