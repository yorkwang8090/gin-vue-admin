// 自动生成模板RobExchangeSymbo
package robot

// robExchangeSymbo表 结构体  RobExchangeSymbo
type RobExchangeSymbo struct {
	Id         uint   `json:"id" form:"id" gorm:"primarykey;column:id;comment:id;size:10;"`                  //id
	ExchangeId *uint  `json:"exchangeId" form:"exchangeId" gorm:"column:exchange_id;comment:交易所id;size:10;"` //交易所id
	Symbol     string `json:"symbol" form:"symbol" gorm:"column:symbol;comment:币种编码;size:10;"`               //币种编码
	Sort       *uint  `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                       //排序
}

// TableName robExchangeSymbo表 RobExchangeSymbo自定义表名 rob_exchange_symbo
func (RobExchangeSymbo) TableName() string {
	return "rob_exchange_symbo"
}
