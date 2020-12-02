func (s *usersService) GetUser(userId int64) (*users.User, user) {
	dao := &users.User{Id: userId}
	if err := dao.Get(userId); err != nil {
		return nil, err
	}
	return dao, nil
}