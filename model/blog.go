package model

type Blog struct {
	ID         uint   `gorm:"primaryKey" json:"id"` // ,omitempty
	Author     string `gorm:"type:varchar(80); default:''" json:"author"`
	Tittle     string `gorm:"type:varchar(250); default:''; not null" json:"tittle"`
	Pictures   string `gorm:"default:''" json:"pictures"`
	Synthesis  string `gorm:"type:varchar(350);default:''" json:"synthesis"`
	Content    string `gorm:"not null" json:"content"`
	time       string `gorm:"type:varchar(30)" json:"time"`
	EmployeeID uint   `json:"employee_id"`
	CategoryID uint8  `json:"category_id"`
	Comments   []Comment
	TimeModel
}

type Blogs []*Blog