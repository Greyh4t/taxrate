package main

import (
	"fmt"
	"taxrate/insurance"
	"taxrate/tax"
)

var (
	annualBonus                           = 50000       // 年终奖
	specialDeductionData                  = 1500 + 3000 // 专项扣除，租房+赡养老人
	largeAmountOfMutualMedicalExpenseFund = 3           // 大额医疗互助资金
)

const (
	oldMinCardinal                     = 6326  // 旧社保基数下限，1-6月份使用
	oldMaxCardinal                     = 33891 // 旧社保基数上限，公积金基数上限与这个一致，1-6月份使用
	oldMinHousingProvidentFundCardinal = 2420  // 旧公积金基数下限，1-6月份使用

	minCardinal                     = 6821  // 新社保基数下限，7-12月份使用
	maxCardinal                     = 35283 // 新社保基数上限，公积金基数上限与这个一致，7-12月份使用
	minHousingProvidentFundCardinal = 2420  // 新公积金基数下限，7-12月份使用
	housingProvidentFundPercent     = 0.12  // 公积金缴费比例
	taxExemptionAmount              = 5000  // 个税起征点
)

// 税前月薪，按照公司工资单填写，不一定是整数，红包奖金什么的也算在内
func MonthlySalary(month int) float64 {
	switch month {
	case 1:
		return 20000
	case 2:
		return 20000
	case 3:
		return 20000
	case 4:
		return 20000
	case 5:
		return 20000
	case 6:
		return 20000
	case 7:
		return 20000
	case 8:
		return 20000
	case 9:
		return 20000
	case 10:
		return 20000
	case 11:
		return 20000
	case 12:
		return 20000
	default:
		return 20000
	}
}

func init() {
	insurance.Init(insurance.Config{
		OldMinCardinal:                     oldMinCardinal,
		OldMaxCardinal:                     oldMaxCardinal,
		OldMinHousingProvidentFundCardinal: oldMinHousingProvidentFundCardinal,
		MinCardinal:                        minCardinal,
		MaxCardinal:                        maxCardinal,
		MinHousingProvidentFundCardinal:    minHousingProvidentFundCardinal,
		HousingProvidentFundPercent:        housingProvidentFundPercent,
	})
}

// 个税免征额
func TaxExemptionAmount(month int) int {
	return taxExemptionAmount
}

// 专项附加扣除
func SpecialDeduction(month int) int {
	return specialDeductionData
}

// 应退或应补税额=[（综合所得收入额-60000元-“三险一金”等专项扣除-子女教育等专项附加扣除-依法确定的其他扣除-符合条件的公益慈善事业捐赠）×适用税率-速算扣除数]-已预缴税额
func calc() {
	fmt.Println("-------------------全年工资-------------------")
	// 累计税前收入总额
	var totalIncome float64
	// 应纳税所得额
	var taxableIncome float64
	// 累计免征额
	var totalTaxExemptionAmount int
	// 累计专项附加扣除
	var totalSpecialDeduction int
	// 累计社保专项扣除
	var totalInsurance float64

	var lastTaxTotal float64
	for i := 1; i <= 12; i++ {
		salary := MonthlySalary(i)
		taxExemptionAmount := TaxExemptionAmount(i)
		specialDeduction := SpecialDeduction(i)
		insurance := insurance.Default.Total(i, salary, largeAmountOfMutualMedicalExpenseFund)

		totalTaxExemptionAmount += taxExemptionAmount
		totalSpecialDeduction += specialDeduction
		totalInsurance += insurance
		taxableIncome += salary - float64(taxExemptionAmount) - float64(specialDeduction) - insurance
		totalIncome += salary

		taxTotal := calcTax(taxableIncome)

		fmt.Printf("%2d月 收入: %9.2f 社保专项扣除: %.2f 累计社保专项扣除: %8.2f 累计免征额: %5d 累计专项附加扣除: %5d 累计应纳税所得额: %9.2f 累计应纳税额: %9.2f 本期申报税额: %9.2f 税后收入: %9.2f\n",
			i, salary, insurance, totalInsurance, totalTaxExemptionAmount, totalSpecialDeduction, taxableIncome, taxTotal, taxTotal-lastTaxTotal, salary-insurance-(taxTotal-lastTaxTotal))

		lastTaxTotal = taxTotal
	}

	fmt.Println("----------------------------------------------")

	taxRate := tax.Default.Rate(taxableIncome)

	taxTotal := calcTax(taxableIncome)

	actualIncome := totalIncome - totalInsurance - taxTotal

	fmt.Printf("税前收入: %.2f\n", totalIncome)
	fmt.Printf("社保缴费: %.2f\n", totalInsurance)
	fmt.Printf("应纳税所得额: %.2f\n", taxableIncome)
	fmt.Printf("适用税率: %s\n", taxRate.String())
	fmt.Printf("速算扣除数: %d\n", taxRate.QuickCalculationDeduction)
	fmt.Printf("应纳税额: %.2f\n", taxTotal)
	fmt.Printf("税后收入: %.2f\n", actualIncome)

	for i, tr := range tax.Default {
		if i == len(tax.Default)-1 {
			continue
		}
		if tr.Rate >= taxRate.Rate {
			fmt.Printf("%s, 剩余可换购期权额度为: %.2f\n", tr.String(), tr.Quota-taxableIncome)
		}
	}
}

func calcTax(taxableIncome float64) float64 {
	taxRate := tax.Default.Rate(taxableIncome)
	tax := taxableIncome*taxRate.Rate - float64(taxRate.QuickCalculationDeduction)
	return insurance.Float(tax).Round(2)
}

func printTaxRateTable() {
	fmt.Println("-------------全年一次性奖金税率表-------------")
	for _, taxRate := range tax.AnnualBonus {
		taxableIncome := taxRate.Quota
		fmt.Printf("年终奖总额: %9.2f, 税率: %2.f%%\n", taxableIncome, taxRate.Rate*100)
	}
}

// 计算年终奖
func calcAnnualBonus() {
	fmt.Println("--------------------年终奖--------------------")
	// 应纳税所得额
	taxableIncome := float64(annualBonus)

	taxRate := tax.AnnualBonus.Rate(taxableIncome)

	taxTotal := taxableIncome*taxRate.Rate - float64(taxRate.QuickCalculationDeduction)
	fmt.Printf("年终奖总额: %.2f\n", taxableIncome)
	fmt.Printf("适用税率: %s\n", taxRate.String())
	fmt.Printf("应纳税额: %.2f\n", insurance.Float(taxTotal).Round(2))
	fmt.Printf("税后收入: %.2f\n", taxableIncome-taxTotal)

	printTaxRateTable()
}

func main() {
	calc()
	calcAnnualBonus()
}
