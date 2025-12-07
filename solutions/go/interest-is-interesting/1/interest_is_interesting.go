package interest

const negativeBalanceInterest float32 = 3.213
const lowBalanceInterest float32 = 0.5
const midBalanceInterest float32 = 1.621
const highBalanceInterest float32 = 2.475
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return negativeBalanceInterest
	case balance < 1000:
		return lowBalanceInterest
	case balance < 5000:
		return midBalanceInterest
	default:
		return highBalanceInterest
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interest := InterestRate(balance)
	return balance * float64(interest) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var yearsCounter int
	for balance < targetBalance {
		yearsCounter++
		balance = AnnualBalanceUpdate(balance)
	}
	return yearsCounter
}
