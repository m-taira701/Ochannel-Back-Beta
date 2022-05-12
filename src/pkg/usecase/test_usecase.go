package usecase

// Resolver動作デモ用UseCase
type TestUsecase struct{}

func (usecase *TestUsecase) Test() string {
	return "hello world";
}