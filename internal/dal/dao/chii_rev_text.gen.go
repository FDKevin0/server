// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameRevisionText = "chii_rev_text"

// RevisionText mapped from table <chii_rev_text>
type RevisionText struct {
	TextID uint32 `gorm:"column:rev_text_id;type:mediumint(9) unsigned;primaryKey;autoIncrement:true" json:"rev_text_id"`
	Text   []byte `gorm:"column:rev_text;type:mediumblob;not null" json:"rev_text"`
}

// TableName RevisionText's table name
func (*RevisionText) TableName() string {
	return TableNameRevisionText
}
