GO validation package (alpha version)
==================================================

#### The main features of this package:
+ Simple and flexible API
+ Many date validators with custom intervals (+1 year, - 1 month, etc)
+ Can validate custom types without any special code
+ Implemented all common validators
+ Custom validators support
+ Error messages API with custom messages with params
+ Options - "required", "lazy", etc. With custom options support
+ Actions - "trim", "lower", etc. With custom actions support

#### Scheduled features:
- Compare validator
- Embedded parameters
- Options "required_with", "required_unless"
- Embedded structs support
- Folding structs support
- JSON tag support
- Functions for testing validation rules

#### Quick examples:

You can use structure tags to specify validators...
```go
type User struct {
	Name           string `valid:"required|min:2|max:100|custom_validator"`
	Email          string `valid:"required|email"            on_update:"ignore"`
	Password       string `valid:"password"                  on_create:"required"`
	PasswordRepeat string `valid:"password|compare:Password" on_create:"required"`
	Birthday       string `valid:"date_gte:02-01-2006,-18Y"` // age must be 18 years +
}

notValidUser := User{}
errors := validation.ValidateStruct(notValidUser, "valid", "on_create")

// Check if has errors
if errors.Empty() {
    // do somthing
} else {
    fmt.Println(errors.JSON())
}

// JSON representation of struct's errors
{
    "Name":[
        "must be greater or equal of 2",
        "message from custom validator"
    ],
    "Email":[
        "must be a valid email address"
    ],
    "Password":[
        "must contains at least english letters in both cases, numbers and have minimum length 8"
    ],
    "Birthday":[
        "must be greater or equal of now -18Y"
    ]
}
```

Or you can validate structure in functional way...

```
type User struct {
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
	Birthday       string
}

func (u User) isValid() validation.ErrorMap {
	return validation.ErrorMap{
	        "Name": validation.ValidateValue(u.Name, is.Required(), is.Min(2), is.Max(100), CustomValidator),
		"Email": validation.ValidateValue(u.Email, is.Required(), is.Email()),
		"Password": validation.ValidateValue(u.Password, is.Password()),
                "PasswordRepeat": validation.ValidateValue(u.PasswordRepeat, is.Compare("Password")),
                "Birthday": validation.ValidateValue(u.Birthday, is.DateGte("02-01-2006", "-18Y")),
	}
}
```
