package main

import (
	"fmt"
	"strings"
)

type StatementType int
type PrepareResult int
type ExecuteResult int

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED_STATEMENT
)

const (
	STATEMENT_INSERT StatementType = iota + 1
	STATEMENT_SELECT
)

const (
	ExecuteSuccess ExecuteResult = iota
	ExecuteTableFull
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
	return &Statement{}
}

func (s *Statement) PrepareStatement(B *InputBuffer) PrepareResult {
	if strings.Index(B.buffer, "insert") == 0 {
		s.Statetype = STATEMENT_INSERT
		argsAssigned, err := fmt.Scanf("insert %d %s %s",
			s.RowToInsert.id, s.RowToInsert.userName, s.RowToInsert.email)
		if err != nil {
			fmt.Println(err)
		}
		if argsAssigned < 3 {
			return PREPARE_UNRECOGNIZED_STATEMENT
		}
		return PREPARE_SUCCESS

	}
	if strings.Index(B.buffer, "select") == 0 {
		s.Statetype = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}
	return PREPARE_UNRECOGNIZED_STATEMENT
}

func (s *Statement) ExecuteStatement() {
	switch s.Statetype {
	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an insert")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do an select")
	}
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
	for i:= 0;i<t.numRows;i++{
		
	}
}
