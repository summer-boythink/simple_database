package main

// insert struct info
const (
	IdSize         = 8
	userNameSize   = 32
	emailSize      = 255
	IdOffset       = 0
	userNameOffset = IdOffset + IdSize
	EmailOffset    = userNameOffset + userNameSize
	RowSize        = IdSize + userNameSize + emailSize
)

const (
	PageSize     = 4096
	TableMaxPage = 100
	RowsPerPage  = PageSize / RowSize
	TableMaxRows = RowsPerPage * TableMaxPage
)

type Table struct {
	pages   *[TableMaxPage][]byte
	numRows int
}

func (t *Table) RowSlot(RowNum int) (pageNum int, byteOffset int) {
	pageNum = RowNum / RowsPerPage
	page := t.pages[pageNum]
	if page == nil {
		t.pages[pageNum] = make([]byte, PageSize)
	}
	rowOffset := RowNum % RowsPerPage
	byteOffset = rowOffset * RowSize
	// TODO:need byteOffset ?
	return
}


func (t *Table) Serialize(r *Row) {
	pageNum, _ := t.RowSlot(t.numRows)
	t.pages[pageNum] = append(t.pages[pageNum], Int32ToBytes(r.id)...)
	t.pages[pageNum] = append(t.pages[pageNum], r.userName[:]...)
	t.pages[pageNum] = append(t.pages[pageNum], r.email[:]...)
}

func (t *Table) deserializeRow(i int) *Row{
	// t.pages
}

func NewTable() *Table {
	return &Table{
		pages:   &[TableMaxPage][]byte{},
		numRows: 0,
	}
}
