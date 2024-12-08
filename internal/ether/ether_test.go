// Copyright 2024 The Contributors to go-eth-app
// This file is part of the go-eth-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

package ether

import "fmt"

func Example_wei() {
	w := Wei(1)
	fmt.Println("ToGwei: ", w.ToGweiBF())
	fmt.Println("ToEther: ", w.ToEtherBF())

	// Output:
	// ToGwei:  1e-09
	// ToEther:  1e-18
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
	fmt.Println(e.ToWeiBI())
	fmt.Println(e.ToGweiBI())

	// Output:
	// 2500000000000000000 Exact
	// 2500000000 Exact
}
