package repository

import (
	"github.com/jmoiron/sqlx"
	"go-ws/internal/entity"
	"log"
)

type MessageRepository interface {
	Create(message *entity.Message) error
	Fetch(request *entity.RequestFetchMessage) ([]*entity.Message, error)
}

type messageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (m *messageRepository) Create(message *entity.Message) error {

	q := `INSERT INTO messages (id, username, recipient, message) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, username, recipient, message`

	err := m.db.QueryRow(q, message.Id, message.Username, message.Recipient, message.Message).Scan(&message.Id, &message.CreatedAt, &message.UpdatedAt, &message.Username, &message.Recipient, &message.Message)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *messageRepository) Fetch(request *entity.RequestFetchMessage) ([]*entity.Message, error) {
	q := `SELECT * FROM messages WHERE username = $1 AND recipient = $2 AND created_at BETWEEN $3 AND $4`

	rows, err := m.db.Query(q, request.Username, request.Recipient, request.StartDate, request.EndDate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := make([]*entity.Message, 0)
	for rows.Next() {
		message := &entity.Message{}
		err = rows.Scan(&message.Id, &message.CreatedAt, &message.UpdatedAt, &message.Username, &message.Recipient, &message.Message)

		res = append(res, message)
	}

	return res, nil
}
