package accountsvc

type accountService struct{}

func NewAccountService() *accountService {
	return &accountService{}
}
