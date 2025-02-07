package repositories_test

import (
	"testing"

	"github.com/yziori/go-intermediate/models"
	"github.com/yziori/go-intermediate/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	//step1 テスト名とテストケースのセットのスライスを作成
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NickNum:  2,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NickNum:  4,
			},
		},
	}

	// step2 step1で作成したものをfor文で回す
	for _, test := range tests {
		// step3 for文の中でサブテスト実施
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: got %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: got %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: got %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NickNum != test.expected.NickNum {
				t.Errorf("NickName: got %d but want %d\n", got.NickNum, test.expected.NickNum)
			}
		})
	}

}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	// テスト対象の関数を実行
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// SelectArticleList関数から得たArticleスライスの長さが期待通りでないならFAILにする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d\n", expectedNum, num)
	}
}
