package flowcontrol

import "strings"

func ControlFlow(op string) {
	switch {
	case strings.Contains(op, "ifelse"):
		callIfElse()
	case strings.Contains(op, "for"):
		callFor()
	case strings.Contains(op, "switch"):
		callSwitch()
	case strings.Contains(op, "logical"):
		callLogicalOperators()
	case strings.Contains(op, "ex"):
		callExercises()
	}
}

func callIfElse() {
	normalIf()
	normalElse()
	elseIf()
	elseIfElse()
}

func callFor() {
	normalFor()
	doubleFor()
	whileFor()
	forWithBreak()
	forWithContinue()
	forAscii()
}

func callSwitch() {
	baseSwitch()
	statementFallthroughSwitch()
	defaultSwitch()
	doubleValueSwitch()
	typeSwitch()
	initializeSwitch()
}

func callLogicalOperators() {
	orOperator()
	andOperator()
	notOperator()
}

func callExercises() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
}
