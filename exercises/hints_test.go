package exercises

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// 문제마다 "정답 명령 줄"로 볼 수 있는 조각. 힌트에 이것이 보이면 답을 준 것이다.
var answerLines = map[int][]string{
	1:  {"mkdir TIL"},
	2:  {"mkdir -p TIL/"},
	3:  {"> TIL/10-12.md"},
	4:  {">> TIL/10-12.md"},
	5:  {"mv TIL/.draft"},
	6:  {"cp TIL/10-13.md", "cp TIL/그파일"},
	7:  {"head -3"},
	8:  {"mv TIL/10_14.md"},
	9:  {"mv 10-15.md", "mv about.md"},
	10: {"rm TIL/temp-note.md"},
	11: {"rm -r trash"},
	12: {"touch mystery/"},
	13: {"pwd > ~"},
	14: {"ls TIL >"},
	15: {"grep -l 포트폴리오", "grep -rl", "cp TIL/파일이름"},
	16: {"| grep 10-"},
	17: {"cp -r TIL"},
	18: {"cp backup/10-12.md"},
	19: {"mv *.png", "mv *.md"},
	20: {`find . -name "*.md" >`},
	21: {"| wc -l >"},
	22: {"diff backup/10-13.md TIL/10-13.md >"},
}

var chapterRef = regexp.MustCompile(`교안|\d+절`)

func TestHintsDoNotGiveAnswers(t *testing.T) {
	for _, ex := range All() {
		for i, h := range ex.Hints {
			for _, ans := range answerLines[ex.ID] {
				if strings.Contains(h, ans) {
					t.Errorf("%s 힌트 %d단계에 정답 명령 %q가 들어 있습니다", ex.Slug, i+1, ans)
				}
			}
			if chapterRef.MatchString(h) {
				t.Errorf("%s 힌트 %d단계가 교안 절을 가리킵니다(교안은 바뀔 수 있음)", ex.Slug, i+1)
			}
		}
	}
}

// 채점 메시지 안의 `명령 대상` 꼴(완성된 명령 줄)을 찾는다. `cat`처럼 명령 이름만 쓴 것은 괜찮다.
var fullCommand = regexp.MustCompile("`(mkdir|rm|mv|cp|echo|touch|head|tail|grep|find|wc|diff) [^`]*`")

func TestMessagesDoNotGiveCommands(t *testing.T) {
	files, _ := filepath.Glob("e*.go")
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if m := fullCommand.FindString(string(b)); m != "" {
			t.Errorf("%s의 채점 메시지에 완성된 명령 %s가 들어 있습니다", f, m)
		}
	}
}
