package corpus

import (
	"testing"
	"time"
)

func TestCorpus(t *testing.T) {
	words := Words{
		QQ:    3261340757,
		Group: 304279325,
		Name:  "Dog",
		Birth: time.Now(),
	}

	for _, v := range []string{
		"HasAndDead",
		"HasNotDead",
		"SuccAdopt",
		"FeedButDead",
		"FeedDead",
		"NoAdopt",
	} {
		sentence, err := words.Execute(v)
		if err != nil {
			t.Error(err)
		}
		t.Log(sentence)
	}
}
