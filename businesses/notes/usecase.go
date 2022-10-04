package notes

type noteUsecase struct {
	noteRepository Repository
}

func NewNoteUsecase(nr Repository) Usecase {
	return &noteUsecase{
		noteRepository: nr,
	}
}

func (nu *noteUsecase) GetAll() []Domain {
	return nu.noteRepository.GetAll()
}
func (nu *noteUsecase) GetByID(id string) Domain {
	return nu.noteRepository.GetByID(id)
}
func (nu *noteUsecase) Create(noteDomain *Domain) Domain {
	return nu.noteRepository.Create(noteDomain)
}
func (nu *noteUsecase) Update(id string, noteDomain *Domain) Domain {
	return nu.noteRepository.Update(id, noteDomain)
}
func (nu *noteUsecase) Delete(id string) bool {
	return nu.noteRepository.Delete(id)
}
func (nu *noteUsecase) Restore(id string) Domain {
	return nu.noteRepository.Restore(id)
}
func (nu *noteUsecase) ForceDelete(id string) bool {
	return nu.noteRepository.ForceDelete(id)
}
