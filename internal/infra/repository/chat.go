package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lassulfi/fclx/chatservice/internal/domain/entity"
	"github.com/lassulfi/fclx/chatservice/internal/infra/db"
)

type ChatRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewChatRepositoryMySQL(dtb *sql.DB) *ChatRepositoryMySQL {
	return &ChatRepositoryMySQL{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *ChatRepositoryMySQL) CreateChat(ctx context.Context, chat *entity.Chat) error {
	err := r.Queries.CreateChat(
		ctx,
		db.CreateChatParams{
			ID:               chat.ID,
			UserID:           chat.UserID,
			InitialMessageID: chat.InitialSystemMessage.Content,
			Status:           chat.Status,
			TokenUsage:       int32(chat.TokenUsage),
			Model:            chat.Config.Model.GetModelName(),
			ModelMaxTokens:   int32(chat.Config.MaxTokens),
			Temperature:      float64(chat.Config.Temperature),
			TopP:             float64(chat.Config.TopP),
			N:                int32(chat.Config.N),
			Stop:             chat.Config.Stop[0],
			MaxTokens:        int32(chat.Config.MaxTokens),
			PresencePenalty:  float64(chat.Config.PresencePenalty),
			FrequencyPenalty: float64(chat.Config.PresencePenalty),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	)

	if err != nil {
		return err
	}
	return nil
}
