package Strategy

import (
	"awesomeProject2/pkg/models/postgresql"
)

type BetMod interface {
	InsertStavki(string, float64, float64, int) error
	CheckStavki(int) (int, error)
	GetStavki(int) (string, error)
	DeleteStavki(int) error
	CheckPolzovatel(int) (int, error)
	GetPolzovatel(int) (float64, error)
	UpdatePolzovatel(float64, int) error
}

// создание пользовательского типа - это первый экземпляр интерфейса
type BetM struct{}

func (b BetM) InsertStavki(r string, s float64, p float64, u int) error {
	return postgresql.InsertStavki(r, s, p, u)
}

func (b BetM) CheckStavki(u int) (int, error) {
	return postgresql.CheckStavki(u)
}

func (b BetM) GetStavki(u int) (string, error) {
	return postgresql.GetStavki(u)
}

func (b BetM) DeleteStavki(id int) error {
	return postgresql.DeleteStavki(id)
}

func (b BetM) CheckPolzovatel(u int) (int, error) {
	return postgresql.CheckPolzovatel(u)
}

func (b BetM) GetPolzovatel(u int) (float64, error) {
	return postgresql.GetPolzovatel(u)
}

func (b BetM) UpdatePolzovatel(p float64, u int) error {
	return postgresql.UpdatePolzovatel(p, u)
}

// создание пользовательского типа - это второй экземпляр интерфейса
type CheckMock struct{}

var CheckStavkiM func(u int) (int, error)
var CheckPolzovatelM func(u int) (int, error)

var GetPolzovatelM func(u int) (float64, error)

var GetStavkiM func(u int) (string, error)

var InsertStavkiM func(r string, b float64, p float64, u int) error

var DeleteStavkiM func(id int) error

var UpdatePolzovatelM func(b float64, u int) error

func (m CheckMock) CheckStavki(u int) (int, error) {
	return CheckStavkiM(u)
}

func (m CheckMock) CheckPolzovatel(u int) (int, error) {
	return CheckPolzovatelM(u)
}

func (m CheckMock) GetPolzovatel(u int) (float64, error) {
	return GetPolzovatelM(u)
}

func (m CheckMock) GetStavki(u int) (string, error) {
	return GetStavkiM(u)
}

func (m CheckMock) InsertStavki(r string, b float64, p float64, u int) error {
	return InsertStavkiM(r, b, p, u)
}

func (m CheckMock) DeleteStavki(id int) error {
	return DeleteStavkiM(id)
}

func (m CheckMock) UpdatePolzovatel(b float64, u int) error {
	return UpdatePolzovatelM(b, u)
}

type Context struct {
	Strategy BetMod
}

func (c *Context) Algorithm(a BetMod) {
	c.Strategy = a
}

func (c *Context) CheckStavki(u int) (int, error) {
	return c.Strategy.CheckStavki(u)
}

func (c *Context) CheckPolzovatel(u int) (int, error) {
	return c.Strategy.CheckPolzovatel(u)
}

func (c *Context) GetPolzovatel(u int) (float64, error) {
	return c.Strategy.GetPolzovatel(u)
}

func (c *Context) GetStavki(u int) (string, error) {
	return c.Strategy.GetStavki(u)
}

func (c *Context) InsertStavki(r string, b float64, p float64, u int) error {
	return c.Strategy.InsertStavki(r, b, p, u)
}

func (c *Context) DeleteStavki(id int) error {
	return c.Strategy.DeleteStavki(id)
}

func (c *Context) UpdatePolzovatel(b float64, u int) error {
	return c.Strategy.UpdatePolzovatel(b, u)
}
