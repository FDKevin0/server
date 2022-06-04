// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameCharacterComment = "chii_crt_comments"

// CharacterComment mapped from table <chii_crt_comments>
type CharacterComment struct {
	ID          uint32 `gorm:"column:crt_pst_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"crt_pst_id"`
	MentionedID uint32 `gorm:"column:crt_pst_mid;type:mediumint(8) unsigned;not null;index:cmt_crt_id,priority:1" json:"crt_pst_mid"` // 关联人物ID
	UID         uint32 `gorm:"column:crt_pst_uid;type:mediumint(8) unsigned;not null;index:crt_pst_uid,priority:1" json:"crt_pst_uid"`
	Related     uint32 `gorm:"column:crt_pst_related;type:mediumint(8) unsigned;not null;index:crt_pst_related,priority:1" json:"crt_pst_related"`
	Dateline    uint32 `gorm:"column:crt_pst_dateline;type:int(10) unsigned;not null" json:"crt_pst_dateline"`
	Content     string `gorm:"column:crt_pst_content;type:mediumtext;not null" json:"crt_pst_content"`
}

// TableName CharacterComment's table name
func (*CharacterComment) TableName() string {
	return TableNameCharacterComment
}
