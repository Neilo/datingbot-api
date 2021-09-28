package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

var instance *DB

//InitMongo init Data Base
func InitMongo(url string) {
	db, err := NewConnection(url)
	if err != nil {
		log.Panicf("DB init false with err: %s", err.Error())
		panic(err)
	}
	instance = db
}

//GetDB return copy mgo instance
func GetDB() (*DB, error) {
	if instance == nil {
		return nil, IsNotConnected
	}
	return instance, nil
}

//DB is connection struct to mongoDB
//use mgo
type DB struct {
	Session *mgo.Session
}

//ShutDown close the connection
func ShutDown() {
	instance.Session.Close()
	instance = nil
}

//NewConnection return new DB connection
//connection with url
//in url must be pass and user if needed
func NewConnection(url string) (*DB, error) {
	var db = DB{}
	var err error
	db.Session, err = mgo.DialWithTimeout(url, connectionTimeout)
	if err != nil {
		return nil, err
	}
	return &db, nil
}

//IsConnected check connection to mongo db Server
func (db *DB) IsConnected() bool {
	return db.Session != nil
}

//Insert insert document to collection
//if collection is not created this function create collection
func (db *DB) Insert(coll string, v ...interface{}) error {
	if !db.IsConnected() {
		return IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Insert(v...)
}

//Find find document in collection
func (db *DB) Find(coll string, query map[string]interface{}, v interface{}) error {
	if !db.IsConnected() {
		return IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	bsonQuery := bson.M{}

	for k, qv := range query {
		bsonQuery[k] = qv
	}

	return sess.DB("").C(coll).Find(bsonQuery).All(v)
}

//FindByID find document by ID
func (db *DB) FindByID(coll string, id string, v interface{}) bool {
	if !db.IsConnected() {
		return false
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	return mgo.ErrNotFound != sess.DB("").C(coll).FindId(id).One(v)
}

//FindAll find all document in collection
func (db *DB) FindAll(coll string, v interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(bson.M{}).All(v)
}

//FindWithQuery you can call this function with query
//you can must use mgo.bson format
func (db *DB) FindWithQuery(coll string, query interface{}, v interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(query).One(v)
}

//FindWithQueryAll you can find all document in collection with this function
//you can call this function with mgo.bson query
func (db *DB) FindWithQueryAll(coll string, query interface{}, v interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(query).All(v)
}

//RemoveWithIDs delete all document in collection by ids
func (db *DB) RemoveWithIDs(coll string, ids interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var sess = db.Session.Copy()
	defer sess.Close()

	_, err := sess.DB("").C(coll).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})

	return err
}

//Update document by query
//warning you can update all document with this query
func (db *DB) Update(coll string, query interface{}, set interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var err error
	var sess = db.Session.Copy()
	defer sess.Close()

	_, err = sess.DB("").C(coll).UpdateAll(query, set)

	return err
}
