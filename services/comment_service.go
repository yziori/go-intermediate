package services

import (
	"github.com/yziori/go-intermediate/models"
	"github.com/yziori/go-intermediate/repositories"
)

// PostCommentHandlerで使うことを想定したサービス
// 引数情報を元に新しいコメントを作成し、結果を返却する
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
