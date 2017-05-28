package dbaccess_test

import (
	"base/DBAccess"
	"base/dbaccess/protocol"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

/*
CREATE TABLE student (
  name varchar(50) DEFAULT NULL,
  age int(11) DEFAULT NULL,
  no int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (no),
  KEY index_student_name (name) USING BTREE
);
*/

var db *dbaccess.DBAccess

func init() {
	db = dbaccess.NewDBAccess("mysql", "root:muchinfo@/test")
	db.ConnDB()
}

func TestDBAccessInsert(t *testing.T) {

	stu := &test.Student{
		Name: proto.String("zouqiang999"),
		Age:  proto.Int32(28),
		NO:   proto.Int32(989),
	}

	err := db.Insert("student", stu)

	if err != nil {
		t.Errorf("Insert err %v", err)
	}
}

func TestDBAccessQuery(t *testing.T) {
	data, err := db.Query("student", []string{}, "", &test.Student{})

	if err == nil {
		for _, v := range data {
			tmp := v.(*test.Student)
			t.Log(tmp.GetAge(), tmp.GetName(), tmp.GetNO())
		}

	} else {
		t.Errorf("Query err ", err)
	}
}

func TestDBAccessUpdate(t *testing.T) {
	stu := &test.Student{
		Name: proto.String("chenyirui"),
		Age:  proto.Int32(30),
	}

	err := db.Update("student", stu, &test.Student{
		Age: proto.Int32(28),
	})

	if err != nil {
		t.Errorf("Update err ", err)
	}
}

func TestDBAccessDelete(t *testing.T) {
	err := db.Delete("student", &test.Student{})

	if err != nil {
		t.Errorf("Delete err ", err)
	}
}

func TestDBAccessCommit(t *testing.T) {
	stu := &test.Student{
		Name: proto.String("xiaoyang"),
		Age:  proto.Int32(20),
		NO:   proto.Int32(0),
	}

	err := db.SetNotAutoCommit()

	if err != nil {
		fmt.Println("SetNotAutoCommit: ", err)
	}

	for i := int32(0); i < 10; i++ {
		stu.NO = proto.Int32(i)
		err = db.Insert("student", stu)

		if err != nil {
			fmt.Println("Insert: ", err)
		}
	}

	db.Commit()
}

func TestDBAccessRollback(t *testing.T) {
	stu := &test.Student{
		Name: proto.String("xiaoyang"),
		Age:  proto.Int32(20),
		NO:   proto.Int32(1),
	}

	err := db.SetNotAutoCommit()

	if err != nil {
		fmt.Println("SetNotAutoCommit: ", err)
	}

	for i := int32(10); i < 20; i++ {
		stu.NO = proto.Int32(i)
		err = db.Insert("student", stu)

		if err != nil {
			fmt.Println("Insert: ", err)
		}
	}

	// 查询
	revData, err := db.Query("student", []string{}, "", &test.Student{})

	if err == nil {
		for _, v := range revData {
			tmp := *v.(*test.Student)
			fmt.Println(tmp.GetName(), tmp.GetName(), tmp.GetAge(), tmp.GetNO())
		}
	}

	db.Rollback()

	// 回滚之后再次查询结果
	revData, err = db.Query("student", []string{}, "", &test.Student{})

	if err == nil {
		for _, v := range revData {
			tmp := *v.(*test.Student)
			fmt.Println(tmp.GetName(), tmp.GetName(), tmp.GetAge(), tmp.GetNO())
		}
	}

}
