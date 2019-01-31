package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 全ての大学イベント情報を取得
func GetAllUniversityEvent() ([]UniversityEvent, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnEvent := []UniversityEvent{}
	err = db.Order("year DESC, month DESC, day DESC").Find(&returnEvent).Error

	return returnEvent, err
}

// レコードの作成
func CreateUniversityEventFromArgment(ue *UniversityEvent) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Create(ue).Error
	return err
}

// レコードの変更
func UpdateUniversityEventFromArgment(updateUniversity *UniversityEvent) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Model(UniversityEvent{}).Where("id = ?", updateUniversity.Id).Updates(updateUniversity).Error
	return err
}

// レコードの削除
func DeleteUniversityEventFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Delete(UniversityEvent{}, "id = ?", id).Error
	return err
}

// レコードの取得
func GetUniversityEventFromId(id int) (UniversityEvent, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnUniversityEvent := UniversityEvent{Id: id}
	err = db.Find(&returnUniversityEvent).Error
	return returnUniversityEvent, err
}
