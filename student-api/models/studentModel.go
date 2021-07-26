package models

type Student struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	Lastname string `json:"lastname"`
	DOB string `json:"dob"`
	Address string `json:"address"`
	Subjects []Subject `gorm:"many2many:student_marks;"`
}

type Subject struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type StudentMarks struct {
	StudentID int  `json:"student_id";gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubjectID int `json:"subject_id";gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Marks int `json:"marks"`
}

func (u *Student) TableName() string {
	return "student"
}
func (u *Subject) TableName() string {
	return "subject"
}
func (u *StudentMarks) TableName() string {
	return "student_marks"
}