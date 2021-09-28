package stores

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

//Photo implement model.PhotoRepo
type Photo struct {
	store *mgo.Session
}

func (p Photo) Get(ctx context.Context, id string) ([]byte, error) {
	reqID := ctx.Value(constant.ReqID).(string)

	select {
	case <-ctx.Done():
		p.store.Close()
		return nil, nil
	default:
		gridFS, err := p.store.DB("").GridFS("photo").Open(id)
		if err != nil {
			log.WithFields(log.Fields{
				"Error":        err,
				constant.ReqID: reqID,
			}).Debugf("Error when GetPhoto")
			return nil, err
		}
		data, err := ioutil.ReadAll(gridFS)
		if err != nil {
			log.WithFields(log.Fields{
				"Error":        err,
				constant.ReqID: reqID,
			}).Debugf("Error when GetPhoto")
			return nil, err
		}
		return data, nil
	}
}

//Store store photo to DB
func (p Photo) Store(ctx context.Context, id string, data []byte) error {
	reqID := ctx.Value(constant.ReqID).(string)

	select {
	case <-ctx.Done():
		p.store.Close()
		return nil
	default:
		gridFS, err := p.store.DB("").GridFS("photo").Create(id)
		if err != nil {
			log.WithFields(log.Fields{
				"Error":        err,
				constant.ReqID: reqID,
			}).Debugf("Error when StorePhoto")
			return err
		}

		sizeFile, err := gridFS.Write(data)
		if err != nil {
			log.WithFields(log.Fields{
				"Error":        err,
				constant.ReqID: reqID,
			}).Debugf("Error when StorePhoto")
			return err
		}
		err = gridFS.Close()
		if err != nil {
			log.WithFields(log.Fields{
				"Error":        err,
				constant.ReqID: reqID,
			}).Debugf("Error when StorePhoto")
			return err
		}
		log.WithFields(log.Fields{
			"Size":         sizeFile,
			constant.ReqID: reqID,
		}).Debugf("Write file")
		return nil
	}
}

//InitPhoto store
func InitPhoto(session *mgo.Session) Photo {
	repo := Photo{
		store: session,
	}
	return repo
}
