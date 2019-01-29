package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// lecture_idから対応する試験情報を取得する
func GetTestFromLectureId(lectureId int) ([]Test, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	test := []Test{}

	err = db.Where(&Test{LectureID: lectureId}).Find(&test).Error
	return test, err
}

// lecture_idから対応する休講情報を取得する
func GetCancelFromLectureId(lectureId int) ([]Cancel, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cancel := []Cancel{}

	err = db.Where(&Cancel{LectureID: lectureId}).Find(&cancel).Error
	return cancel, err
}

// lecture_idから対応する教室変更情報を取得する
func GetChangeRoomFromLectureId(lectureId int) ([]ChangeRoom, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	change := []ChangeRoom{}

	err = db.Where(&ChangeRoom{LectureID: lectureId}).Find(&change).Error
	return change, err
}

// 試験情報を登録する
func CreateTestFromArgment(createTest *Test) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Create(createTest).Error
	return err
}

// 休講情報を作成
func CreateCancelFromArgment(createCancel *Cancel) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Create(createCancel).Error
	return err
}

// 教室変更情報を登録
func CreateChangeRoomFromArgment(createRoom *ChangeRoom) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Create(createRoom).Error
	return err
}

// idから休講情報を取得する
func GetCancelFromId(id int) (Cancel, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnCancel := Cancel{Id: id}

	err = db.First(&returnCancel).Error

	return returnCancel, err
}

// idから試験情報を取得する
func GetTestFromId(id int) (Test, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnTest := Test{Id: id}

	err = db.First(&returnTest).Error

	return returnTest, err
}

// idから教室変更情報を取得する
func GetChangeRoomFromId(id int) (ChangeRoom, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	returnChangeRoom := ChangeRoom{Id: id}

	err = db.First(&returnChangeRoom).Error

	return returnChangeRoom, err
}

// idから休講情報を削除する
func DeleteCancelFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Delete(&Cancel{Id: id}).Error
	return err
}

// idから試験情報を削除する
func DeleteTestFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Delete(&Test{Id: id}).Error
	return err
}

// idから教室変更を削除する
func DeleteChangeRoomFromId(id int) error {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Delete(&ChangeRoom{Id: id}).Error
	return err
}

// 全ての休講情報を取得
func GetAllCancel() ([]Cancel, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cancel := []Cancel{}
	err = db.Find(&cancel).Error
	return cancel, err
}
// 全ての試験情報を取得
func GetAllTest() ([]Test, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	test := []Test{}
	err = db.Find(&test).Error
	return test, err
}
// 全ての教室変更情報を取得
func GetAllChangeRoom() ([]ChangeRoom, error) {
	db, err := open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	changeRoom := []ChangeRoom{}
	err = db.Find(&changeRoom).Error
	return changeRoom, err
}
