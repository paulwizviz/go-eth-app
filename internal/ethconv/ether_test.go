package ethconv

import "fmt"

func Example_wei() {
	w := Wei(1)
	fmt.Println("ToGwei: ", w.ToGwei())
	fmt.Println("ToEther: ", w.ToEther())

	// Output:
	// ToGwei:  1e-09
	// ToEther:  1e-18
}

func Example_gwei() {
	g := Gwei(2.5)
	fmt.Println(g.ToWei())
	fmt.Println(g.ToEther())

	// Output:
	// 2500000000 Exact
	// 2.5e-09
}

func Example_ether() {
	e := Ether(2.5)
	fmt.Println(e.ToWei())
	fmt.Println(e.ToGwei())

	// Output:
	// 2500000000000000000 Exact
	// 2500000000 Exact
}
