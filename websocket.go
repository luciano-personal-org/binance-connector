package binance_connector

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

type WebsocketStreamClient struct {
	Endpoint   string
	IsCombined bool
}

func NewWebsocketStreamClient(isCombined bool, baseURL ...string) *WebsocketStreamClient {
	// Set default base URL to production WS URL
	url := "wss://stream.binance.com:9443"

	if len(baseURL) > 0 {
		url = baseURL[0]
	}

	// Append to baseURL based on whether the client is for combined streams or not
	if isCombined {
		url += "/stream?streams="
	} else {
		url += "/ws"
	}

	return &WebsocketStreamClient{
		Endpoint:   url,
		IsCombined: isCombined,
	}
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  24 * time.Hour, // 24 hours connected, it is the maximum time allowed by the Binance server
		EnableCompression: false,
		// HandshakeTimeout:  time.Duration(10) * time.Second,
		// TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	c, httpResponse, err := Dialer.Dial(cfg.Endpoint, headers)
	if err != nil {
		fmt.Printf("Connecting to: %s\n", cfg.Endpoint)
		fmt.Printf("HTTP Response Status: %s\n", httpResponse.Status)
		fmt.Printf("HTTP Response Body: %s\n", httpResponse.Body)
		fmt.Printf("HTTP Response TLS NegotiatedProtocol: %s\n", httpResponse.TLS.NegotiatedProtocol)
		switch err.(type) {
		case *websocket.CloseError:
			err = fmt.Errorf("websocket.CloseError: %v", err)
		case *websocket.HandshakeError:
			err = fmt.Errorf("websocket.Handshake: %v", err)
		}
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneCh)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					if !silent {
						errHandler(err)
					}
					stopCh <- struct{}{}
					return
				}
				handler(message)
			}
		}()

		for {
			select {
			case <-stopCh:
				silent = true
				return
			case <-doneCh:
			}
		}
	}()
	return

}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				return
			}
		}
	}()
}
