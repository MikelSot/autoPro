package model

type TechnicalReview struct {
	ID                uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	CarFeatures       string `gorm:"type:varchar(350); default:''" json:"car_features"`
	CommonFaults      string `gorm:"type:varchar(250); default:''" json:"common_faults"`
	TechnicalAnalysis string `gorm:"type:varchar(350); default:''" json:"technical_analysis"`
	VehicleType       string `gorm:"type:varchar(30); default:''" json:"vehicle_type"`
	RevisionType      bool   `gorm:"default:false" json:"revision_type"`
	Requirements      string `gorm:"type:varchar(350); default:''" json:"requirements"`
	ServicesVehicle   string `gorm:"type:varchar(50); default:''" json:"services_vehicle"`
	PickUp            bool   `gorm:"default:false; not null" json:"pick_up"`
	Arrangements      string `gorm:"type:varchar(350); default:''" json:"arrangements"`
	VehicleCondition  bool   `gorm:"default:false; not null" json:"vehicle_condition"`
	ClientID          uint   `gorm:"not null" json:"client_id"`
	TimeModel
}

type TechnicalReviews []*TechnicalReview