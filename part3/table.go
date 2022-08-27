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
	numRows int
	pages   *[TableMaxPage][]byte
}

func (t *Table) RowSlot(RowNum int) (pageNum int, rowOffset int) {
	pageNum = RowNum / RowsPerPage
	page := t.pages[pageNum]
	if page == nil {
		t.pages[pageNum] = make([]byte, 0)
	}
	rowOffset = RowNum % RowsPerPage
	return
}

func (t *Table) Serialize(r *Row) {
	pageNum, _ := t.RowSlot(t.numRows)
	t.pages[pageNum] = append(t.pages[pageNum], Int32ToBytes(r.id)...)
	t.pages[pageNum] = append(t.pages[pageNum], r.userName[:]...)
	t.pages[pageNum] = append(t.pages[pageNum], r.email[:]...)
}

func (t *Table) deserializeRow(i int) *Row {
	row := &Row{}
	pageNum, rowOffset := t.RowSlot(i)
	row.GetRows(t, pageNum, rowOffset)
	return row
}

func NewTable() *Table {
	return &Table{
		pages:   &[TableMaxPage][]byte{},
		numRows: 0,
	}
}
