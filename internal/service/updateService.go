package service

func (u *UserService) UpdateService(id int, newUsername string) error {
	err := u.Repo.Update(id, newUsername)
	return err
}
