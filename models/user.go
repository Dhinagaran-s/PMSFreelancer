package models

// RegistrationUserUsingPassword => User Model for registration via email and password
type RegistrationUserUsingPassword struct {
	Email       string
	Password    string
	Fullname    string
	Location    int
	Description string
	Skills      string
	Username    string
}

type RegistrationUserUsingGoogle struct {
	Email       string
	GoogleID    string
	Fullname    string
	Location    int
	Description string
	Skills      string
	Username    string
}

type LoginUserPassword struct {
	Email    string
	Password string
}

type GoogleResponse struct {
	ID             string
	Email          string
	Verified_email bool
	Picture        string
}

type RegistrationInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UpdateResetPassword struct {
	Password string
	Token    string
}

type DatabaseResetPassword struct {
	Email  string
	Token  string
	Expire string
}

type UserProfile struct {
	ID          string                  `json:"id"`
	Fullname    string                  `json:"fullname"`
	Email       string                  `json:"email"`
	Description string                  `json:"description"`
	Education   []EducationReturnValue  `json:"education"`
	Skill       []UserSkills            `json:"skill"`
	Experience  []ExperienceReturnValue `json:"experience"`
	Picture     string                  `json:"picture"`
	Username    string                  `json:"username"`
	Location    string                  `json:"location"`
	Member      string                  `json:"member"`
	Portfolio   []PortfolioDatabase     `json:"portfolio"`
}

type QueryUserProfile struct {
	ID          string
	Firstname   string
	LastName    string
	Email       string
	Description string
	Picture     string
	CreatedAt   string
	Username    string
	Location    string
	Skills      string
}
