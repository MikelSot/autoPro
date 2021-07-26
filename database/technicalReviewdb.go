package database

import "github.com/MikelSot/autoPro/model"


type TechnicalReviewDao struct {
	technicalReviewDao model.TechnicalReview
}

func NewTechnicalReviewDao() TechnicalReviewDao {
	return TechnicalReviewDao{}
}

func (t *TechnicalReviewDao) Create(review *model.TechnicalReview) error {
	DB().Create(&review)
	return nil
}

func (t *TechnicalReviewDao) Update(ID uint, review *model.TechnicalReview) error {
	reviewID := model.TechnicalReview{}
	reviewID.ID = ID
	DB().Model(&reviewID).Updates(review)
	return nil
}

func (t *TechnicalReviewDao) GetByID(ID uint) (*model.TechnicalReview, error) {
	review :=model.TechnicalReview{}
	DB().First(&review, ID)
	return &review, nil
}

func (t *TechnicalReviewDao) GetAll(max int) (*model.TechnicalReviews, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	reviews := model.TechnicalReviews{}
	DB().Limit(max).Find(&reviews)
	return &reviews, nil
}

func (t *TechnicalReviewDao) DeleteSoft(ID uint) error {
	review := model.TechnicalReview{}
	review.ID = ID
	DB().Delete(&review)
	return nil
}

func (t *TechnicalReviewDao) DeletePermanent(ID uint) error {
	review := model.TechnicalReview{}
	review.ID = ID
	DB().Unscoped().Delete(&review)
	return nil
}

func (t *TechnicalReviewDao) AllReviewClient(ID uint, max int) (*model.TechnicalReviews, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	reviews := model.TechnicalReviews{}
	DB().Limit(max).Find(&reviews, "client_id = ?", ID).Order("id desc")
	return &reviews, nil
}
