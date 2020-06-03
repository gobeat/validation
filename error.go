package validation

const (
	// Common
	True  = "true"
	False = "false"
	Empty = ""

	// Validator errors
	invalidTypeError          = "%s type is not supported"
	invalidArgumentError      = "invalid argument found"
	notIntegerValueError      = "value is not an integer number"
	notFloatValueError        = "value is not a float number"
	notNumberValueError       = "value is not a number"
	notStringNumberValueError = "value (string) is not a number"
	notNilValueError          = "value is not nil"
	notStringValueError       = "value is not a string"

	// Checkers errors
	invalidEmailFormatError    = "value is not an email address"
	emptyListFoundError        = "%v is not expected to appear in an empty list"
	itemNotFoundInListError    = "%v is not appeared in %v"
	maxValueError              = "maximum value is %v"
	maxLengthValueError        = "maximum length is %v"
	lengthValueError           = "length must be %v"
	minValueError              = "minimum value is %v"
	minLengthValueError        = "minimum length is %v"
	invalidRangeFormatError    = "invalid range format"
	notInRangeError            = "value must be between %v and %v"
	invalidRegexPatternError   = "%v is not a valid pattern"
	regexValueError            = "value does not match %v"
	unableToFetchResourceError = "unable to fetch appropriate resource"
	notUniqueValueError        = "value already exists. It must be unique"
	emptyValueError            = "value must not be empty"
	nilValueError              = "value must not be nil"
)
