package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
	"fmt"
	"strings"

	m "KUMA-server/models"
)

// 検索フォーム
func SearchLecture(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	render(c, "search_lecture.tmpl", gin.H{"title": "Lecture Search Page"})
}

// post情報から検索条件を取得し、講義情報を取得し、
// レンダリングする
func SearchResultLecture(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	resultLecture := getSearchLectureFromPosted(c)

	searchResultLecture, err := m.SearchLectureFromArgment(&resultLecture)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	fmt.Println(searchResultLecture)

	render(c, "search_result.tmpl", gin.H{"title": "Search Results Page", "lecture": searchResultLecture})

}

// 講義のオプション（休講、試験、教室変更）をデータベースから取得し、
// レンダリングする
func ShowLectureOption(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	lecture, err := m.GetLectureDataFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	test, err := m.GetTestFromLectureId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	cancel, err := m.GetCancelFromLectureId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	change, err := m.GetChangeRoomFromLectureId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	render(c, "lecture_option_list.tmpl", gin.H{"title": "Option Page", "lecture": lecture, "test": test, "cancel": cancel, "change": change})
}

// 休講情報登録フォーム
func RegisterCancelGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("lectureId"))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	lecture, err := m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	render(c, "register_option_cancel.tmpl", gin.H{"title": "Register Cancel Page", "lecture": lecture})
}

// 休講情報の登録
func RegisterCancelPost(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.PostForm("lectureid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	// Lecture情報が取得できない場合
	// リダイレクトする
	_, err = m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	cancel, ok := getCancelFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/lecture/option/cancel/register/"+idStr)
		return
	}

	err = m.CreateCancelFromArgment(&cancel)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
}

// 試験情報登録フォーム
func RegisterTestGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("lectureId"))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	lecture, err := m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	render(c, "register_option_test.tmpl", gin.H{"title": "Register Test Page", "lecture": lecture})
}

// 試験情報の登録
func RegisterTestPost(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.PostForm("lectureid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	_, err = m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	test, ok := getTestFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/lecture/option/test/register/"+idStr)
		return
	}

	err = m.CreateTestFromArgment(&test)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
}

// 教室変更情報登録フォーム
func RegisterChangeRoomGet(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("lectureId"))
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	lecture, err := m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	render(c, "register_option_change_room.tmpl", gin.H{"title": "Register Change Room Page", "lecture": lecture})
}

// 教室変更情報の登録
func RegisterChangeRoomPost(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.PostForm("lectureid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	_, err = m.GetLectureDataFromId(id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	room, ok := getChangeRoomFromPosted(c)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/lecture/option/change_room/register/"+idStr)
		return
	}

	err = m.CreateChangeRoomFromArgment(&room)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}
	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
}

// 休講情報の詳細 
func ShowCancel(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	cancel, err := m.GetCancelFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	lecture, err := m.GetLectureDataFromId(cancel.LectureID)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	render(c, "show_cancel.tmpl", gin.H{"title": "Cancel Info Page", "cancel": cancel, "name": lecture.LectureName})
}

// 試験情報の詳細
func ShowTest(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	test, err := m.GetTestFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	lecture, err := m.GetLectureDataFromId(test.LectureID)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	render(c, "show_test.tmpl", gin.H{"title": "Test Info Page", "test": test, "name": lecture.LectureName})
}

// 教室変更の詳細
func ShowChangeRoom(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	changeRoom, err := m.GetChangeRoomFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	lecture, err := m.GetLectureDataFromId(changeRoom.LectureID)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+idStr)
		return
	}

	render(c, "show_change_room.tmpl", gin.H{"title": "Change Room Info Page", "room": changeRoom, "name": lecture.LectureName})
}

// 休講情報の削除
func DeleteCancel(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err :=strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}
	cancel, err := m.GetCancelFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
	}

	lectureIdStr := cancel.LectureID
	lectureId := strconv.Itoa(lectureIdStr)

	err = m.DeleteCancelFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
		return
	}

	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
}

// 試験情報の削除
func DeleteTest(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	test, err := m.GetTestFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
	}

	lectureIdStr := test.LectureID
	lectureId := strconv.Itoa(lectureIdStr)

	err = m.DeleteTestFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
		return
	}

	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
}

// 教室変更情報の削除
func DeleteChangeRoom(c *gin.Context) {
	_, err := sessionCheck(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
		return
	}

	changeRoom, err := m.GetChangeRoomFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/search")
	}

	lectureIdStr := changeRoom.LectureID
	lectureId := strconv.Itoa(lectureIdStr)

	err = m.DeleteChangeRoomFromId(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
		return
	}

	c.Redirect(http.StatusSeeOther, "/lecture/option/list/"+lectureId)
}

func getSearchLectureFromPosted(c *gin.Context) m.Lecture {
	name       := c.PostForm("name")
	teachar    := c.PostForm("teachar")
	yearStr    := c.PostForm("year")
	quarterStr := c.PostForm("quarter")    // example: "1"
	weekStr    := c.PostForm("week")    // example: "3"
	timeStr    := c.PostForm("time")    // example: "4"

	var returnError error
	var year, quarter, week, time int
	var err error
	var o m.Other

	if yearStr != "" {
		year, err    = strconv.Atoi(yearStr)
		if err != nil { returnError = err }
	}
	if quarterStr != "" {
		quarter, err = strconv.Atoi(quarterStr)
		if err != nil { returnError = err }
	}
	if weekStr != "" {
		week, err    = strconv.Atoi(weekStr)
		if err != nil { returnError = err
		} else { o.Week = week }
	}
	if timeStr != "" {
		time, err    = strconv.Atoi(timeStr)
		if err != nil { returnError = err
		} else { o.Time = time }
	}

	fmt.Println(returnError)

	other := []m.Other{o}
	searchLecture := m.Lecture{LectureName: name, Teachar: teachar, Year: year, Quarter: quarter, Others: other}

	return searchLecture
}

func getCancelFromPosted(c *gin.Context) (m.Cancel, bool) {
	// postされた情報の取得
	lectureId           := c.PostForm("lectureid")
	dateStrBeforeFormat := c.PostForm("date")
	comment             := c.PostForm("comment")

	id, _ := strconv.Atoi(lectureId)

	// 空の入力があるか
	var postDataEmpty = false

	var month, day int
	var date []string
	var err error

	if dateStrBeforeFormat == "" {
		postDataEmpty = true
	} else {
		date = strings.Split(dateStrBeforeFormat, "-")
		month, err = strconv.Atoi(date[1])
		if err != nil { postDataEmpty = true }
		day, err = strconv.Atoi(date[2])
		if err != nil { postDataEmpty = true }
	}

	cancel := m.Cancel{LectureID: id, Month: month, Day: day, Comment: comment}

	return cancel, !postDataEmpty
}

func getTestFromPosted(c *gin.Context) (m.Test, bool) {
	lectureId        := c.PostForm("lectureid")
	dateStrBeforeFormat := c.PostForm("date")
	room             := c.PostForm("room")
	comment          := c.PostForm("comment")

	id, _ := strconv.Atoi(lectureId)

	// 空の入力があるか
	postDataEmpty := false

	var month, day int
	var err error

	if dateStrBeforeFormat == "" {
		postDataEmpty = true
	} else {
		date := strings.Split(dateStrBeforeFormat, "-")
		month, err = strconv.Atoi(date[1])
		if err != nil { postDataEmpty = true }
		day, err = strconv.Atoi(date[2])
		if err != nil { postDataEmpty = true }
	}
	if room == "" {
		postDataEmpty = true
	}

	test := m.Test{LectureID: id, Month: month, Day: day, ClassRoom: room, Comment: comment}
	return test, !postDataEmpty
}

// postされた情報からchnageRoomの構造体を返す
func getChangeRoomFromPosted(c *gin.Context) (m.ChangeRoom, bool) {
	lectureId        := c.PostForm("lectureid")
	dateStrBeforeFormat := c.PostForm("date")
	roomAfterChange  := c.PostForm("room")

	id, _ := strconv.Atoi(lectureId)

	// 空の入力があるか
	postDataEmpty := false

	var month, day int
	var err error

	if dateStrBeforeFormat == "" {
		postDataEmpty = true
	} else {
		date := strings.Split(dateStrBeforeFormat, "-")
		month, err = strconv.Atoi(date[1])
		if err != nil { postDataEmpty = true }
		day, err = strconv.Atoi(date[2])
		if err != nil { postDataEmpty = true }
	}
	if roomAfterChange == "" {
		postDataEmpty = true
	}

	room := m.ChangeRoom{LectureID: id, Month: month, Day: day, ClassRoom: roomAfterChange}

	return room, !postDataEmpty
}
