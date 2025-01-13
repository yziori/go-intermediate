package repositories

import (
	"database/sql"

	"github.com/yziori/go-intermediate/models"
)

// 新規コメントをDBにinsertする関数
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now())
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	row, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	commentArray := make([]models.Comment, 0)
	for row.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		if err := row.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime); err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
