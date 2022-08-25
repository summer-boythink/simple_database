package main

import (
	"fmt"
	"strings"
)

type StatementType int
type PrepareResult int

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED_STATEMENT
)

const (
	STATEMENT_INSERT StatementType = iota + 1
	STATEMENT_SELECT
)

type Statement struct {
	Statetype StatementType
}

func NewStatement() *Statement {
	return &Statement{}
}

func (s *Statement) PrepareStatement(B *InputBuffer) PrepareResult {
	if strings.Index(B.buffer, "insert") == 0 {
		s.Statetype = STATEMENT_INSERT
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
