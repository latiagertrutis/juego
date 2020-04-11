// ///////////////////////////////////////////////////////////////////
// Filename: sql.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sat Apr 11 21:08:03 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	Db             *sql.DB
	GetSpritesheet *sql.Stmt
	GetSprites     *sql.Stmt
}

var (
	GlobalDB Db
)

func (d *Db) Init(Path string) error {
	var err error

	d.Db, err = sql.Open("sqlite3", Path)
	if err != nil {
		return err
	}

	err = d.Db.Ping()
	if err != nil {
		return err
	}

	d.GetSpritesheet, err = d.Db.Prepare("SELECT Path, SpritesNum FROM Spritesheets WHERE ID=?;")
	if err != nil {
		return err
	}
	d.GetSprites, err = d.Db.Prepare("SELECT BoundsX0, BoundsY0, BoundsXf, BoundsYf, SizeX, SizeY  FROM Sprites WHERE SpriteSheet=?;")
	if err != nil {
		return err
	}
	return nil
}
