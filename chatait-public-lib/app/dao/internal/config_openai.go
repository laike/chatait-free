// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ConfigOpenaiDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ConfigOpenaiDao struct {
	gmvc.M                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB              // DB is the raw underlying database management object.
	Table   string              // Table is the table name of the DAO.
	Columns configOpenaiColumns // Columns contains all the columns of Table that for convenient usage.
}

// ConfigOpenaiColumns defines and stores column names for table c_config_openai.
type configOpenaiColumns struct {
	Id        string // ID
	Title     string // 标题
	ApiKey    string // api_key
	Proxy     string // 代理地址
	MaxTokens string // 最大Token
	Status    string // 是否启用 1启用 0不启用
	CallNum   string // 接口调用次数
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

func NewConfigOpenaiDao() *ConfigOpenaiDao {
	return &ConfigOpenaiDao{
		M:     g.DB("default").Model("c_config_openai").Safe(),
		DB:    g.DB("default"),
		Table: "c_config_openai",
		Columns: configOpenaiColumns{
			Id:        "id",
			Title:     "title",
			ApiKey:    "api_key",
			Proxy:     "proxy",
			MaxTokens: "max_tokens",
			Status:    "status",
			CallNum:   "call_num",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
		},
	}
}
