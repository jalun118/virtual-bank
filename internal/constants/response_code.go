package constants

const (
	DuplicateRecord = iota + 1000
	RecordNotFound
	RecordFound

	InputError

	// user
	UserNotFound
	PasswordNotMatch

	// account type
	AccountTypeNotFound

	// account
	PinNotMatch
	AccountNotFound
	DestinationAccountNotFound
	NotEnoughBalance

	// validation
	FieldRequired
	FieldWrongValue
	FieldBoolean
	FieldNumber
	FieldEmail
	FieldAlphaNumber
	FieldAlphaNumberSpecial
	FieldLen
	MinimumNumber
	MaximumNumber
	FieldAscii
	FieldNotAllowChar
	NotCompleteForm
	UnsupportedContentType
	ColumnNull
	ReachingTheLimit
	AccessUnauthorized

	// database
	SuccessInsert
	SuccessUpdate
	SuccessRead
	SuccessDelete

	InternalServerError
	DatabaseError
)

var ValidatorMsg = map[string]int{
	"required":        FieldRequired,
	"boolean":         FieldBoolean,
	"email":           FieldEmail,
	"number":          FieldNumber,
	"min":             MinimumNumber,
	"max":             MaximumNumber,
	"len":             FieldLen,
	"alphanum":        FieldAlphaNumber,
	"alphanumspecial": FieldAlphaNumberSpecial,
	"ascii":           FieldAscii,
	"notallowchar":    FieldNotAllowChar,
}
