package paginator

import "github.com/jinzhu/gorm"

type QueryParams struct {
	DB      *gorm.DB
	Models  interface{}
	Current uint
	PerPage uint
	OrderBy string
}

type QueryIface interface {
	getPerPage() uint
	getCurrent() uint
	getOrderBy() string
	count() uint
	find() interface{}
}

type Query struct {
	db      *gorm.DB
	models  interface{}
	perPage uint
	current uint
	orderBy string
}

func NewQuery(p QueryParams) QueryIface {
	var per uint = 10
	if p.PerPage != 0 {
		per = p.PerPage
	}
	var current uint = 1
	if p.Current != 0 {
		current = p.Current
	}
	return &Query{
		db:      p.DB,
		models:  p.Models,
		perPage: per,
		current: current,
		orderBy: p.OrderBy,
	}
}

func (q *Query) getPerPage() uint {
	return q.perPage
}
func (q *Query) getCurrent() uint {
	return q.current
}
func (q *Query) getOrderBy() string {
	return q.orderBy
}
func (q *Query) count() uint {
	var cnt uint
	q.db.Find(q.models).Count(&cnt)
	return cnt
}
func (q *Query) find() interface{} {
	offset := q.perPage*(q.current-1) + 1
	query := q.db.Limit(q.perPage).Offset(offset)
	if q.orderBy != "" {
		query = query.Order(q.orderBy)
	}
	query.Find(q.models)
	return q.models
}
