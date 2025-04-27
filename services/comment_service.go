package services

import (
	"github.com/yziori/go-intermediate/apperrors"
	"github.com/yziori/go-intermediate/models"
	"github.com/yziori/go-intermediate/repositories"
)

// PostCommentHandlerで使うことを想定したサービス
// 引数情報を元に新しいコメントを作成し、結果を返却する
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
