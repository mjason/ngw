package ngw

import (
	"labix.org/v2/mgo"
)

var MongodbUrl string
var MongodbDB string

func initMgo() {
	MongodbUrl = "127.0.0.1"
}

func Session(f func(*mgo.Session)) (err error) {
	session, err := mgo.Dial(MongodbUrl)
	if err != nil {
		return
	}
	session.SetMode(mgo.Monotonic, true)
	f(session)
	session.Close()
	return
}

func Collection(name string, f func(*mgo.Collection)) (err error) {
	err = Session(func(s *mgo.Session) {
		f(s.DB(MongodbDB).C(name))
	})
	return
}
