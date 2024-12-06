package ethconv

import "math/big"

const (
	kwei          = int64(1_000)                     // wei
	mwei          = int64(1_000_000)                 // wei
	gwei          = int64(1_000_000_000)             // wei
	microetherWei = int64(1_000_000_000_000)         // wei
	millietherWei = int64(1_000_000_000_000_000)     // wei
	etherWei      = int64(1_000_000_000_000_000_000) // wei

	etherGwei = int64(1_000_000_000) // gwei
)

// Wei is the smallest unit of Ether
type Wei int64

// ToGwei returns Wei in big.Float
func (w Wei) ToGwei() *big.Float {
	numerator := big.NewFloat(float64(w))
	denominator := big.NewFloat(float64(gwei))
	result := new(big.Float)
	return result.Quo(numerator, denominator)
}

// ToEther returns Ether in big.Float
func (w Wei) ToEther() *big.Float {
	convFactor := big.NewFloat(float64(etherWei))
	numerator := big.NewFloat(float64(w))
	result := new(big.Float)
	return result.Quo(numerator, convFactor)
}

// Gwei is 1,000,000,000 wei
type Gwei float64

// ToWei returns Wei in Big.Int and accuracy
func (g Gwei) ToWei() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(gwei))
	multiplicant := big.NewFloat(float64(g))
	result := new(big.Float)
	return result.Mul(multiplicant, conversionFactor).Int(nil)
}

// ToEther returns Ether big.Float
func (g Gwei) ToEther() *big.Float {
	conversionFactor := big.NewFloat(float64(etherGwei))
	nominator := big.NewFloat(float64(g))
	result := new(big.Float)
	return result.Quo(nominator, conversionFactor)
}

// Ether is 1,000,000,000,000,000,000 wei
type Ether float64

// ToWei returns wei in big.Int and accuracy in big.Accuracy
func (e Ether) ToWei() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(etherWei))
	multiplicant := big.NewFloat(float64(e))
	return new(big.Float).Mul(multiplicant, conversionFactor).Int(nil)
}

// ToGwei returns gwei in big.Int and accuracy in big.Accuracy
func (e Ether) ToGwei() (*big.Int, big.Accuracy) {
	conversionFactor := big.NewFloat(float64(etherGwei))
	multiplicant := big.NewFloat(float64(e))
	return new(big.Float).Mul(multiplicant, conversionFactor).Int(nil)
}
