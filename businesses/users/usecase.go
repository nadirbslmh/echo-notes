package users

type UserUsecase struct {
	userRepository Repository
}

func NewUserUsecase(ur Repository) Usecase {
	return &UserUsecase{
		userRepository: ur,
	}
}

func (uu *UserUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUsecase) Login(userDomain *Domain) string {
	return uu.userRepository.Login(userDomain)
}
