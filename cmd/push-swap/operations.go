package main

import "fmt"

type Operation struct {
	name string
	exec func(*Stack, *Stack)
}

var operations = map[string]Operation{
	"pa": {"pa", pa},
	"pb": {"pb", pb},
	"sa": {"sa", sa},
	"sb": {"sb", sb},
	"ss": {"ss", ss},
	"ra": {"ra", ra},
	"rb": {"rb", rb},
	"rr": {"rr", rr},
	"rra": {"rra", rra},
	"rrb": {"rrb", rrb},
	"rrr": {"rrr", rrr},
}

// pa: push top of b to a
func pa(a, b *Stack) {
	if value, ok := b.Pop(); ok {
		a.Push(value)
	}
}

// pb: push top of a to b
func pb(a, b *Stack) {
	if value, ok := a.Pop(); ok {
		b.Push(value)
	}
}

// sa: swap first 2 elements of stack a
func sa(a, b *Stack) {
	if len(a.values) >= 2 {
		a.values[len(a.values)-1], a.values[len(a.values)-2] = a.values[len(a.values)-2], a.values[len(a.values)-1]
	}
}

// sb: swap first 2 elements of stack b
func sb(a, b *Stack) {
	if len(b.values) >= 2 {
		b.values[len(b.values)-1], b.values[len(b.values)-2] = b.values[len(b.values)-2], b.values[len(b.values)-1]
	}
}

// ss: execute sa and sb
func ss(a, b *Stack) {
	sa(a, b)
	sb(a, b)
}

// ra: rotate a (shift up all elements by 1, first becomes last)
func ra(a, b *Stack) {
	if len(a.values) >= 2 {
		first := a.values[len(a.values)-1]
		for i := len(a.values) - 1; i > 0; i-- {
			a.values[i] = a.values[i-1]
		}
		a.values[0] = first
	}
}

// rb: rotate b
func rb(a, b *Stack) {
	if len(b.values) >= 2 {
		first := b.values[len(b.values)-1]
		for i := len(b.values) - 1; i > 0; i-- {
			b.values[i] = b.values[i-1]
		}
		b.values[0] = first
	}
}

// rr: execute ra and rb
func rr(a, b *Stack) {
	ra(a, b)
	rb(a, b)
}

// rra: reverse rotate a (shift down all elements by 1, last becomes first)
func rra(a, b *Stack) {
	if len(a.values) >= 2 {
		last := a.values[0]
		for i := 0; i < len(a.values)-1; i++ {
			a.values[i] = a.values[i+1]
		}
		a.values[len(a.values)-1] = last
	}
}

// rrb: reverse rotate b
func rrb(a, b *Stack) {
	if len(b.values) >= 2 {
		last := b.values[0]
		for i := 0; i < len(b.values)-1; i++ {
			b.values[i] = b.values[i+1]
		}
		b.values[len(b.values)-1] = last
	}
}

// rrr: execute rra and rrb
func rrr(a, b *Stack) {
	rra(a, b)
	rrb(a, b)
}

func executeOperation(opName string, a, b *Stack) error {
	if op, exists := operations[opName]; exists {
		op.exec(a, b)
		return nil
	}
	return fmt.Errorf("unknown operation: %s", opName)
}
