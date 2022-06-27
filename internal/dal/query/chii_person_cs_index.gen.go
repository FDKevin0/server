// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newPersonSubjects(db *gorm.DB) personSubjects {
	_personSubjects := personSubjects{}

	_personSubjects.personSubjectsDo.UseDB(db)
	_personSubjects.personSubjectsDo.UseModel(&dao.PersonSubjects{})

	tableName := _personSubjects.personSubjectsDo.TableName()
	_personSubjects.ALL = field.NewField(tableName, "*")
	_personSubjects.PrsnType = field.NewString(tableName, "prsn_type")
	_personSubjects.PersonID = field.NewField(tableName, "prsn_id")
	_personSubjects.PrsnPosition = field.NewUint16(tableName, "prsn_position")
	_personSubjects.SubjectID = field.NewField(tableName, "subject_id")
	_personSubjects.SubjectTypeID = field.NewUint8(tableName, "subject_type_id")
	_personSubjects.Summary = field.NewString(tableName, "summary")
	_personSubjects.PrsnAppearEps = field.NewString(tableName, "prsn_appear_eps")
	_personSubjects.Subject = personSubjectsHasOneSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_personSubjects.Person = personSubjectsHasOnePerson{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Person", "dao.Person"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Person.Fields", "dao.PersonField"),
		},
	}

	_personSubjects.fillFieldMap()

	return _personSubjects
}

type personSubjects struct {
	personSubjectsDo personSubjectsDo

	ALL           field.Field
	PrsnType      field.String
	PersonID      field.Field
	PrsnPosition  field.Uint16
	SubjectID     field.Field
	SubjectTypeID field.Uint8
	Summary       field.String
	PrsnAppearEps field.String
	Subject       personSubjectsHasOneSubject

	Person personSubjectsHasOnePerson

	fieldMap map[string]field.Expr
}

func (p personSubjects) Table(newTableName string) *personSubjects {
	p.personSubjectsDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personSubjects) As(alias string) *personSubjects {
	p.personSubjectsDo.DO = *(p.personSubjectsDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personSubjects) updateTableName(table string) *personSubjects {
	p.ALL = field.NewField(table, "*")
	p.PrsnType = field.NewString(table, "prsn_type")
	p.PersonID = field.NewField(table, "prsn_id")
	p.PrsnPosition = field.NewUint16(table, "prsn_position")
	p.SubjectID = field.NewField(table, "subject_id")
	p.SubjectTypeID = field.NewUint8(table, "subject_type_id")
	p.Summary = field.NewString(table, "summary")
	p.PrsnAppearEps = field.NewString(table, "prsn_appear_eps")

	p.fillFieldMap()

	return p
}

func (p *personSubjects) WithContext(ctx context.Context) *personSubjectsDo {
	return p.personSubjectsDo.WithContext(ctx)
}

func (p personSubjects) TableName() string { return p.personSubjectsDo.TableName() }

func (p personSubjects) Alias() string { return p.personSubjectsDo.Alias() }

func (p *personSubjects) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personSubjects) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["prsn_type"] = p.PrsnType
	p.fieldMap["prsn_id"] = p.PersonID
	p.fieldMap["prsn_position"] = p.PrsnPosition
	p.fieldMap["subject_id"] = p.SubjectID
	p.fieldMap["subject_type_id"] = p.SubjectTypeID
	p.fieldMap["summary"] = p.Summary
	p.fieldMap["prsn_appear_eps"] = p.PrsnAppearEps

}

func (p personSubjects) clone(db *gorm.DB) personSubjects {
	p.personSubjectsDo.ReplaceDB(db)
	return p
}

type personSubjectsHasOneSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a personSubjectsHasOneSubject) Where(conds ...field.Expr) *personSubjectsHasOneSubject {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a personSubjectsHasOneSubject) WithContext(ctx context.Context) *personSubjectsHasOneSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a personSubjectsHasOneSubject) Model(m *dao.PersonSubjects) *personSubjectsHasOneSubjectTx {
	return &personSubjectsHasOneSubjectTx{a.db.Model(m).Association(a.Name())}
}

type personSubjectsHasOneSubjectTx struct{ tx *gorm.Association }

func (a personSubjectsHasOneSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a personSubjectsHasOneSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a personSubjectsHasOneSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a personSubjectsHasOneSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a personSubjectsHasOneSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a personSubjectsHasOneSubjectTx) Count() int64 {
	return a.tx.Count()
}

type personSubjectsHasOnePerson struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a personSubjectsHasOnePerson) Where(conds ...field.Expr) *personSubjectsHasOnePerson {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a personSubjectsHasOnePerson) WithContext(ctx context.Context) *personSubjectsHasOnePerson {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a personSubjectsHasOnePerson) Model(m *dao.PersonSubjects) *personSubjectsHasOnePersonTx {
	return &personSubjectsHasOnePersonTx{a.db.Model(m).Association(a.Name())}
}

type personSubjectsHasOnePersonTx struct{ tx *gorm.Association }

func (a personSubjectsHasOnePersonTx) Find() (result *dao.Person, err error) {
	return result, a.tx.Find(&result)
}

func (a personSubjectsHasOnePersonTx) Append(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a personSubjectsHasOnePersonTx) Replace(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a personSubjectsHasOnePersonTx) Delete(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a personSubjectsHasOnePersonTx) Clear() error {
	return a.tx.Clear()
}

func (a personSubjectsHasOnePersonTx) Count() int64 {
	return a.tx.Count()
}

type personSubjectsDo struct{ gen.DO }

func (p personSubjectsDo) Debug() *personSubjectsDo {
	return p.withDO(p.DO.Debug())
}

func (p personSubjectsDo) WithContext(ctx context.Context) *personSubjectsDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personSubjectsDo) ReadDB(ctx context.Context) *personSubjectsDo {
	return p.WithContext(ctx).Clauses(dbresolver.Read)
}

func (p personSubjectsDo) WriteDB(ctx context.Context) *personSubjectsDo {
	return p.WithContext(ctx).Clauses(dbresolver.Write)
}

func (p personSubjectsDo) Clauses(conds ...clause.Expression) *personSubjectsDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personSubjectsDo) Returning(value interface{}, columns ...string) *personSubjectsDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personSubjectsDo) Not(conds ...gen.Condition) *personSubjectsDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personSubjectsDo) Or(conds ...gen.Condition) *personSubjectsDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personSubjectsDo) Select(conds ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personSubjectsDo) Where(conds ...gen.Condition) *personSubjectsDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personSubjectsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *personSubjectsDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p personSubjectsDo) Order(conds ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personSubjectsDo) Distinct(cols ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personSubjectsDo) Omit(cols ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personSubjectsDo) Join(table schema.Tabler, on ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personSubjectsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personSubjectsDo) RightJoin(table schema.Tabler, on ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personSubjectsDo) Group(cols ...field.Expr) *personSubjectsDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personSubjectsDo) Having(conds ...gen.Condition) *personSubjectsDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personSubjectsDo) Limit(limit int) *personSubjectsDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personSubjectsDo) Offset(offset int) *personSubjectsDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personSubjectsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personSubjectsDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personSubjectsDo) Unscoped() *personSubjectsDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personSubjectsDo) Create(values ...*dao.PersonSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personSubjectsDo) CreateInBatches(values []*dao.PersonSubjects, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personSubjectsDo) Save(values ...*dao.PersonSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personSubjectsDo) First() (*dao.PersonSubjects, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonSubjects), nil
	}
}

func (p personSubjectsDo) Take() (*dao.PersonSubjects, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonSubjects), nil
	}
}

func (p personSubjectsDo) Last() (*dao.PersonSubjects, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonSubjects), nil
	}
}

func (p personSubjectsDo) Find() ([]*dao.PersonSubjects, error) {
	result, err := p.DO.Find()
	return result.([]*dao.PersonSubjects), err
}

func (p personSubjectsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.PersonSubjects, err error) {
	buf := make([]*dao.PersonSubjects, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personSubjectsDo) FindInBatches(result *[]*dao.PersonSubjects, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personSubjectsDo) Attrs(attrs ...field.AssignExpr) *personSubjectsDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personSubjectsDo) Assign(attrs ...field.AssignExpr) *personSubjectsDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personSubjectsDo) Joins(fields ...field.RelationField) *personSubjectsDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personSubjectsDo) Preload(fields ...field.RelationField) *personSubjectsDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personSubjectsDo) FirstOrInit() (*dao.PersonSubjects, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonSubjects), nil
	}
}

func (p personSubjectsDo) FirstOrCreate() (*dao.PersonSubjects, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonSubjects), nil
	}
}

func (p personSubjectsDo) FindByPage(offset int, limit int) (result []*dao.PersonSubjects, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p personSubjectsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p *personSubjectsDo) withDO(do gen.Dao) *personSubjectsDo {
	p.DO = *do.(*gen.DO)
	return p
}
