package exercises

import "shellings-ko/internal/exercise"

func init() { register(6, "06-rename", "06 이름 오타", false, setupRename, checkRename) }

func setupRename(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10_14.md": til1014,
	})
}

func checkRename(dir string) exercise.Result {
	body, ok := exercise.Read(dir, "TIL/10-14.md")
	if !ok {
		return exercise.Missing(dir, "TIL/10-14.md")
	}
	if exercise.Exists(dir, "TIL/10_14.md") {
		return exercise.Fail("오타 이름 `TIL/10_14.md`가 아직 남아 있습니다. 복사가 아니라 이름 바꾸기(`mv`)를 쓰세요")
	}
	if !exercise.Same(body, til1014) {
		return exercise.Fail("`TIL/10-14.md`의 내용이 원래 파일과 다릅니다. 처음 상태로 되돌려 이름만 바꿔 보세요")
	}
	return exercise.Pass("이름의 오타를 고쳤습니다")
}
