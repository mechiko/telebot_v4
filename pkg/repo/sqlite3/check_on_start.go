package sqlite3

func (r *repository) CheckOnStart() error {
	defer r.App.GetRecovery().RecoverLog("TestOnStart()")
	// dbService := r.App.GetDbService()
	// // создаем структуру БД если файл БД был только что создан
	// if !dbService.IsCreated() {
	// 	if err := r.create(); err != nil {
	// 		return fmt.Errorf("r.create() %w", err)
	// 	}
	// 	dbService.SetCreated()
	// }

	// all_names, _ := r.GetTaskNames().Get()
	// r.App.DebugLog().Str("test_on_start", "all_names").Msgf("%v", all_names)
	// names := r.GetTaskNames().GetNames()
	// r.App.DebugLog().Str("test_on_start", "names").Msgf("%v", names)
	// name, _ := r.GetTaskNames().GetByName("Запрос справок 1 по остаткам")
	// r.App.DebugLog().Str("test_on_start", "name").Msgf("%v", name)
	// r.App.Shutdown()
	return nil
}
