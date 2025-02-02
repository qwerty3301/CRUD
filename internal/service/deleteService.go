package service

func (u *UserService) DeleteService(id int) error {
	err := u.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
