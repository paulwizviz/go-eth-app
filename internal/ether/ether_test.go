package ether

import "fmt"

func Example_gas() {
	g := Gas(26)
	fmt.Println(g.HexString())

	// Output:
	// 0x1a
}

func Example_wei() {
	w := Wei(1)
	fmt.Println("ToGwei: ", w.ToGweiBF())
	fmt.Println("ToEther: ", w.ToEtherBF())

	w = Wei(26)
	fmt.Println("HexString: ", w.HexString())

	// Output:
	// ToGwei:  1e-09
	// ToEther:  1e-18
	// HexString:  0x1a
}

func Example_gwei() {
	g := Gwei(2.5)
	fmt.Println(g.ToWeiBI())
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
