package tax

import "fmt"

// 个人所得税税率表（综合所得适用）
var Default TaxRateTable

// 全年一次性奖金税率表
var AnnualBonus TaxRateTable

func init() {
	Default = TaxRateTable{
		{Quota: 36000, Rate: 0.03, QuickCalculationDeduction: 0},
		{Quota: 144000, Rate: 0.1, QuickCalculationDeduction: 2520},
		{Quota: 300000, Rate: 0.2, QuickCalculationDeduction: 16920},
		{Quota: 420000, Rate: 0.25, QuickCalculationDeduction: 31920},
		{Quota: 660000, Rate: 0.30, QuickCalculationDeduction: 52920},
		{Quota: 960000, Rate: 0.35, QuickCalculationDeduction: 85920},
		{Quota: 960001, Rate: 0.45, QuickCalculationDeduction: 181920}, // 大于960000的都是45%
	}
	AnnualBonus = TaxRateTable{
		{Quota: 36000, Rate: 0.03, QuickCalculationDeduction: 0},
		{Quota: 144000, Rate: 0.1, QuickCalculationDeduction: 210},
		{Quota: 300000, Rate: 0.2, QuickCalculationDeduction: 1410},
		{Quota: 420000, Rate: 0.25, QuickCalculationDeduction: 2660},
		{Quota: 660000, Rate: 0.3, QuickCalculationDeduction: 4410},
		{Quota: 960000, Rate: 0.35, QuickCalculationDeduction: 7160},
		{Quota: 960001, Rate: 0.45, QuickCalculationDeduction: 15160}, // 大于960000的都是45%
	}
}

type TaxRateTable []TaxRate

func (t TaxRateTable) Rate(income float64) TaxRate {
	for _, x := range t {
		if income <= x.Quota {
			return x
		}
	}
	return t[len(t)-1]
}

type TaxRate struct {
	Quota                     float64 // 应纳税所得额
	Rate                      float64 // 税率
	QuickCalculationDeduction int     // 速算扣除数
}

func (tr *TaxRate) String() string {
	if tr.Quota > 960000 {
		return fmt.Sprintf("应纳税所得额 %.f 以上 %.f%%", tr.Quota, tr.Rate*100)
	}
	return fmt.Sprintf("应纳税所得额 %.f 以内 %.f%%", tr.Quota, tr.Rate*100)
}
