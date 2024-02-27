package domain

// Start the domain
func (d *Domain) Start(file string) error {
	// Load the configuration
	config, err := d.LoadConfig(file)
	if err != nil {
		return err
	}

	// Start the server
	err = d.Server(config)
	if err != nil {
		return err
	}

	return nil
}
