// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameRevisionHistory = "chii_rev_history"

// RevisionHistory mapped from table <chii_rev_history>
type RevisionHistory struct {
	ID        uint32 `gorm:"column:rev_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true;index:rev_id,priority:1" json:"rev_id"`
	Type      uint8  `gorm:"column:rev_type;type:tinyint(3) unsigned;not null;index:rev_crt_id,priority:1;index:rev_id,priority:2" json:"rev_type"` // 条目，角色，人物
	Mid       uint32 `gorm:"column:rev_mid;type:mediumint(8) unsigned;not null;index:rev_crt_id,priority:2" json:"rev_mid"`                         // 对应条目，人物的ID
	TextID    uint32 `gorm:"column:rev_text_id;type:mediumint(9) unsigned;not null" json:"rev_text_id"`
	CreatedAt uint32 `gorm:"column:rev_dateline;type:int(10) unsigned;not null" json:"rev_dateline"`
	CreatorID uint32 `gorm:"column:rev_creator;type:mediumint(8) unsigned;not null;index:rev_crt_creator,priority:1;index:rev_id,priority:3" json:"rev_creator"`
	Summary   string `gorm:"column:rev_edit_summary;type:varchar(200);not null" json:"rev_edit_summary"`
}

// TableName RevisionHistory's table name
func (*RevisionHistory) TableName() string {
	return TableNameRevisionHistory
}
