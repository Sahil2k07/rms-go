package dto

type CreateJobDto struct {
	Title       string `json:"title" validate:"required,excludesall=;=:"`
	Description string `json:"description" validate:"required,excludesall=;=:"`
	CompanyName string `json:"companyName" validate:"required,excludesall=(;)=:"`
}
