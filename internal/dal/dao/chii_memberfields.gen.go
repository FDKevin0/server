// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameMemberField = "chii_memberfields"

// MemberField mapped from table <chii_memberfields>
type MemberField struct {
	UID       uint32 `gorm:"column:uid;type:mediumint(8) unsigned;primaryKey;autoIncrement:false;default:0" json:"uid"`
	Site      string `gorm:"column:site;type:varchar(75);not null" json:"site"`
	Location  string `gorm:"column:location;type:varchar(30);not null" json:"location"`
	Bio       string `gorm:"column:bio;type:text;not null" json:"bio"`
	Privacy   string `gorm:"column:privacy;type:mediumtext;not null" json:"privacy"`
	Blocklist string `gorm:"column:blocklist;type:mediumtext;not null" json:"blocklist"`
}

// TableName MemberField's table name
func (*MemberField) TableName() string {
	return TableNameMemberField
}