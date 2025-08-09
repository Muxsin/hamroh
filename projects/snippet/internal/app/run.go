package app

func (a *app) Run() error {
	if err := a.router.Run(":" + a.configs.HTTPServerPort); err != nil {
		return err
	}

	return nil
}
