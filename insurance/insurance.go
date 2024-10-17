package insurance

import (
	"fmt"
	"math"
)

var Default InsuranceTable

type Config struct {
	OldMinCardinal                     int     // 旧社保基数下限，1-6月份使用
	OldMaxCardinal                     int     // 旧社保基数上限，公积金基数上限与这个一致，1-6月份使用
	OldMinHousingProvidentFundCardinal int     // 旧公积金基数下限，1-6月份使用
	MinCardinal                        int     // 新社保基数下限，7-12月份使用
	MaxCardinal                        int     // 新社保基数上限，公积金基数上限与这个一致，7-12月份使用
	MinHousingProvidentFundCardinal    int     // 新公积金基数下限，7-12月份使用
	HousingProvidentFundPercent        float64 // 公积金缴费比例
}

// 初始化社保数据
func Init(conf Config) {
	Default.endowmentInsurance = append(Default.endowmentInsurance,
		*NewInsurance(conf.OldMinCardinal, conf.OldMaxCardinal, 0.08).SetEffectiveMonth(1, 6))
	Default.medicare = append(Default.medicare,
		*NewInsurance(conf.OldMinCardinal, conf.OldMaxCardinal, 0.02).SetEffectiveMonth(1, 6))
	Default.unemployment = append(Default.unemployment,
		*NewInsurance(conf.OldMinCardinal, conf.OldMaxCardinal, 0.005).SetEffectiveMonth(1, 6))
	Default.housingProvidentFund = append(Default.housingProvidentFund,
		*NewInsurance(conf.OldMinHousingProvidentFundCardinal, conf.OldMaxCardinal,
			conf.HousingProvidentFundPercent).SetEffectiveMonth(1, 6))

	Default.endowmentInsurance = append(Default.endowmentInsurance,
		*NewInsurance(conf.MinCardinal, conf.MaxCardinal, 0.08).SetEffectiveMonth(7, 12))
	Default.medicare = append(Default.medicare,
		*NewInsurance(conf.MinCardinal, conf.MaxCardinal, 0.02).SetEffectiveMonth(7, 12))
	Default.unemployment = append(Default.unemployment,
		*NewInsurance(conf.MinCardinal, conf.MaxCardinal, 0.005).SetEffectiveMonth(7, 12))
	Default.housingProvidentFund = append(Default.housingProvidentFund,
		*NewInsurance(conf.MinHousingProvidentFundCardinal, conf.MaxCardinal,
			conf.HousingProvidentFundPercent).SetEffectiveMonth(7, 12))
}

// 社保表
type InsuranceTable struct {
	endowmentInsurance   []Insurance // 养老
	medicare             []Insurance // 医疗
	unemployment         []Insurance // 失业
	housingProvidentFund []Insurance // 公积金
}

// 社保总额
func (i *InsuranceTable) Total(month int, income float64, largeAmountOfMutualMedicalExpenseFund int) float64 {
	return i.Endowment(month, income) + i.Medicare(month, income, largeAmountOfMutualMedicalExpenseFund) +
		i.UnemploymentBenefits(month, income) + i.HousingProvidentFund(month, income)
}

// 养老
func (i *InsuranceTable) Endowment(month int, income float64) float64 {
	return i.calc(month, income, i.endowmentInsurance).Round(2)
}

// 带大额医疗互助金的医保
func (i *InsuranceTable) Medicare(month int, income float64, largeAmountOfMutualMedicalExpenseFund int) float64 {
	return i.calc(month, income, i.medicare).Round(2) + float64(largeAmountOfMutualMedicalExpenseFund)
}

// 失业
func (i *InsuranceTable) UnemploymentBenefits(month int, income float64) float64 {
	return i.calc(month, income, i.unemployment).Round(2)
}

// 公积金
func (i *InsuranceTable) HousingProvidentFund(month int, income float64) float64 {
	return float64(i.calc(month, income, i.housingProvidentFund).Int())
}

func (i *InsuranceTable) calc(month int, income float64, insurances []Insurance) Float {
	for _, data := range insurances {
		r, ok := data.Total(month, income)
		if ok {
			return r
		}
	}
	panic(fmt.Sprintf("not found data for month: %d", month))
}

type Insurance struct {
	minCardinal       float64 // 缴费基数下限
	maxCardinal       float64 // 缴费基数上限
	percent           float64 // 缴费比例
	minEffectiveMonth int     // 生效月份
	maxEffectiveMonth int     // 生效月份
}

func NewInsurance(minCardinal, maxCardinal int, percent float64) *Insurance {
	return &Insurance{
		minCardinal:       float64(minCardinal),
		maxCardinal:       float64(maxCardinal),
		percent:           percent,
		minEffectiveMonth: 1,
		maxEffectiveMonth: 12,
	}
}

// 设置生效月份
func (i *Insurance) SetEffectiveMonth(min, max int) *Insurance {
	if min < 1 {
		panic(fmt.Sprintf("unexpected month: %d", min))
	}
	if max > 12 {
		panic(fmt.Sprintf("unexpected month: %d", max))
	}

	i.minEffectiveMonth = min
	i.maxEffectiveMonth = max
	return i
}

func (i *Insurance) Total(month int, income float64) (Float, bool) {
	if month < i.minEffectiveMonth || month > i.maxEffectiveMonth {
		return 0, false
	}

	cardinal := income
	if income > i.maxCardinal {
		cardinal = i.maxCardinal
	} else if income < i.minCardinal {
		cardinal = i.minCardinal
	}

	return Float(float64(cardinal) * i.percent), true
}

type Float float64

// 四舍五入 int
func (f Float) Int() int {
	return int(f.Round(0))
}

// 四舍五入，percision 保留几位小数
func (f Float) Round(precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(float64(f)*p+0.5) / p
}
