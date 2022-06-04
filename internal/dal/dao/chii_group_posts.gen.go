// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameGroupTopicComment = "chii_group_posts"

// GroupTopicComment mapped from table <chii_group_posts>
type GroupTopicComment struct {
	ID          uint32 `gorm:"column:grp_pst_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"grp_pst_id"`
	MentionedID uint32 `gorm:"column:grp_pst_mid;type:mediumint(8) unsigned;not null;index:pss_topic_id,priority:1" json:"grp_pst_mid"`
	UID         uint32 `gorm:"column:grp_pst_uid;type:mediumint(8) unsigned;not null;index:grp_pst_uid,priority:1" json:"grp_pst_uid"`
	Related     uint32 `gorm:"column:grp_pst_related;type:mediumint(8) unsigned;not null;index:grp_pst_related,priority:1" json:"grp_pst_related"` // 关联回复ID
	Content     string `gorm:"column:grp_pst_content;type:mediumtext;not null" json:"grp_pst_content"`
	State       bool   `gorm:"column:grp_pst_state;type:tinyint(1) unsigned;not null" json:"grp_pst_state"`
	Dateline    uint32 `gorm:"column:grp_pst_dateline;type:int(10) unsigned;not null" json:"grp_pst_dateline"`
}

// TableName GroupTopicComment's table name
func (*GroupTopicComment) TableName() string {
	return TableNameGroupTopicComment
}
