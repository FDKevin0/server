// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameSubject = "chii_subjects"

// Subject mapped from table <chii_subjects>
type Subject struct {
	ID          uint32       `gorm:"column:subject_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"subject_id"`
	TypeID      uint8        `gorm:"column:subject_type_id;type:smallint(6) unsigned;not null;index:subject_idx_cn,priority:2;index:subject_type_id,priority:1;index:order_by_name,priority:2;index:browser,priority:2" json:"subject_type_id"`
	Name        string       `gorm:"column:subject_name;type:varchar(80);not null;index:subject_name,priority:1;index:order_by_name,priority:5" json:"subject_name"`
	NameCN      string       `gorm:"column:subject_name_cn;type:varchar(80);not null;index:subject_name_cn,priority:1" json:"subject_name_cn"`
	UID         string       `gorm:"column:subject_uid;type:varchar(20);not null" json:"subject_uid"` // isbn / imdb
	Creator     uint32       `gorm:"column:subject_creator;type:mediumint(8) unsigned;not null;index:subject_creator,priority:1" json:"subject_creator"`
	Dateline    uint32       `gorm:"column:subject_dateline;type:int(10) unsigned;not null" json:"subject_dateline"`
	Image       string       `gorm:"column:subject_image;type:varchar(255);not null" json:"subject_image"`
	Platform    uint16       `gorm:"column:subject_platform;type:smallint(6) unsigned;not null;index:subject_platform,priority:1;index:order_by_name,priority:4;index:browser,priority:4" json:"subject_platform"`
	Infobox     string       `gorm:"column:field_infobox;type:mediumtext;not null" json:"field_infobox"`
	Summary     string       `gorm:"column:field_summary;type:mediumtext;not null" json:"field_summary"`            // summary
	Field5      string       `gorm:"column:field_5;type:mediumtext;not null" json:"field_5"`                        // author summary
	Volumes     uint32       `gorm:"column:field_volumes;type:mediumint(8) unsigned;not null" json:"field_volumes"` // 卷数
	Eps         uint32       `gorm:"column:field_eps;type:mediumint(8) unsigned;not null" json:"field_eps"`
	Wish        uint32       `gorm:"column:subject_wish;type:mediumint(8) unsigned;not null" json:"subject_wish"`
	Collect     uint32       `gorm:"column:subject_collect;type:mediumint(8) unsigned;not null" json:"subject_collect"`
	Doing       uint32       `gorm:"column:subject_doing;type:mediumint(8) unsigned;not null" json:"subject_doing"`
	OnHold      uint32       `gorm:"column:subject_on_hold;type:mediumint(8) unsigned;not null" json:"subject_on_hold"` // 搁置人数
	Dropped     uint32       `gorm:"column:subject_dropped;type:mediumint(8) unsigned;not null" json:"subject_dropped"` // 抛弃人数
	Series      bool         `gorm:"column:subject_series;type:tinyint(1) unsigned;not null;index:subject_series,priority:1;index:order_by_name,priority:3;index:browser,priority:3" json:"subject_series"`
	SeriesEntry uint32       `gorm:"column:subject_series_entry;type:mediumint(8) unsigned;not null;index:subject_series_entry,priority:1" json:"subject_series_entry"`
	IdxCn       string       `gorm:"column:subject_idx_cn;type:varchar(1);not null;index:subject_idx_cn,priority:1" json:"subject_idx_cn"`
	Airtime     uint8        `gorm:"column:subject_airtime;type:tinyint(1) unsigned;not null;index:subject_airtime,priority:1" json:"subject_airtime"`
	Nsfw        bool         `gorm:"column:subject_nsfw;type:tinyint(1);not null;index:subject_nsfw,priority:1" json:"subject_nsfw"`
	Ban         uint8        `gorm:"column:subject_ban;type:tinyint(1) unsigned;not null;index:subject_ban,priority:1;index:order_by_name,priority:1;index:browser,priority:1" json:"subject_ban"`
	Fields      SubjectField `gorm:"foreignKey:subject_id;references:field_sid" json:"fields"`
}

// TableName Subject's table name
func (*Subject) TableName() string {
	return TableNameSubject
}
