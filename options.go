package validation

// If this option is present value must be not empty
const Required Option = "required"

// If this option is present validation will not be performed for this field
const Ignore Option = "ignore"

// If this option is present validation will be performed before the first error
const Lazy Option = "lazy"

// Validation option
type Option string

// List of present options
type OptionList []Option

// List of possible options
var Options = OptionList{
	Required,
	Ignore,
	Lazy,
}

// Check what option exists
func (o OptionList) Has(name Option) bool {
	for _, option := range o {
		if name == option {
			return true
		}
	}
	return false
}

// Add option to list
func (o *OptionList) Add(option Option) {
	*o = append(*o, option)
}
