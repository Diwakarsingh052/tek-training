package models

func (s *Service) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}

	err := s.db.Migrator().DropTable(&User{}, &Inventory{})
	if err != nil {
		return err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = s.db.Migrator().AutoMigrate(&User{}, &Inventory{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
