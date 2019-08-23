package paginator

type Meta struct {
	Total   uint
	PerPage uint
	Current uint
}

type Paginator struct {
	Meta *Meta
	Data interface{}
}

func Paginate(q QueryIface) *Paginator {
	meta := Meta{
		Total:   q.count(),
		PerPage: q.getPerPage(),
		Current: q.getCurrent(),
	}
	return &Paginator{
		Meta: &meta,
		Data: q.find(),
	}
}
