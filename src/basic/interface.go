package main

import (
  "fmt"
)

type salaryCaculate interface {
  CaculateSalary() int
}

type engineer struct {
  basicpay int
}

type manager struct {
  basicpay, pf int
}

func (e engineer) CaculateSalary() int {
  return e.basicpay
}

func (m manager) CaculateSalary() int {
  return m.basicpay + m.pf
}

func totalExpense(s []salaryCaculate) int {
  expense := 0

  for _, v := range s {
    expense = expense + v.CaculateSalary()
  }

  return expense
}

func main() {
  e1 := engineer{30}
  p1 := manager{30, 10}

  s := []salaryCaculate{e1, p1}
  fmt.Println("Total salary is", totalExpense(s))
}
