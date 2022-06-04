// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameGroupTopic = "chii_group_topics"

// GroupTopic mapped from table <chii_group_topics>
type GroupTopic struct {
	ID       uint32 `gorm:"column:grp_tpc_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"grp_tpc_id"`
	Gid      uint32 `gorm:"column:grp_tpc_gid;type:mediumint(8) unsigned;not null;index:grp_tpc_gid,priority:1" json:"grp_tpc_gid"`
	UID      uint32 `gorm:"column:grp_tpc_uid;type:mediumint(8) unsigned;not null;index:grp_tpc_uid,priority:1" json:"grp_tpc_uid"`
	Title    string `gorm:"column:grp_tpc_title;type:varchar(80);not null" json:"grp_tpc_title"`
	Dateline uint32 `gorm:"column:grp_tpc_dateline;type:int(10) unsigned;not null" json:"grp_tpc_dateline"`
	Lastpost uint32 `gorm:"column:grp_tpc_lastpost;type:int(10) unsigned;not null;index:grp_tpc_lastpost,priority:1" json:"grp_tpc_lastpost"`
	Replies  uint32 `gorm:"column:grp_tpc_replies;type:mediumint(8) unsigned;not null" json:"grp_tpc_replies"`
	State    bool   `gorm:"column:grp_tpc_state;type:tinyint(1) unsigned;not null" json:"grp_tpc_state"`
	Display  bool   `gorm:"column:grp_tpc_display;type:tinyint(1) unsigned;not null;index:grp_tpc_display,priority:1;default:1" json:"grp_tpc_display"`
}

// TableName GroupTopic's table name
func (*GroupTopic) TableName() string {
	return TableNameGroupTopic
}
