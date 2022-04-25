from typing import List

from sqlalchemy import Date, Enum, Float, Index, Table, Column, String, and_, text
from sqlalchemy.orm import foreign, relationship, declarative_base
from sqlalchemy.dialects.mysql import (
    ENUM,
    YEAR,
    INTEGER,
    TINYINT,
    VARCHAR,
    SMALLINT,
    MEDIUMINT,
    MEDIUMTEXT,
)

Base = declarative_base()
metadata = Base.metadata


class ChiiCharacter(Base):
    __tablename__ = "chii_characters"

    crt_id = Column(MEDIUMINT(8), primary_key=True)
    crt_name = Column(String(255, "utf8_unicode_ci"), nullable=False)
    crt_role = Column(TINYINT(4), nullable=False, index=True, comment="角色，机体，组织。。")
    crt_infobox: str = Column(MEDIUMTEXT, nullable=False)
    crt_summary: str = Column(MEDIUMTEXT, nullable=False)
    crt_img = Column(String(255, "utf8_unicode_ci"), nullable=False)
    crt_comment = Column(MEDIUMINT(9), nullable=False, server_default=text("'0'"))
    crt_collects = Column(MEDIUMINT(8), nullable=False)
    crt_dateline = Column(INTEGER(10), nullable=False)
    crt_lastpost = Column(INTEGER(11), nullable=False)
    crt_lock = Column(
        TINYINT(4), nullable=False, index=True, server_default=text("'0'")
    )
    crt_img_anidb = Column(VARCHAR(255), nullable=False)
    crt_anidb_id = Column(MEDIUMINT(8), nullable=False)
    crt_ban = Column(TINYINT(3), nullable=False, index=True, server_default=text("'0'"))
    crt_redirect = Column(INTEGER(10), nullable=False, server_default=text("'0'"))
    crt_nsfw = Column(TINYINT(1), nullable=False)

    subjects: List["ChiiCrtSubjectIndex"] = relationship(
        "ChiiCrtSubjectIndex",
        primaryjoin=(
            lambda: ChiiCharacter.crt_id == foreign(ChiiCrtSubjectIndex.crt_id)
        ),
        lazy="raise_on_sql",
        back_populates="character",
    )  # type: ignore

    field: "ChiiPersonField" = relationship(
        "ChiiPersonField",
        lazy="raise_on_sql",
        primaryjoin=(lambda: foreign(ChiiPersonField.prsn_id) == ChiiCharacter.crt_id),
        uselist=False,
    )  # type: ignore


class ChiiCrtCastIndex(Base):
    __tablename__ = "chii_crt_cast_index"

    crt_id = Column(MEDIUMINT(9), primary_key=True, nullable=False)
    prsn_id: int = Column(MEDIUMINT(9), primary_key=True, nullable=False, index=True)
    subject_id: int = Column(MEDIUMINT(9), primary_key=True, nullable=False, index=True)
    subject_type_id: int = Column(
        TINYINT(3), nullable=False, index=True, comment="根据人物归类查询角色，动画，书籍，游戏"
    )
    summary = Column(
        String(255, "utf8_unicode_ci"), nullable=False, comment="幼年，男乱马，女乱马，变身形态，少女形态。。"
    )


class ChiiCrtSubjectIndex(Base):
    __tablename__ = "chii_crt_subject_index"

    crt_id = Column(MEDIUMINT(9), primary_key=True, nullable=False)
    subject_id = Column(MEDIUMINT(9), primary_key=True, nullable=False, index=True)
    subject_type_id = Column(TINYINT(4), nullable=False, index=True)
    crt_type: int = Column(TINYINT(4), nullable=False, index=True, comment="主角，配角")
    ctr_appear_eps = Column(MEDIUMTEXT, nullable=False, comment="可选，角色出场的的章节")
    crt_order = Column(TINYINT(3), nullable=False)

    character: "ChiiCharacter" = relationship(
        "ChiiCharacter",
        lazy="raise_on_sql",
        primaryjoin=(
            lambda: ChiiCharacter.crt_id == foreign(ChiiCrtSubjectIndex.crt_id)
        ),
        innerjoin=True,
        uselist=False,
        back_populates="subjects",
    )  # type: ignore
    subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiCrtSubjectIndex.subject_id)
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
        back_populates="characters",
    )  # type: ignore


class ChiiEpisode(Base):
    __tablename__ = "chii_episodes"
    __table_args__ = (Index("ep_subject_id_2", "ep_subject_id", "ep_ban", "ep_sort"),)

    ep_id = Column(MEDIUMINT(8), primary_key=True)
    ep_subject_id = Column(MEDIUMINT(8), nullable=False, index=True)
    ep_sort: float = Column(
        Float, nullable=False, index=True, server_default=text("'0'")
    )
    ep_type = Column(TINYINT(1), nullable=False)
    ep_disc = Column(
        TINYINT(3),
        nullable=False,
        index=True,
        server_default=text("'0'"),
        comment="碟片数",
    )
    ep_name = Column(String(80), nullable=False)
    ep_name_cn = Column(String(80), nullable=False)
    ep_rate = Column(TINYINT(3), nullable=False)
    ep_duration = Column(String(80), nullable=False)
    ep_airdate = Column(String(80), nullable=False)
    ep_online = Column(MEDIUMTEXT, nullable=False)
    ep_comment = Column(MEDIUMINT(8), nullable=False)
    ep_resources = Column(MEDIUMINT(8), nullable=False)
    ep_desc = Column(MEDIUMTEXT, nullable=False)
    ep_dateline = Column(INTEGER(10), nullable=False)
    ep_lastpost = Column(INTEGER(10), nullable=False, index=True)
    ep_lock = Column(TINYINT(3), nullable=False, server_default=text("'0'"))
    ep_ban = Column(TINYINT(3), nullable=False, index=True, server_default=text("'0'"))

    subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiEpisode.ep_subject_id)
        ),
        foreign_keys=[ep_subject_id],
        innerjoin=True,
        lazy="raise_on_sql",
        back_populates="episodes",
        uselist=False,
    )  # type: ignore


t_chii_person_alias = Table(
    "chii_person_alias",
    metadata,
    Column("prsn_cat", ENUM("prsn", "crt"), nullable=False),
    Column("prsn_id", MEDIUMINT(9), nullable=False, index=True),
    Column("alias_name", String(255, "utf8_unicode_ci"), nullable=False),
    Column("alias_type", TINYINT(4), nullable=False),
    Column("alias_key", String(10, "utf8_unicode_ci"), nullable=False),
    Index("prsn_cat", "prsn_cat", "prsn_id"),
)


class ChiiPersonCollect(Base):
    __tablename__ = "chii_person_collects"
    __table_args__ = (
        Index("prsn_clt_cat", "prsn_clt_cat", "prsn_clt_mid"),
        {"comment": "人物收藏"},
    )

    prsn_clt_id = Column(MEDIUMINT(8), primary_key=True)
    prsn_clt_cat = Column(Enum("prsn", "crt"), nullable=False)
    prsn_clt_mid = Column(MEDIUMINT(8), nullable=False, index=True)
    prsn_clt_uid = Column(MEDIUMINT(8), nullable=False, index=True)
    prsn_clt_dateline = Column(INTEGER(10), nullable=False)


class ChiiPersonCsIndex(Base):
    __tablename__ = "chii_person_cs_index"
    __table_args__ = {"comment": "subjects' credits/creator & staff (c&s)index"}

    prsn_type = Column(ENUM("prsn", "crt"), primary_key=True, nullable=False)
    prsn_id = Column(
        MEDIUMINT(9),
        primary_key=True,
        nullable=False,
        index=True,
    )
    prsn_position: int = Column(
        SMALLINT(5), primary_key=True, nullable=False, index=True, comment="监督，原案，脚本,.."
    )
    subject_id = Column(
        MEDIUMINT(9),
        primary_key=True,
        nullable=False,
        index=True,
    )
    subject_type_id: int = Column(TINYINT(4), nullable=False, index=True)
    summary = Column(MEDIUMTEXT, nullable=False)
    prsn_appear_eps = Column(MEDIUMTEXT, nullable=False, comment="可选，人物参与的章节")

    person: "ChiiPerson" = relationship(
        "ChiiPerson",
        primaryjoin=lambda: and_(
            ChiiPerson.prsn_ban == 0,
            foreign(ChiiPersonCsIndex.prsn_id) == ChiiPerson.prsn_id,
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
    )  # type: ignore
    subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=(
            lambda: foreign(ChiiPersonCsIndex.subject_id) == ChiiSubject.subject_id
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
        back_populates="persons",
    )  # type: ignore


class ChiiPersonField(Base):
    __tablename__ = "chii_person_fields"
    __table_args__ = {"extend_existing": True}

    prsn_id = Column(INTEGER(8), primary_key=True, nullable=False, index=True)
    prsn_cat = Column(ENUM("prsn", "crt"), nullable=False)
    gender = Column(TINYINT(4), nullable=False)
    bloodtype = Column(TINYINT(4), nullable=False)
    birth_year = Column(YEAR(4), nullable=False)
    birth_mon = Column(TINYINT(2), nullable=False)
    birth_day = Column(TINYINT(2), nullable=False)
    __mapper_args__ = {"polymorphic_on": prsn_cat, "polymorphic_identity": "prsn"}


class ChiiCharacterField(ChiiPersonField):
    __mapper_args__ = {"polymorphic_identity": "crt"}


t_chii_person_relationship = Table(
    "chii_person_relationship",
    metadata,
    Column("prsn_type", ENUM("prsn", "crt"), nullable=False),
    Column("prsn_id", MEDIUMINT(9), nullable=False),
    Column("relat_prsn_type", ENUM("prsn", "crt"), nullable=False),
    Column("relat_prsn_id", MEDIUMINT(9), nullable=False),
    Column("relat_type", SMALLINT(6), nullable=False, comment="任职于，从属,聘用,嫁给，"),
    Index("relat_prsn_type", "relat_prsn_type", "relat_prsn_id"),
    Index("prsn_type", "prsn_type", "prsn_id"),
)


class ChiiPerson(Base):
    __tablename__ = "chii_persons"
    __table_args__ = {"comment": "（现实）人物表"}

    prsn_id = Column(MEDIUMINT(8), primary_key=True)
    prsn_name = Column(String(255, "utf8_unicode_ci"), nullable=False)
    prsn_type = Column(TINYINT(4), nullable=False, index=True, comment="个人，公司，组合")
    prsn_infobox: str = Column(MEDIUMTEXT, nullable=False)
    prsn_producer = Column(TINYINT(1), nullable=False, index=True)
    prsn_mangaka = Column(TINYINT(1), nullable=False, index=True)
    prsn_artist = Column(TINYINT(1), nullable=False, index=True)
    prsn_seiyu = Column(TINYINT(1), nullable=False, index=True)
    prsn_writer = Column(
        TINYINT(4), nullable=False, index=True, server_default=text("'0'"), comment="作家"
    )
    prsn_illustrator = Column(
        TINYINT(4), nullable=False, index=True, server_default=text("'0'"), comment="绘师"
    )
    prsn_actor = Column(TINYINT(1), nullable=False, index=True, comment="演员")
    prsn_summary: str = Column(MEDIUMTEXT, nullable=False)
    prsn_img = Column(String(255, "utf8_unicode_ci"), nullable=False)
    prsn_img_anidb = Column(VARCHAR(255), nullable=False)
    prsn_comment = Column(MEDIUMINT(9), nullable=False)
    prsn_collects = Column(MEDIUMINT(8), nullable=False)
    prsn_dateline = Column(INTEGER(10), nullable=False)
    prsn_lastpost = Column(INTEGER(11), nullable=False)
    prsn_lock = Column(TINYINT(4), nullable=False, index=True)
    prsn_anidb_id = Column(MEDIUMINT(8), nullable=False)
    prsn_ban = Column(
        TINYINT(3), nullable=False, index=True, server_default=text("'0'")
    )
    prsn_redirect = Column(INTEGER(10), nullable=False, server_default=text("'0'"))
    prsn_nsfw = Column(TINYINT(1), nullable=False)

    subjects: List["ChiiPersonCsIndex"] = relationship(
        "ChiiPersonCsIndex",
        primaryjoin=lambda: and_(
            ChiiPerson.prsn_ban == 0,
            foreign(ChiiPersonCsIndex.prsn_id) == ChiiPerson.prsn_id,
        ),
        lazy="raise_on_sql",
        back_populates="person",
    )


class ChiiSubjectField(Base):
    __tablename__ = "chii_subject_fields"
    __table_args__ = (
        Index("query_date", "field_sid", "field_date"),
        Index("field_year_mon", "field_year", "field_mon"),
    )

    field_sid = Column(MEDIUMINT(8), primary_key=True)
    field_tid = Column(
        SMALLINT(6), nullable=False, index=True, server_default=text("'0'")
    )
    field_tags = Column(MEDIUMTEXT, nullable=False)
    field_rate_1 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_2 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_3 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_4 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_5 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_6 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_7 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_8 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_9 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_rate_10 = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    field_airtime = Column(TINYINT(1), nullable=False, index=True)
    field_rank = Column(
        INTEGER(10), nullable=False, index=True, server_default=text("'0'")
    )
    field_year = Column(YEAR(4), nullable=False, index=True, comment="放送年份")
    field_mon = Column(TINYINT(2), nullable=False, comment="放送月份")
    field_week_day = Column(TINYINT(1), nullable=False, comment="放送日(星期X)")
    # 对于默认的零值 '0000-00-00' 会被解析成字符串。
    # 非零值会被处理成 `datetime.date`
    field_date = Column(Date, nullable=False, index=True, comment="放送日期")
    field_redirect = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))

    subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiSubjectField.field_sid)
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
        back_populates="fields",
    )  # type: ignore


class ChiiSubjectRelations(Base):
    """
    这个表带有 comment，也没有主键，所以生成器用的是 `Table` 而不是现在的class。
    """

    __tablename__ = "chii_subject_relations"
    __table_args__ = (
        Index(
            "rlt_relation_type",
            "rlt_relation_type",
            "rlt_subject_id",
            "rlt_related_subject_id",
        ),
        Index(
            "rlt_subject_id",
            "rlt_subject_id",
            "rlt_related_subject_id",
            "rlt_vice_versa",
            unique=True,
        ),
        Index(
            "rlt_related_subject_type_id", "rlt_related_subject_type_id", "rlt_order"
        ),
    )
    rlt_subject_id = Column(
        "rlt_subject_id",
        MEDIUMINT(8),
        nullable=False,
        comment="关联主 ID",
    )
    rlt_subject_type_id = Column(
        "rlt_subject_type_id", TINYINT(3), nullable=False, index=True
    )
    rlt_relation_type: int = Column(
        "rlt_relation_type", SMALLINT(5), nullable=False, comment="关联类型"
    )
    rlt_related_subject_id = Column(
        "rlt_related_subject_id",
        MEDIUMINT(8),
        nullable=False,
        comment="关联目标 ID",
    )
    rlt_related_subject_type_id: int = Column(
        "rlt_related_subject_type_id", TINYINT(3), nullable=False, comment="关联目标类型"
    )
    rlt_vice_versa = Column("rlt_vice_versa", TINYINT(1), nullable=False)
    rlt_order = Column("rlt_order", TINYINT(3), nullable=False, comment="关联排序")

    __mapper_args__ = {
        "primary_key": [rlt_subject_id, rlt_related_subject_id, rlt_vice_versa]
    }

    src_subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=lambda: (
            ChiiSubject.subject_id == foreign(ChiiSubjectRelations.rlt_subject_id)
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
        back_populates="relations",
    )  # type: ignore

    dst_subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=lambda: (
            ChiiSubject.subject_id
            == foreign(ChiiSubjectRelations.rlt_related_subject_id)
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=False,
        back_populates="related",
    )  # type: ignore


class ChiiSubject(Base):
    __tablename__ = "chii_subjects"
    __table_args__ = (
        Index(
            "order_by_name",
            "subject_ban",
            "subject_type_id",
            "subject_series",
            "subject_platform",
            "subject_name",
        ),
        Index(
            "browser",
            "subject_ban",
            "subject_type_id",
            "subject_series",
            "subject_platform",
        ),
        Index("subject_idx_cn", "subject_idx_cn", "subject_type_id"),
    )

    subject_id = Column(MEDIUMINT(8), primary_key=True)
    subject_type_id = Column(
        SMALLINT(6), nullable=False, index=True, server_default=text("'0'")
    )
    subject_name = Column(String(80), nullable=False, index=True)
    subject_name_cn = Column(String(80), nullable=False, index=True)
    subject_uid = Column(String(20), nullable=False, comment="isbn / imdb")
    subject_creator = Column(MEDIUMINT(8), nullable=False, index=True)
    subject_dateline = Column(INTEGER(10), nullable=False, server_default=text("'0'"))
    subject_image = Column(String(255), nullable=False)
    subject_platform = Column(
        SMALLINT(6), nullable=False, index=True, server_default=text("'0'")
    )
    field_infobox = Column(MEDIUMTEXT, nullable=False)
    field_summary = Column(MEDIUMTEXT, nullable=False, comment="summary")
    field_5 = Column(MEDIUMTEXT, nullable=False, comment="author summary")
    field_volumes = Column(
        MEDIUMINT(8), nullable=False, server_default=text("'0'"), comment="卷数"
    )
    field_eps: int = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    subject_wish: int = Column(MEDIUMINT(8), nullable=False, server_default=text("'0'"))
    subject_collect: int = Column(
        MEDIUMINT(8), nullable=False, server_default=text("'0'")
    )
    subject_doing: int = Column(
        MEDIUMINT(8), nullable=False, server_default=text("'0'")
    )
    subject_on_hold: int = Column(
        MEDIUMINT(8), nullable=False, server_default=text("'0'"), comment="搁置人数"
    )
    subject_dropped: int = Column(
        MEDIUMINT(8), nullable=False, server_default=text("'0'"), comment="抛弃人数"
    )
    subject_series = Column(
        TINYINT(1), nullable=False, index=True, server_default=text("'0'")
    )
    subject_series_entry = Column(
        MEDIUMINT(8), nullable=False, index=True, server_default=text("'0'")
    )
    subject_idx_cn = Column(String(1), nullable=False)
    subject_airtime = Column(TINYINT(1), nullable=False, index=True)
    subject_nsfw = Column(TINYINT(1), nullable=False, index=True)
    subject_ban = Column(
        TINYINT(1), nullable=False, index=True, server_default=text("'0'")
    )

    persons: List[ChiiPersonCsIndex] = relationship(
        "ChiiPersonCsIndex",
        primaryjoin=(
            lambda: foreign(ChiiPersonCsIndex.subject_id) == ChiiSubject.subject_id
        ),
        lazy="raise_on_sql",
        back_populates="subject",
    )  # type: ignore

    characters: List[ChiiCrtSubjectIndex] = relationship(
        "ChiiCrtSubjectIndex",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiCrtSubjectIndex.subject_id)
        ),
        order_by=ChiiCrtSubjectIndex.crt_order,
        lazy="raise_on_sql",
        back_populates="subject",
    )  # type: ignore

    episodes: List[ChiiEpisode] = relationship(
        "ChiiEpisode",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiEpisode.ep_subject_id)
        ),
        lazy="raise_on_sql",
        order_by=(ChiiEpisode.ep_disc, ChiiEpisode.ep_type, ChiiEpisode.ep_sort),
        back_populates="subject",
    )  # type: ignore

    fields: ChiiSubjectField = relationship(
        "ChiiSubjectField",
        lazy="raise_on_sql",
        primaryjoin=(
            lambda: foreign(ChiiSubjectField.field_sid) == ChiiSubject.subject_id
        ),
        back_populates="subject",
        uselist=False,
    )  # type: ignore

    relations: List[ChiiSubjectRelations] = relationship(
        "ChiiSubjectRelations",
        primaryjoin=lambda: (
            ChiiSubject.subject_id == foreign(ChiiSubjectRelations.rlt_subject_id)
        ),
        order_by=ChiiSubjectRelations.rlt_order,
        lazy="raise_on_sql",
        back_populates="src_subject",
    )

    related: List[ChiiSubjectRelations] = relationship(
        "ChiiSubjectRelations",
        primaryjoin=lambda: (
            ChiiSubject.subject_id
            == foreign(ChiiSubjectRelations.rlt_related_subject_id)
        ),
        lazy="raise_on_sql",
        back_populates="dst_subject",
    )

    users: List["ChiiSubjectInterest"] = relationship(
        "ChiiSubjectInterest",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiSubjectInterest.subject_id)
        ),
        lazy="raise_on_sql",
        innerjoin=True,
        uselist=True,
        back_populates="subject",
    )  # type: ignore

    @property
    def locked(self) -> bool:
        return self.subject_ban == 2

    @property
    def ban(self) -> bool:
        return self.subject_ban == 1


class ChiiSubjectInterest(Base):
    __tablename__ = "chii_subject_interests"
    __table_args__ = (
        Index("user_collects", "interest_subject_type", "interest_uid"),
        Index(
            "tag_subject_id", "interest_subject_type", "interest_type", "interest_uid"
        ),
        Index(
            "subject_lasttouch",
            "interest_subject_id",
            "interest_private",
            "interest_lasttouch",
        ),
        Index(
            "user_collect_type",
            "interest_subject_type",
            "interest_type",
            "interest_uid",
            "interest_private",
            "interest_collect_dateline",
        ),
        Index(
            "subject_collect",
            "interest_subject_id",
            "interest_type",
            "interest_private",
            "interest_collect_dateline",
        ),
        Index(
            "subject_comment",
            "interest_subject_id",
            "interest_has_comment",
            "interest_private",
            "interest_lasttouch",
        ),
        Index("interest_id", "interest_uid", "interest_private"),
        Index(
            "user_collect_latest",
            "interest_subject_type",
            "interest_type",
            "interest_uid",
            "interest_private",
        ),
        Index(
            "top_subject",
            "interest_subject_id",
            "interest_subject_type",
            "interest_doing_dateline",
        ),
        Index(
            "subject_rate", "interest_subject_id", "interest_rate", "interest_private"
        ),
        Index("interest_type_2", "interest_type", "interest_uid"),
        Index(
            "interest_uid_2", "interest_uid", "interest_private", "interest_lasttouch"
        ),
        Index("user_interest", "interest_uid", "interest_subject_id", unique=True),
        Index("interest_subject_id", "interest_subject_id", "interest_type"),
    )

    id = Column("interest_id", INTEGER(10), primary_key=True)
    user_id = Column("interest_uid", MEDIUMINT(8), nullable=False, index=True)
    subject_id = Column("interest_subject_id", MEDIUMINT(8), nullable=False, index=True)
    subject_type = Column(
        "interest_subject_type",
        SMALLINT(6),
        nullable=False,
        index=True,
        server_default=text("'0'"),
    )
    rate = Column(
        "interest_rate",
        TINYINT(3),
        nullable=False,
        index=True,
        server_default=text("'0'"),
    )
    type = Column(
        "interest_type",
        TINYINT(1),
        nullable=False,
        index=True,
        server_default=text("'0'"),
    )
    has_comment = Column("interest_has_comment", TINYINT(1), nullable=False, default=0)
    comment = Column("interest_comment", MEDIUMTEXT, nullable=False, default="")
    tag: str = Column("interest_tag", MEDIUMTEXT, nullable=False, default="")
    ep_status = Column(
        "interest_ep_status",
        MEDIUMINT(8),
        nullable=False,
        server_default=text("'0'"),
    )
    vol_status = Column(
        "interest_vol_status",
        MEDIUMINT(8),
        nullable=False,
        comment="卷数",
        default=0,
    )
    wish_dateline = Column(
        "interest_wish_dateline", INTEGER(10), nullable=False, default=0
    )
    doing_dateline = Column(
        "interest_doing_dateline", INTEGER(10), nullable=False, default=0
    )
    collect_dateline = Column(
        "interest_collect_dateline", INTEGER(10), nullable=False, index=True, default=0
    )
    on_hold_dateline = Column(
        "interest_on_hold_dateline", INTEGER(10), nullable=False, default=0
    )
    dropped_dateline = Column(
        "interest_dropped_dateline", INTEGER(10), nullable=False, default=0
    )
    last_touch = Column(
        "interest_lasttouch",
        INTEGER(10),
        nullable=False,
        index=True,
        server_default=text("'0'"),
    )
    private = Column(
        "interest_private", TINYINT(1), nullable=False, index=True, default=0
    )

    subject: "ChiiSubject" = relationship(
        "ChiiSubject",
        primaryjoin=(
            lambda: ChiiSubject.subject_id == foreign(ChiiSubjectInterest.subject_id)
        ),
        lazy="raise_on_sql",
        uselist=False,
        back_populates="users",
    )  # type: ignore
