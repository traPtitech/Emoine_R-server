package dbschema

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// State represents a row from 'emoine.state'.
type State struct {
	ID        uint16         `json:"id"`        // id
	Status    string         `json:"status"`    // status
	Info      sql.NullString `json:"info"`      // info
	CreatedAt sql.NullTime   `json:"createdAt"` // createdAt
	UpdatedAt sql.NullTime   `json:"updatedAt"` // updatedAt
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [State] exists in the database.
func (s *State) Exists() bool {
	return s._exists
}

// Deleted returns true when the [State] has been marked for deletion
// from the database.
func (s *State) Deleted() bool {
	return s._deleted
}

// Insert inserts the [State] to the database.
func (s *State) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO emoine.state (` +
		`status, info, createdAt, updatedAt` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, s.Status, s.Info, s.CreatedAt, s.UpdatedAt)
	res, err := db.ExecContext(ctx, sqlstr, s.Status, s.Info, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	s.ID = uint16(id)
	// set exists
	s._exists = true
	return nil
}

// Update updates a [State] in the database.
func (s *State) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE emoine.state SET ` +
		`status = ?, info = ?, createdAt = ?, updatedAt = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, s.Status, s.Info, s.CreatedAt, s.UpdatedAt, s.ID)
	if _, err := db.ExecContext(ctx, sqlstr, s.Status, s.Info, s.CreatedAt, s.UpdatedAt, s.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [State] to the database.
func (s *State) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for [State].
func (s *State) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO emoine.state (` +
		`id, status, info, createdAt, updatedAt` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`status = VALUES(status), info = VALUES(info), createdAt = VALUES(createdAt), updatedAt = VALUES(updatedAt)`
	// run
	logf(sqlstr, s.ID, s.Status, s.Info, s.CreatedAt, s.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, s.ID, s.Status, s.Info, s.CreatedAt, s.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the [State] from the database.
func (s *State) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM emoine.state ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, s.ID)
	if _, err := db.ExecContext(ctx, sqlstr, s.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// States retrieves all rows from 'emoine.state' as a [State].
func States(ctx context.Context, db DB, limit, int, offset int) ([]State, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, status, info, createdAt, updatedAt ` +
		`FROM emoine.state ` +
		`LIMIT ? OFFSET ?`
	// run
	logf(sqlstr, limit, offset)

	rows, err := db.QueryContext(ctx, sqlstr, limit, offset)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []State
	for rows.Next() {
		s := State{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.ID, &s.Status, &s.Info, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// StateByID retrieves a row from 'emoine.state' as a [State].
//
// Generated from index 'state_id_pkey'.
func StateByID(ctx context.Context, db DB, id uint16) (*State, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, status, info, createdAt, updatedAt ` +
		`FROM emoine.state ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	s := State{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&s.ID, &s.Status, &s.Info, &s.CreatedAt, &s.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}
