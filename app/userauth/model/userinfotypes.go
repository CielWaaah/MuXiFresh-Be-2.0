package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Avatar      string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	NickName    string             `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	StudentID   string             `bson:"student_id,omitempty" json:"student_id,omitempty"`
	UserType    string             `bson:"user_type,omitempty" json:"user_type,omitempty"`
	EntryFormID primitive.ObjectID `bson:"entry_form_id,omitempty" json:"entry_form_id,omitempty"`
	ScheduleID  primitive.ObjectID `bson:"schedule_id,omitempty" json:"schedule_id"`
	TestChoice  []string           `bson:"test_choice,omitempty" json:"test_choice,omitempty"`
	TestResult  struct {
		LeQunXing   int64
		YouHengXing int64
		XingFenXing int64
		CongHuiXing int64
		JiaoJiXing  int64
		HuaiYiXing  int64
		WenDingXing int64
	} `bson:"test_result,omitempty" json:"test_result,omitempty"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
