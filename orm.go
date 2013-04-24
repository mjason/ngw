package ngw

import (
	"labix.org/v2/mgo"
)

var _ip string
var _db string

func InitORM(ip, db string) {
	_ip = ip
	_db = db
}

type ORM struct {
	IP string
	DB string
}

func GetOrm() ORM {
	return ORM{_ip, _db}
}

func (self ORM) Session() (session *mgo.Session, err error) {
	session, err = mgo.Dial(self.IP)
	if err != nil {
		return
	}
	session.SetMode(mgo.Monotonic, true)
	return
}

func (self ORM) GetDB() (session *mgo.Session, db *mgo.Database, err error) {
	session, err = self.Session()
	if err != nil {
		return
	}
	db = session.DB(self.DB)
	return
}

type TableOrm struct {
	TableName string
	orm       ORM
}

func SetTable(tableName string) TableOrm {
	tableOrm := TableOrm{TableName: tableName}
	tableOrm.orm = GetOrm()
	return tableOrm
}

func (self TableOrm) Save(data interface{}) (err error) {
	session, db, err := self.orm.GetDB()
	if err != nil {
		return
	}
	defer session.Close()
	c := db.C(self.TableName)
	err = c.Insert(data)
	if err != nil {
		return
	}
	return
}

func (self TableOrm) FindAll(query, data interface{}) (err error) {
	session, db, err := self.orm.GetDB()
	if err != nil {
		return
	}
	defer session.Close()
	c := db.C(self.TableName)
	c.Find(query).All(data)
	return
}

func (self TableOrm) FindOne(query, data interface{}) (err error) {
	session, db, err := self.orm.GetDB()
	if err != nil {
		return
	}
	defer session.Close()
	c := db.C(self.TableName)
	c.Find(query).One(data)
	return
}

// 验证数据是否已经存在
func (self TableOrm) ISExist(query interface{}) (status bool, e error) {
	session, db, err := self.orm.GetDB()
	if err != nil {
		return
	}
	defer session.Close()
	c := db.C(self.TableName)
	count, err := c.Find(query).Count()
	if err != nil {
		return
	}
	if count == 0 {
		// 不存在
		status = false
	} else {
		// 存在
		status = true
	}
	return

}

func (self TableOrm) Update(query, data interface{}) (err error) {
	session, db, err := self.orm.GetDB()
	if err != nil {
		return
	}
	defer session.Close()
	c := db.C(self.TableName)
	err = c.Update(query, data)
	if err != nil {
		return
	}
	return
}
