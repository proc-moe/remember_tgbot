package dal

type Client struct {
	tg *TGProxy
}

var Cli *Client

func InitDal() {
	Cli = &Client{
		tg: InitTGProxy(),
	}
}
