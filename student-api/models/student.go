package models

import (
	"freshers-bootcamp/student-api/config"
)

func GetAllStudents(student *[]Student) error {
	if err := config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

func CreateStudent(student *Student) error {
	if err := config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSubjects(subject *[]Subject) error {
	if err := config.DB.Find(subject).Error; err != nil {
		return err
	}
	return nil
}

func CreateSubject(subject *Subject) error {
	if err := config.DB.Create(subject).Error; err != nil {
		return err
	}
	return nil
}

func CreateStudentMarks(marks *StudentMarks) error {
	if err := config.DB.Create(marks).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentMarksByID(marks *[]StudentMarks, id string) error {
	if err := config.DB.Where("student_id = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}