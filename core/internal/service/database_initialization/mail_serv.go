package database_initialization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	registerHandler(func() {
		sqlList := []string{
			`--  domain 
			CREATE TABLE IF NOT EXISTS domain (
				domain varchar(255) NOT NULL,
				a_record varchar(255) NOT NULL DEFAULT '',
				mailboxes int NOT NULL DEFAULT 50,                       -- Number of emails created
				mailbox_quota BIGINT NOT NULL DEFAULT 5368709120,        -- Default space size of mailbox
				quota BIGINT NOT NULL DEFAULT 10737418240,               -- Domain name quota
				rate_limit INT DEFAULT 12,                               -- How many emails per second
				create_time INT NOT NULL default 0,
				active SMALLINT NOT NULL DEFAULT 1,
				PRIMARY KEY (domain)
			)`,

			`--  mailbox 
			CREATE TABLE IF NOT EXISTS mailbox (
				username varchar(255) NOT NULL,
				password varchar(255) NOT NULL,
				password_encode varchar(255) NOT NULL,
				full_name varchar(255) NOT NULL,
				is_admin smallint NOT NULL DEFAULT 0,
				maildir varchar(255) NOT NULL,
				quota bigint NOT NULL DEFAULT 0,
				local_part varchar(255) NOT NULL,
				domain varchar(255) NOT NULL,
				create_time int NOT NULL default 0,
				update_time int NOT NULL default 0,
				active SMALLINT NOT NULL DEFAULT 1,
				PRIMARY KEY (username)
			)`,

			`--  alias 
			CREATE TABLE IF NOT EXISTS alias (
				address varchar(255) NOT NULL,
				goto text NOT NULL,
				domain varchar(255) NOT NULL,
				create_time int NOT NULL default 0,
				update_time int NOT NULL default 0,
				active smallint NOT NULL DEFAULT 1,
				PRIMARY KEY (address)
			)`,

			`--  bm_bcc 
			CREATE TABLE IF NOT EXISTS bm_bcc (
				id SERIAL PRIMARY KEY,
				type TEXT NOT NULL, -- 'sender' or 'recipient'
				address TEXT NOT NULL,
				goto TEXT NOT NULL,
				domain varchar(255) NOT NULL,
				create_time int NOT NULL default 0,
				update_time int NOT NULL default 0,
				active smallint NOT NULL DEFAULT 1
			)`,

			`--  alias_domain 
			CREATE TABLE IF NOT EXISTS alias_domain (
				alias_domain varchar(255) NOT NULL, 
				target_domain varchar(255) NOT NULL,
				create_time int NOT NULL default 0,
				update_time int NOT NULL default 0,
				active smallint NOT NULL DEFAULT 1,
				PRIMARY KEY (alias_domain)
			)`,
		}

		for _, sql := range sqlList {
			_, err := g.DB().Exec(context.Background(), sql)
			if err != nil {
				g.Log().Error(context.Background(), "Failed to create mail server tables:", err)
				return
			}
		}
	})
}
