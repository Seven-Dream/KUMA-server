package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 全ての学生イベント情報を取得
func GetAllStudentEvent() ([]StudentEvent, error){
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	event := []StudentEvent{}
	err = db.Find(&event).Error

	return event, err
}

// 学生イベントの作成
func CreateStudentEventFromArgment(se *StudentEvent) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Create(se).Error
	return err
}

func DeleteStudentEventFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Delete(&StudentEvent{}, "id = ?", id).Error
	return err
}

func UpdateStudentEventFromArgment(se *StudentEvent) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Debug().Model(&StudentEvent{}).Where("id = ?", se.Id).Updates(se).Error
	return err
}

func GetStudentEventFromId(id int) (StudentEvent, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	se := StudentEvent{Id: id}
	err = db.First(&se).Error
	return se, err
}
