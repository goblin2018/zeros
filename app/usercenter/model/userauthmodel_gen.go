// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"

	cacheZerosUserAuthIdPrefix              = "cache:zeros:userAuth:id:"
	cacheZerosUserAuthAuthTypeAuthKeyPrefix = "cache:zeros:userAuth:authType:authKey:"
	cacheZerosUserAuthUserIdAuthTypePrefix  = "cache:zeros:userAuth:userId:authType:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error)
		FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error)
		Update(ctx context.Context, data *UserAuth) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		UserId     int64     `db:"user_id"`
		AuthKey    string    `db:"auth_key"`  // 平台唯一id
		AuthType   string    `db:"auth_type"` // 平台类型
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	zerosUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	zerosUserAuthIdKey := fmt.Sprintf("%s%v", cacheZerosUserAuthIdPrefix, id)
	zerosUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, zerosUserAuthAuthTypeAuthKeyKey, zerosUserAuthIdKey, zerosUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	zerosUserAuthIdKey := fmt.Sprintf("%s%v", cacheZerosUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, zerosUserAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error) {
	zerosUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthAuthTypeAuthKeyPrefix, authType, authKey)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, zerosUserAuthAuthTypeAuthKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `auth_type` = ? and `auth_key` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, authType, authKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error) {
	zerosUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthUserIdAuthTypePrefix, userId, authType)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, zerosUserAuthUserIdAuthTypeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `auth_type` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, authType); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, data *UserAuth) (sql.Result, error) {
	zerosUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	zerosUserAuthIdKey := fmt.Sprintf("%s%v", cacheZerosUserAuthIdPrefix, data.Id)
	zerosUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, zerosUserAuthAuthTypeAuthKeyKey, zerosUserAuthIdKey, zerosUserAuthUserIdAuthTypeKey)
	return ret, err
}

func (m *defaultUserAuthModel) Update(ctx context.Context, newData *UserAuth) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	zerosUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	zerosUserAuthIdKey := fmt.Sprintf("%s%v", cacheZerosUserAuthIdPrefix, data.Id)
	zerosUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheZerosUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.AuthKey, newData.AuthType, newData.Id)
	}, zerosUserAuthAuthTypeAuthKeyKey, zerosUserAuthIdKey, zerosUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheZerosUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
