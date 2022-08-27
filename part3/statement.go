package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type StatementType int
type PrepareResult int
type ExecuteResult int

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED_STATEMENT
	PREPARE_SYNTAX_ERROR
)

const (
	STATEMENT_INSERT StatementType = iota + 1
	STATEMENT_SELECT
)

const (
	ExecuteSuccess ExecuteResult = iota
	ExecuteTableFull
	ExecuteError
)

type Row struct {
	id       int
	userName [32]byte
	email    [255]byte
}

type Statement struct {
	Statetype   StatementType
	RowToInsert *Row
}

func NewStatement() *Statement {
	return &Statement{
		Statetype:   0,
		RowToInsert: &Row{},
	}
}

func (s *Statement) PrepareStatement(B *InputBuffer) PrepareResult {
	if strings.Index(B.buffer, "insert") == 0 {
		s.Statetype = STATEMENT_INSERT
		res := strings.Split(B.buffer, " ")
		if len(res) != 4 {
			return PREPARE_SYNTAX_ERROR
		}
		id, err := strconv.Atoi(res[1])
		if err != nil {
			fmt.Println(err)
			return PREPARE_SYNTAX_ERROR
		}
		s.RowToInsert.id = id
		s.RowToInsert.userName = StringToArray[[userNameSize]byte](res[2])
		s.RowToInsert.email = StringToArray[[emailSize]byte](res[3])

		return PREPARE_SUCCESS

	}
	if strings.Index(B.buffer, "select") == 0 {
		s.Statetype = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}
	return PREPARE_UNRECOGNIZED_STATEMENT
}

func (s *Statement) ExecuteStatement(t *Table) ExecuteResult {
	switch s.Statetype {
	case STATEMENT_INSERT:
		return s.ExecuteInsert(t)
	case STATEMENT_SELECT:
		return s.ExecuteSelect(t)
	}
	return ExecuteError
}

func (s *Statement) ExecuteInsert(t *Table) ExecuteResult {
	if t.numRows >= TableMaxRows {
		return ExecuteTableFull
	}
	t.Serialize(s.RowToInsert)
	t.numRows += 1
	return ExecuteSuccess
}

func (s *Statement) ExecuteSelect(t *Table) ExecuteResult {
	for i := 0; i < t.numRows; i++ {
		r := t.deserializeRow(i)
		fmt.Printf("(%d %s %s)\n", r.id, ArrayToString(r.userName), ArrayToString(r.email))
	}
	return ExecuteSuccess
}

func (r *Row) GetRows(t *Table, pageNum int, rowOffset int) {
	byteOffset := rowOffset * RowSize
	r.id = int(binary.BigEndian.Uint32(
		t.pages[pageNum][byteOffset+IdOffset : byteOffset+userNameOffset]))

	userName := [userNameSize]byte{}
	copy(userName[:], t.pages[pageNum][byteOffset+userNameOffset:byteOffset+EmailOffset])
	r.userName = userName
	email := [emailSize]byte{}
	copy(email[:], t.pages[pageNum][byteOffset+EmailOffset:])
	r.email = email
}
