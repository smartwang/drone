package datastore

import (
	"github.com/smartwang/drone/model"
	"github.com/smartwang/drone/store/datastore/sql"

	"github.com/russross/meddler"
)

func (db *datastore) PermFind(user *model.User, repo *model.Repo) (*model.Perm, error) {
	stmt := sql.Lookup(db.driver, "perms-find-user-repo")
	data := new(model.Perm)
	err := meddler.QueryRow(db, data, stmt, user.ID, repo.ID)
	return data, err
}

func (db *datastore) PermUpsert(perm *model.Perm) error {
	stmt := sql.Lookup(db.driver, "perms-insert-replace-lookup")
	_, err := db.Exec(stmt,
		perm.UserID,
		perm.Repo,
		perm.Pull,
		perm.Push,
		perm.Admin,
		perm.Synced,
	)
	return err
}

func (db *datastore) PermBatch(perms []*model.Perm) (err error) {
	for _, perm := range perms {
		err = db.PermUpsert(perm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *datastore) PermDelete(perm *model.Perm) error {
	stmt := sql.Lookup(db.driver, "perms-delete-user-repo")
	_, err := db.Exec(stmt, perm.UserID, perm.RepoID)
	return err
}

func (db *datastore) PermFlush(user *model.User, before int64) error {
	stmt := sql.Lookup(db.driver, "perms-delete-user-date")
	_, err := db.Exec(stmt, user.ID, before)
	return err
}
