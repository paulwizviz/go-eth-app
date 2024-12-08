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

// Package ether contains operations related to unit of Ether
//
// | Unit	             | Wei Equivalent                          | Description                             |
// | ------------------- | --------------------------------------- | --------------------------------------  |
// | Wei                 | 10^0 = 1 wei                            | The smallest unit of Ether.             |
// | Kwei (Babbage)      | 10^3 = 1,000 wei                        | Thousand wei.                           |
// | Mwei (Lovelace)     | 10^6 = 1,000,000 wei                    | Million wei.                            |
// | Gwei (Shannon)      | 10^9 = 1,000,000,000 wei                | Billion wei, often used for gas prices. |
// | Microether (Szabo)  | 10^{12} = 1,000,000,000,000 wei         | Trillion wei.                           |
// | Milliether (Finney) | 10^{15} = 1,000,000,000,000,000 wei     | Quadrillion wei.                        |
// | Ether               | 10^{18} = 1,000,000,000,000,000,000 wei | 1 Ether.                                |
package ether
