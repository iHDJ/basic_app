package enum

type EnumerationsQueryForm struct {
	ID       int
	Name     string
	IsSystem bool
	IsOpened bool

	Page     int
	PageSize int
}
