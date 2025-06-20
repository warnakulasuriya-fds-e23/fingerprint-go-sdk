package core

type customTransparencyContents struct {
}

func (c *customTransparencyContents) Accepts(key string) bool {
	return true
}

func (c *customTransparencyContents) Accept(key, mime string, data []byte) error {
	//fmt.Printf("%d B  %s %s \n", len(data), mime, key)
	return nil
}
