package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/overlineink/nellcorp-challenge-jorge-costa/domain/entities"
)

type AccountRepositoryDb struct {
	Db *gorm.DB
}

func (r AccountRepositoryDb) FindAccountById(id string) (*entities.Account, error) {
	var account entities.Account
	r.Db.Preload("Transactions").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("account not found")
	}
	return &account, nil
}

func (r *AccountRepositoryDb) Save(account *entities.Account) error {
	err := r.Db.Save(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r AccountRepositoryDb) Register(account *entities.Account) (*entities.Account, error) {
	err := r.Db.Create(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}
