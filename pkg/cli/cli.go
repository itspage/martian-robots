package cli

type CLI struct{}

func (c *CLI) ReadLine(line string) error {
	return nil
}

func (c *CLI) Output() ([]string, error) {
	return nil, nil
}
