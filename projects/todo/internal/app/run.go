package app

func (a *App) Run() error {
	if err := a.router.Run(":" + a.configs.HTTPServerPort); err != nil {
		return err
	}

	return nil
}
