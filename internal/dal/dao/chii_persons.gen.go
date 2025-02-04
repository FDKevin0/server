// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNamePerson = "chii_persons"

// Person mapped from table <chii_persons>
type Person struct {
	ID          uint32      `gorm:"column:prsn_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"prsn_id"`
	Name        string      `gorm:"column:prsn_name;type:varchar(255);not null" json:"prsn_name"`
	Type        uint8       `gorm:"column:prsn_type;type:tinyint(4) unsigned;not null;index:prsn_type,priority:1" json:"prsn_type"` // 个人，公司，组合
	Infobox     string      `gorm:"column:prsn_infobox;type:mediumtext;not null" json:"prsn_infobox"`
	Producer    bool        `gorm:"column:prsn_producer;type:tinyint(1);not null;index:prsn_producer,priority:1" json:"prsn_producer"`
	Mangaka     bool        `gorm:"column:prsn_mangaka;type:tinyint(1);not null;index:prsn_mangaka,priority:1" json:"prsn_mangaka"`
	Artist      bool        `gorm:"column:prsn_artist;type:tinyint(1);not null;index:prsn_artist,priority:1" json:"prsn_artist"`
	Seiyu       bool        `gorm:"column:prsn_seiyu;type:tinyint(1);not null;index:prsn_seiyu,priority:1" json:"prsn_seiyu"`
	Writer      bool        `gorm:"column:prsn_writer;type:tinyint(4);not null;index:prsn_writer,priority:1" json:"prsn_writer"`                // 作家
	Illustrator bool        `gorm:"column:prsn_illustrator;type:tinyint(4);not null;index:prsn_illustrator,priority:1" json:"prsn_illustrator"` // 绘师
	Actor       bool        `gorm:"column:prsn_actor;type:tinyint(1);not null;index:prsn_actor,priority:1" json:"prsn_actor"`                   // 演员
	Summary     string      `gorm:"column:prsn_summary;type:mediumtext;not null" json:"prsn_summary"`
	Img         string      `gorm:"column:prsn_img;type:varchar(255);not null" json:"prsn_img"`
	ImgAnidb    string      `gorm:"column:prsn_img_anidb;type:varchar(255);not null" json:"prsn_img_anidb"`
	Comment     uint32      `gorm:"column:prsn_comment;type:mediumint(9) unsigned;not null" json:"prsn_comment"`
	Collects    uint32      `gorm:"column:prsn_collects;type:mediumint(8) unsigned;not null" json:"prsn_collects"`
	Dateline    uint32      `gorm:"column:prsn_dateline;type:int(10) unsigned;not null" json:"prsn_dateline"`
	Lastpost    uint32      `gorm:"column:prsn_lastpost;type:int(11) unsigned;not null" json:"prsn_lastpost"`
	Lock        int8        `gorm:"column:prsn_lock;type:tinyint(4);not null;index:prsn_lock,priority:1" json:"prsn_lock"`
	AnidbID     uint32      `gorm:"column:prsn_anidb_id;type:mediumint(8) unsigned;not null" json:"prsn_anidb_id"`
	Ban         uint8       `gorm:"column:prsn_ban;type:tinyint(3) unsigned;not null;index:prsn_ban,priority:1" json:"prsn_ban"`
	Redirect    uint32      `gorm:"column:prsn_redirect;type:int(10) unsigned;not null" json:"prsn_redirect"`
	Nsfw        bool        `gorm:"column:prsn_nsfw;type:tinyint(1) unsigned;not null" json:"prsn_nsfw"`
	Fields      PersonField `gorm:"foreignKey:prsn_id;polymorphic:Owner;polymorphicValue:prsn" json:"fields"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
