package datatypes

// represents basic hiragana from database
type Hiragana struct {
	HiraId      uint `gorm:"column:Hira_ID;primarykey"`
	Symbol      string
	Translation string
}

// represents extended hiragana from database
type HiraganaExt struct {
	HiraId      uint `gorm:"column:Hira_Ext_ID;primarykey"`
	Symbol      string
	Translation string
}
