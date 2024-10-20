# taxrate

工资税率计算器，默认参数适用于北京市 2024 年数据

> 修改文件中各项参数，编译运行即可输出计算结果

### 示例参数

```go
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
```

### 运行结果

```shell
-------------------全年工资-------------------
 1月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除:  4503.00 累计免征额:  5000 累计专项附加扣除:  4500 累计应纳税所得额:   5997.00 累计应纳税额:    179.91 本期申报税额:    179.91 税后收入:  15317.09
 2月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除:  9006.00 累计免征额: 10000 累计专项附加扣除:  9000 累计应纳税所得额:  11994.00 累计应纳税额:    359.82 本期申报税额:    179.91 税后收入:  15317.09
 3月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 13509.00 累计免征额: 15000 累计专项附加扣除: 13500 累计应纳税所得额:  17991.00 累计应纳税额:    539.73 本期申报税额:    179.91 税后收入:  15317.09
 4月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 18012.00 累计免征额: 20000 累计专项附加扣除: 18000 累计应纳税所得额:  23988.00 累计应纳税额:    719.64 本期申报税额:    179.91 税后收入:  15317.09
 5月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 22515.00 累计免征额: 25000 累计专项附加扣除: 22500 累计应纳税所得额:  29985.00 累计应纳税额:    899.55 本期申报税额:    179.91 税后收入:  15317.09
 6月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 27018.00 累计免征额: 30000 累计专项附加扣除: 27000 累计应纳税所得额:  35982.00 累计应纳税额:   1079.46 本期申报税额:    179.91 税后收入:  15317.09
 7月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 31521.00 累计免征额: 35000 累计专项附加扣除: 31500 累计应纳税所得额:  41979.00 累计应纳税额:   1677.90 本期申报税额:    598.44 税后收入:  14898.56
 8月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 36024.00 累计免征额: 40000 累计专项附加扣除: 36000 累计应纳税所得额:  47976.00 累计应纳税额:   2277.60 本期申报税额:    599.70 税后收入:  14897.30
 9月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 40527.00 累计免征额: 45000 累计专项附加扣除: 40500 累计应纳税所得额:  53973.00 累计应纳税额:   2877.30 本期申报税额:    599.70 税后收入:  14897.30
10月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 45030.00 累计免征额: 50000 累计专项附加扣除: 45000 累计应纳税所得额:  59970.00 累计应纳税额:   3477.00 本期申报税额:    599.70 税后收入:  14897.30
11月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 49533.00 累计免征额: 55000 累计专项附加扣除: 49500 累计应纳税所得额:  65967.00 累计应纳税额:   4076.70 本期申报税额:    599.70 税后收入:  14897.30
12月 收入:  20000.00 社保专项扣除: 4503.00 累计社保专项扣除: 54036.00 累计免征额: 60000 累计专项附加扣除: 54000 累计应纳税所得额:  71964.00 累计应纳税额:   4676.40 本期申报税额:    599.70 税后收入:  14897.30
----------------------------------------------
税前收入: 240000.00
社保缴费: 54036.00
应纳税所得额: 71964.00
适用税率: 应纳税所得额 144000 以内 10%
速算扣除数: 2520
应纳税额: 4676.40
税后收入: 181287.60
应纳税所得额 144000 以内 10%, 剩余可换购期权额度为: 72036.00
应纳税所得额 300000 以内 20%, 剩余可换购期权额度为: 228036.00
应纳税所得额 420000 以内 25%, 剩余可换购期权额度为: 348036.00
应纳税所得额 660000 以内 30%, 剩余可换购期权额度为: 588036.00
应纳税所得额 960000 以内 35%, 剩余可换购期权额度为: 888036.00
--------------------年终奖--------------------
年终奖总额: 50000.00
适用税率: 应纳税所得额 144000 以内 10%
应纳税额: 4790.00
税后收入: 45210.00
-------------全年一次性奖金税率表-------------
年终奖总额:  36000.00, 税率:  3%
年终奖总额: 144000.00, 税率: 10%
年终奖总额: 300000.00, 税率: 20%
年终奖总额: 420000.00, 税率: 25%
年终奖总额: 660000.00, 税率: 30%
年终奖总额: 960000.00, 税率: 35%
年终奖总额: 960001.00, 税率: 45%
```
