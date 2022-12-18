package dal

type Client struct {
	Tg *TGProxy
}

var Cli *Client

func Init() {
	Cli = &Client{
		Tg: InitTGProxy(),
	}
}
