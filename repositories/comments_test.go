package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yziori/go-intermediate/models"
	"github.com/yziori/go-intermediate/repositories"
)

// SelectCommentList関数のテスト
func TestSelectCommentList(t *testing.T) {
	const articleID = 1

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Fatalf("want comment of articleID %d but got %d\n", articleID, comment.ArticleID)
		}
	}

}

// InsertComment関数のテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test comment",
	}
	expectedCommentID := 3

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id id expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	// テストデータを削除
	t.Cleanup(func() {
		const deleteSQL = `
			delete from comments
			where message = ?;
		`

		testDB.Exec(deleteSQL, newComment.Message)

		const resetSQL = `ALTER TABLE comments AUTO_INCREMENT = 3;`
		testDB.Exec(resetSQL)
	})
}
