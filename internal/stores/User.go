package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/brotherhood228/dating-bot-api/internal/utils"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const userTable = "user"

//UserStore implementation model.UserRepository
type UserStore struct {
	db    *sqlx.DB
	table string
}

//InitUserStore
// return UserStore with userTable as table
// db is *sqlx.DB
func InitUserStore(db *sqlx.DB) UserStore {
	s := UserStore{
		db:    db,
		table: userTable,
	}
	return s
}

//Store save user in DB
//use sqlx.DB to store
//perfect use pool Connections
func (store UserStore) Store(ctx context.Context, user *model.User) (string, error) {
	tx := store.db.MustBegin()
	reqID := utils.GetReqIDFromContext(ctx)
	log.WithFields(log.Fields{
		constant.ReqID: reqID,
	}).Debug("Enter in userRepo.Store")

	select {
	case <-ctx.Done():
		err := tx.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Store Err when commit tx.")
		}
		log.WithFields(log.Fields{
			constant.ReqID: reqID,
		}).Debug("Cancel in userRepo.Store")
		return "", nil
	default:
		_, err := tx.NamedExec(`INSERT INTO public."bot_users" (id,chat_id,name,description,age,gender,find_sex,rating,longitude,latitude) VALUES (:id,:chat_id,:name,:description,:age,:gender,:find_sex,:rating,:longitude,:latitude)`,
			user)
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Store Err when create user.")
			errCommit := tx.Commit()
			if errCommit != nil {
				log.WithFields(log.Fields{
					constant.ReqID: reqID,
					"error":        err.Error(),
				}).Error("userRepo.Store Err when commit tx.")
			}
			return "", err
		}
		err = tx.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Store Err when commit tx.")
			return "", err
		}
		return user.UserNameTG, err
	}
}

//Update user in DB
//use sqlx.DB
//perfect use pool Connections
//update all fields except id(UserNameTG)
func (store UserStore) Update(ctx context.Context, id *string, user *model.User) error {
	tx := store.db.MustBegin()
	reqID := utils.GetReqIDFromContext(ctx)
	log.WithFields(log.Fields{
		constant.ReqID: reqID,
	}).Debug("Enter in userRepo.Update")

	select {
	case <-ctx.Done():
		err := tx.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Update Err when commit tx.")
		}
		log.WithFields(log.Fields{
			constant.ReqID: reqID,
		}).Debug("Cancel in userRepo.Update")
		return nil
	default:
		/*query := fmt.Sprintf(`
			UPDATE public."bot_users"
			   SET  chat_id = :chat_id,
					name = :name,
					description = :description,
					age = :age,
					gender = :gender,
					find_sex = :find_sex,
					rating = :rating,
					longitude = :longitude,
					latitude = :latitude
			 WHERE id = '%s';
		`, *id)*/
		_, err := tx.NamedExec(`UPDATE public."bot_users"    SET  chat_id = :chat_id,name = :name,description = :description,age = :age,gender = :gender,find_sex = :find_sex,rating = :rating,longitude = :longitude,latitude = :latitude WHERE id = :id;`, user)
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Update Err when create user.")
			errCommit := tx.Commit()
			if errCommit != nil {
				log.WithFields(log.Fields{
					constant.ReqID: reqID,
					"error":        err.Error(),
				}).Error("userRepo.Update Err when commit tx.")
			}
			return err
		}
		err = tx.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Update Err when commit tx.")
			return err
		}
		return err
	}
}

//Find user in DB
//use sqlx.DB
//perfect use pool Connections
//find by id(UserNameTG)
func (store UserStore) Find(ctx context.Context, id *string) (*model.User, error) {
	tx := store.db.MustBegin()
	reqID := utils.GetReqIDFromContext(ctx)
	log.WithFields(log.Fields{
		constant.ReqID: reqID,
	}).Debug("Enter in userRepo.Find")

	select {
	case <-ctx.Done():
		err := tx.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Find Err when commit tx.")
		}
		log.WithFields(log.Fields{
			constant.ReqID: reqID,
		}).Debug("Cancel in userRepo.Find")
		return nil, err
	default:
		query := fmt.Sprintf(`
			SELECT * FROM bot_users WHERE id='%s';
		`, *id)
		var user model.User
		rows, err := tx.Queryx(query)

		for rows.Next() {
			err = rows.StructScan(&user)
			if err != nil {
				return nil, err
			}
		}
		errCommit := tx.Commit()
		if errCommit != nil {
			log.WithFields(log.Fields{
				constant.ReqID: reqID,
				"error":        err.Error(),
			}).Error("userRepo.Find Err when commit tx.")
		}
		return &user, err
	}
}

//Pick выбирает рандомного пользователя, которого еще не лайкал пользователь
func (store UserStore) Pick(ctx context.Context, id *string) (*model.User, error) {
	//query := fmt.Sprintf(`
	//		SELECT *
	//		FROM   bot_users
	//		WHERE  id NOT IN (
	//			SELECT object
	//			FROM   bot_likes
	//			WHERE  subject = '%s'
	//		)
	//		AND gender IN (
	//			SELECT find_sex
	//			FROM   bot_users
	//			WHERE  id = '%s'
	//		) order by random() limit 1;
	//	`, *id, *id)

	var user model.User
	err := store.db.GetContext(ctx, &user, `SELECT * FROM   bot_users WHERE  id NOT IN (SELECT object FROM   bot_likes	WHERE subject = $1 ) AND gender IN (SELECT find_sex FROM bot_users WHERE  id = $2 ) AND id NOT IN (SELECT id FROM bot_users WHERE id = $3) order by random() limit 1;`, id, id, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}
