package languages

type ClientInterface interface {
	List() []*Language
	Get(slug string) *Language
}

type Client struct {
}

func NewClient() ClientInterface {
	return &Client{}
}

func (c *Client) List() []*Language {
	return []*Language{
		&Language{"Go", "go"},
		&Language{"JavaScript", "javascript"},
		&Language{"Ruby", "ruby"},
		&Language{"Python", "python"},
	}
}

func (c *Client) Get(slug string) *Language {
	return &Language{"Go", "go"}
}
