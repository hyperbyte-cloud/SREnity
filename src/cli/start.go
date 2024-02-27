package cli

// CLI Start is to Start the cli in server mode
func (c *CLI) Start(file string) error {
	// Start the server
	err := c.Domain.Start(file)
	if err != nil {
		return err
	}

	// wait forever
	select {}
}
