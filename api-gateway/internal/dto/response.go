package dto


// APIResponse универсальная структура ответа для ошибок/сообщений
type APIResponse struct {
    Error string `json:"error" example:"something went wrong"`
}