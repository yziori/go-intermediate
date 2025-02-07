package repositories_test

import (
	"testing"

	"github.com/yziori/go-intermediate/models"
	"github.com/yziori/go-intermediate/repositories"
	"github.com/yziori/go-intermediate/repositories/testdata"

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
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
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
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// SelectArticleList関数から得たArticleスライスの長さが期待通りでないならFAILにする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d\n", expectedNum, num)
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const deleteSQL = `
		DELETE FROM articles
		WHERE title = ? AND contents = ? AND username = ?
	`
		testDB.Exec(deleteSQL, article.Title, article.Contents, article.UserName)

		// Auto Incrementのリセットを追加
		const resetSQL = `ALTER TABLE articles AUTO_INCREMENT = 3`
		_, err = testDB.Exec(resetSQL)
		if err != nil {
			t.Errorf("failed to reset AUTO_INCREMENT: %v", err)
		}
	})
}

// UpdateNickNum関数のテスト
func TestUpdateNickNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get before article")
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get after article")
	}

	if after.NickNum-before.NickNum != 1 {
		t.Errorf("NickNum is not updated: before %d, after %d\n", before.NickNum, after.NickNum)
	}

	// テスト後に元に戻す
	t.Cleanup(func() {
		const resetSQL = `
		UPDATE articles
		SET nice = 2
		WHERE article_id = 1
	`
		_, err = testDB.Exec(resetSQL)
		if err != nil {
			t.Errorf("failed to reset nice: %v", err)
		}
	})
}
