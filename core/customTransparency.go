package core

type CustomTransparencyContents struct {
}

func (c *CustomTransparencyContents) Accepts(key string) bool {
	return true
}

func (c *CustomTransparencyContents) Accept(key, mime string, data []byte) error {
	//fmt.Printf("%d B  %s %s \n", len(data), mime, key)
	return nil
}
