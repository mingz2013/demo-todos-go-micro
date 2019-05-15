package datastore

import (
	"gopkg.in/mgo.v2"
)

var sessionInstance *mgo.Session

func CreateSession(host string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	sessionInstance = session
	return session, nil

}

func GetSession() *mgo.Session {
	return sessionInstance
}
