package contactinformation

import "gorm.io/gorm"

type Repository interface {
	Create(contactInformation *ContactInformation) (*ContactInformation, error)
	FindWithPagination(limit, offset int) ([]ContactInformation, int64, error)
	FindByID(id uint) (*ContactInformation, error)
	UpdateFields(id uint, fields map[string]interface{}) (*ContactInformation, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(contactInformation *ContactInformation) (*ContactInformation, error) {
	if err := r.db.Create(contactInformation).Error; err != nil {
		return nil, err
	}
	return contactInformation, nil
}

func (r *repo) UpdateFields(id uint, fields map[string]interface{}) (*ContactInformation, error) {
	var contactInformation ContactInformation
	if err := r.db.Model(&contactInformation).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	if err := r.db.First(&contactInformation, id).Error; err != nil {
		return nil, err
	}
	return &contactInformation, nil
}

func (r *repo) FindWithPagination(limit, offset int) ([]ContactInformation, int64, error) {
	var contactInformations []ContactInformation
	var total int64
	if err := r.db.Model(&ContactInformation{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Limit(limit).Offset(offset).Find(&contactInformations).Error; err != nil {
		return nil, 0, err
	}
	return contactInformations, total, nil
}

func (r *repo) FindByID(id uint) (*ContactInformation, error) {
	var contactInformation ContactInformation
	if err := r.db.First(&contactInformation, id).Error; err != nil {
		return nil, err
	}
	return &contactInformation, nil
}
func (r *repo) Delete(id uint) error {
	if err := r.db.Delete(&ContactInformation{}, id).Error; err != nil {
		return err
	}
	return nil
}
