package interfaces

type GqBuilder interface {
	Select(columns ...string) GqBuilder
	Insert(tableName string, columns ...string) GqBuilder
	Values(values string, args ...any) GqBuilder
	From(tableName string) GqBuilder
	Where(conditions string, args ...any) GqBuilder
	Build() (string, any, error)
}
