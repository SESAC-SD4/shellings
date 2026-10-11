package exercises

import "shellings-ko/internal/exercise"

func init() { register(5, "05-read-and-copy", "05 목록대로 복사", false, setupReadAndCopy, checkReadAndCopy) }

const pickFile = "10-13.md"

func setupReadAndCopy(dir string) error {
	return exercise.WriteFiles(dir, map[string]string{
		"TIL/10-12.md": til1012,
		"TIL/10-13.md": til1013,
		"TIL/10-14.md": til1014,
		"portfolio/":   "",
		"index.txt":    "포트폴리오에 올릴 글: " + pickFile + "\n",
	})
}

func checkReadAndCopy(dir string) exercise.Result {
	for _, name := range exercise.List(dir, "portfolio") {
		if name != pickFile {
			return exercise.Fail("`portfolio`에 다른 글이 들어 있습니다: `" + name + "`. `index.txt`를 열어 어느 글인지 확인하세요")
		}
	}
	body, ok := exercise.Read(dir, "portfolio/"+pickFile)
	if !ok {
		return exercise.Fail("`portfolio`에 아직 올릴 글이 없습니다. 어느 글인지는 `index.txt`에 적혀 있습니다")
	}
	if !exercise.IsFile(dir, "TIL/"+pickFile) {
		return exercise.Fail("원본 `TIL/" + pickFile + "`이 사라졌습니다. 옮기기(`mv`)가 아니라 복사(`cp`)해야 합니다")
	}
	if !exercise.Same(body, til1013) {
		return exercise.Fail("`portfolio/" + pickFile + "`의 내용이 원본과 다릅니다. `cp`로 복사해 보세요")
	}
	return exercise.Pass("`index.txt`대로 글을 포트폴리오에 복사했습니다")
}
