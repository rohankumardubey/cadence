// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uber/cadence/common/persistence/sql/sqlplugin"
)

const (
	createDomainQuery = `INSERT INTO 
 domains (id, name, is_global, data, data_encoding)
 VALUES($1, $2, $3, $4, $5)`

	updateDomainQuery = `UPDATE domains 
 SET name = $1, data = $2, data_encoding = $3
 WHERE shard_id=54321 AND id = $4`

	getDomainPart = `SELECT id, name, is_global, data, data_encoding FROM domains`

	getDomainByIDQuery   = getDomainPart + ` WHERE shard_id=$1 AND id = $2`
	getDomainByNameQuery = getDomainPart + ` WHERE shard_id=$1 AND name = $2`

	listDomainsQuery      = getDomainPart + ` WHERE shard_id=$1 ORDER BY id LIMIT $2`
	listDomainsRangeQuery = getDomainPart + ` WHERE shard_id=$1 AND id > $2 ORDER BY id LIMIT $3`

	deleteDomainByIDQuery   = `DELETE FROM domains WHERE shard_id=$1 AND id = $2`
	deleteDomainByNameQuery = `DELETE FROM domains WHERE shard_id=$1 AND name = $2`

	getDomainMetadataQuery    = `SELECT notification_version FROM domain_metadata`
	lockDomainMetadataQuery   = `SELECT notification_version FROM domain_metadata FOR UPDATE`
	updateDomainMetadataQuery = `UPDATE domain_metadata SET notification_version = $1 WHERE notification_version = $2`
)

const (
	shardID = 54321
)

var errMissingArgs = errors.New("missing one or more args for API")

// InsertIntoDomain inserts a single row into domains table
func (pdb *db) InsertIntoDomain(ctx context.Context, row *sqlplugin.DomainRow) (sql.Result, error) {
	return pdb.driver.ExecContext(ctx, sqlplugin.DbDefaultShard, createDomainQuery, row.ID, row.Name, row.IsGlobal, row.Data, row.DataEncoding)
}

// UpdateDomain updates a single row in domains table
func (pdb *db) UpdateDomain(ctx context.Context, row *sqlplugin.DomainRow) (sql.Result, error) {
	return pdb.driver.ExecContext(ctx, sqlplugin.DbDefaultShard, updateDomainQuery, row.Name, row.Data, row.DataEncoding, row.ID)
}

// SelectFromDomain reads one or more rows from domains table
func (pdb *db) SelectFromDomain(ctx context.Context, filter *sqlplugin.DomainFilter) ([]sqlplugin.DomainRow, error) {
	switch {
	case filter.ID != nil || filter.Name != nil:
		return pdb.selectFromDomain(ctx, filter)
	case filter.PageSize != nil && *filter.PageSize > 0:
		return pdb.selectAllFromDomain(ctx, filter)
	default:
		return nil, errMissingArgs
	}
}

func (pdb *db) selectFromDomain(ctx context.Context, filter *sqlplugin.DomainFilter) ([]sqlplugin.DomainRow, error) {
	var err error
	var row sqlplugin.DomainRow
	switch {
	case filter.ID != nil:
		err = pdb.driver.GetContext(ctx, sqlplugin.DbDefaultShard, &row, getDomainByIDQuery, shardID, *filter.ID)
	case filter.Name != nil:
		err = pdb.driver.GetContext(ctx, sqlplugin.DbDefaultShard, &row, getDomainByNameQuery, shardID, *filter.Name)
	}
	if err != nil {
		return nil, err
	}
	return []sqlplugin.DomainRow{row}, err
}

func (pdb *db) selectAllFromDomain(ctx context.Context, filter *sqlplugin.DomainFilter) ([]sqlplugin.DomainRow, error) {
	var err error
	var rows []sqlplugin.DomainRow
	switch {
	case filter.GreaterThanID != nil:
		err = pdb.driver.SelectContext(ctx, sqlplugin.DbDefaultShard, &rows, listDomainsRangeQuery, shardID, *filter.GreaterThanID, *filter.PageSize)
	default:
		err = pdb.driver.SelectContext(ctx, sqlplugin.DbDefaultShard, &rows, listDomainsQuery, shardID, filter.PageSize)
	}
	return rows, err
}

// DeleteFromDomain deletes a single row in domains table
func (pdb *db) DeleteFromDomain(ctx context.Context, filter *sqlplugin.DomainFilter) (sql.Result, error) {
	var err error
	var result sql.Result
	switch {
	case filter.ID != nil:
		result, err = pdb.driver.ExecContext(ctx, sqlplugin.DbDefaultShard, deleteDomainByIDQuery, shardID, filter.ID)
	default:
		result, err = pdb.driver.ExecContext(ctx, sqlplugin.DbDefaultShard, deleteDomainByNameQuery, shardID, filter.Name)
	}
	return result, err
}

// LockDomainMetadata acquires a write lock on a single row in domain_metadata table
func (pdb *db) LockDomainMetadata(ctx context.Context) error {
	var row sqlplugin.DomainMetadataRow
	err := pdb.driver.GetContext(ctx, sqlplugin.DbDefaultShard, &row.NotificationVersion, lockDomainMetadataQuery)
	return err
}

// SelectFromDomainMetadata reads a single row in domain_metadata table
func (pdb *db) SelectFromDomainMetadata(ctx context.Context) (*sqlplugin.DomainMetadataRow, error) {
	var row sqlplugin.DomainMetadataRow
	err := pdb.driver.GetContext(ctx, sqlplugin.DbDefaultShard, &row.NotificationVersion, getDomainMetadataQuery)
	return &row, err
}

// UpdateDomainMetadata updates a single row in domain_metadata table
func (pdb *db) UpdateDomainMetadata(ctx context.Context, row *sqlplugin.DomainMetadataRow) (sql.Result, error) {
	return pdb.driver.ExecContext(ctx, sqlplugin.DbDefaultShard, updateDomainMetadataQuery, row.NotificationVersion+1, row.NotificationVersion)
}
