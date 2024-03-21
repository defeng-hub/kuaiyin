package ftpsync

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

type FTPClient struct {
	Host     string
	Username string
	Password string
	client   *ftp.ServerConn
}

var FtpClientConn *FTPClient

func NewFTPClient(host string, username string, password string) *FTPClient {
	return &FTPClient{
		Host:     host,
		Username: username,
		Password: password,
	}
}

func (c *FTPClient) Connect() error {
	var err error
	//c.client, err = ftp.Dial(fmt.Sprintf("%s:%d", c.Host, c.Port))
	c.client, err = ftp.Dial(c.Host)
	if err != nil {
		return err
	}

	err = c.client.Login(c.Username, c.Password)
	if err != nil {
		c.Disconnect()
		return err
	}

	err = c.client.Type(ftp.TransferTypeBinary)
	if err != nil {
		return err
	}
	return nil
}

func (c *FTPClient) Disconnect() {
	if c.client != nil {
		_ = c.client.Quit()
	}
}

func (c *FTPClient) Reconnect() error {
	c.Disconnect()
	return c.Connect()
}

func (c *FTPClient) Upload(localPath string, remotePath string) error {
	if c.client == nil {
		err := c.Connect()
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		if err := c.client.NoOp(); err != nil {
			fmt.Println("FTP client is disconnected.try to reconnected")
			err := c.Connect()
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			fmt.Println("FTP client is connected. go on uploading")
		}
	}
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = c.client.Stor(remotePath, file)
	if err != nil {
		return err
	}

	return nil
}
