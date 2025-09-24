package core

import (
	"encoding/binary"
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID        `json:"-" gorm:"column:id;" db:"id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"  db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"  db:"updated_at"`
}

func NewSQLModel() SQLModel {
	now := time.Now().UTC()

	return SQLModel{
		Id:        uuid.New(),
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
func (sqlModel *SQLModel) Mask(objectId int) {
	uid := NewUID(binary.BigEndian.Uint32(sqlModel.Id[:4]), objectId, 1)
	sqlModel.FakeId = &uid
}
