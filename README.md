# Binance Spot Go Connector

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

This is a lightweight, high-performance Go library that works as a connector to the [Binance public API](https://github.com/binance/binance-spot-api-docs). It provides a clean, idiomatic Go interface for interacting with Binance's spot trading, market data, and account management APIs.

## ✨ Features

- **Complete API Coverage**: Support for all major Binance Spot API endpoints
- **WebSocket Support**: Real-time market data and user data streams
- **WebSocket API**: Real-time trading operations via WebSocket
- **Type Safety**: Strongly typed request/response structures
- **Flexible Authentication**: Support for API key and signed requests
- **Error Handling**: Comprehensive error handling with detailed API error responses
- **Rate Limiting**: Built-in respect for Binance rate limits
- **Testnet Support**: Easy switching between production and testnet environments
- **Debug Mode**: Built-in logging for debugging and monitoring
- **Context Support**: Full context.Context support for request cancellation and timeouts

## 🚀 Quick Start

### Installation

```shell
go get github.com/luciano-personal-org/binance-connector
```

### Import

```golang
import (
    binance_connector "github.com/luciano-personal-org/binance-connector"
)
```

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
    // Initialize client
    client := binance_connector.NewClient("your-api-key", "your-secret-key")
    
    // Get account information
    account, err := client.NewGetAccountService().Do(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(binance_connector.PrettyPrint(account))
}
```

## 📚 API Coverage

### Account & Trading
- **Order Management**: Create, cancel, modify orders (including OCO orders)
- **Account Information**: Get account details, balances, trading status
- **Trade History**: Query order history, trade history, and prevented matches
- **Order Types**: Market, Limit, Stop-Loss, Take-Profit, OCO orders

### Market Data  
- **Real-time Data**: Ticker prices, order book depth, recent trades
- **Historical Data**: Kline/candlestick data, historical trades
- **Statistics**: 24hr ticker statistics, price change statistics
- **Exchange Info**: Trading rules, symbol information, filters

### Wallet Operations
- **Asset Management**: Deposit/withdrawal history, asset details
- **Transfers**: Internal transfers, universal transfers
- **Dust Conversion**: Convert small balances to BNB
- **API Management**: API key permissions, trading status

### Margin Trading
- **Margin Orders**: Place and manage margin orders
- **Account Management**: Margin account details, isolated margin
- **Borrowing/Lending**: Query max borrow, repay history
- **Risk Management**: Liquidation records, margin ratios

### Sub-Account Management
- **Account Creation**: Create and manage sub-accounts
- **Asset Management**: Transfer assets between sub-accounts  
- **Futures & Margin**: Enable futures/margin for sub-accounts
- **Monitoring**: Query sub-account status and transaction statistics

### FIAT Operations
- **Payment History**: Query fiat payment transactions
- **Deposit/Withdrawal**: FIAT deposit and withdrawal history
## 🔐 Authentication

### Basic Client Setup
```go
// Initialize with API credentials
client := binance_connector.NewClient("your-api-key", "your-secret-key")

// Or specify custom base URL (optional)
client := binance_connector.NewClient("your-api-key", "your-secret-key", "https://api.binance.com")
```

### API Key Requirements

Different endpoints require different authentication levels:

- **Public endpoints** (market data): No authentication required
- **API-KEY required**: Requires valid API key in request header
- **SIGNED required**: Requires API key + signature for private operations

### Client Configuration

```go
client := binance_connector.NewClient("api-key", "secret-key", "base-url")

// Enable debug logging
client.Debug = true

// Adjust request timestamp (useful for server time sync issues)
client.TimeOffset = -1000 // milliseconds adjustment

// Custom HTTP client with timeout
client.HTTPClient = &http.Client{
    Timeout: 10 * time.Second,
}
```

## 📈 REST API Examples

### Trading Operations

#### Create Market Order
```go
package main

import (
    "context"
    "fmt"
    "log"

    binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
    // Initialize client with testnet
    client := binance_connector.NewClient("your-api-key", "your-secret-key", "https://testnet.binance.vision")

    // Create market buy order
    order, err := client.NewCreateOrderService().
        Symbol("BTCUSDT").
        Side("BUY").
        Type("MARKET").
        Quantity(0.001).
        Do(context.Background())
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Order created:", binance_connector.PrettyPrint(order))
}
```

#### Create Limit Order
```go
// Create limit sell order
order, err := client.NewCreateOrderService().
    Symbol("BTCUSDT").
    Side("SELL").
    Type("LIMIT").
    TimeInForce("GTC").
    Quantity(0.001).
    Price(45000.00).
    Do(context.Background())
```

#### Cancel Order
```go
// Cancel order by order ID
result, err := client.NewCancelOrderService().
    Symbol("BTCUSDT").
    OrderId(12345).
    Do(context.Background())
```

#### OCO (One-Cancels-Other) Orders
```go
// Create OCO order
ocoOrder, err := client.NewNewOCOService().
    Symbol("BTCUSDT").
    Side("SELL").
    Quantity(1.0).
    Price(50000.00).        // Limit price
    StopPrice(45000.00).    // Stop price
    StopLimitPrice(44800.00). // Stop limit price
    Do(context.Background())
```

### Market Data

#### Get Order Book
```go
// Get order book depth
orderBook, err := client.NewOrderBookService().
    Symbol("BTCUSDT").
    Limit(100).
    Do(context.Background())
```

#### Get Kline Data
```go
// Get candlestick data
klines, err := client.NewKlinesService().
    Symbol("BTCUSDT").
    Interval("1h").
    StartTime(1640995200000). // Unix timestamp in milliseconds
    EndTime(1641081600000).
    Limit(500).
    Do(context.Background())
```

#### Get 24hr Ticker Statistics
```go
// Get 24hr ticker for single symbol
ticker, err := client.NewTicker24hrService().
    Symbol("BTCUSDT").
    Do(context.Background())

// Get 24hr ticker for all symbols
allTickers, err := client.NewTicker24hrService().
    Do(context.Background())
```

### Account Information

#### Get Account Details
```go
// Get account information
account, err := client.NewGetAccountService().
    Do(context.Background())
```

#### Get Trade History
```go
// Get account trade list
trades, err := client.NewGetMyTradesService().
    Symbol("BTCUSDT").
    StartTime(1640995200000).
    EndTime(1641081600000).
    Limit(500).
    Do(context.Background())
```

### Request Options

You can use request options to customize individual requests:

```go
// Use custom receive window
order, err := client.NewCreateOrderService().
    Symbol("BTCUSDT").
    Side("BUY").
    Type("MARKET").
    Quantity(0.001).
    Do(context.Background(), binance_connector.WithRecvWindow(10000))
```

### Error Handling

```go
order, err := client.NewCreateOrderService().
    Symbol("BTCUSDT").
    Side("BUY").
    Type("MARKET").
    Quantity(0.001).
    Do(context.Background())

if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*handlers.APIError); ok {
        fmt.Printf("API Error: Code=%d, Message=%s\n", apiErr.Code, apiErr.Message)
    } else {
        fmt.Printf("Network/Other Error: %v\n", err)
    }
    return
}
```

### 📁 Examples Directory

Comprehensive examples for all endpoints can be found in the `examples/` directory:

- `examples/account/` - Trading and account operations
- `examples/market/` - Market data endpoints  
- `examples/wallet/` - Wallet and asset operations
- `examples/margin/` - Margin trading
- `examples/subaccount/` - Sub-account management
- `examples/websocket/` - WebSocket stream examples
- `examples/websocket_api/` - WebSocket API examples

## 🌐 WebSocket Streams

WebSocket streams provide real-time market data and user account updates with low latency.

### Stream Client Initialization

```go
// Single stream client (for individual symbol streams)
wsClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.binance.com:9443")

// Combined stream client (for multiple streams in one connection)
wsClient := binance_connector.NewWebsocketStreamClient(true, "wss://stream.binance.com:9443")

// Testnet
wsClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
```

### Parameters
- **`isCombined`** (boolean): 
  - `false`: Individual stream mode (`/ws/` endpoint)
  - `true`: Combined stream mode (`/stream?streams=` endpoint)
- **`baseURL`** (optional string): WebSocket base URL
  - Defaults to `"wss://stream.binance.com:9443"` if not specified

### Market Data Streams

#### Order Book Depth Stream
```go
package main

import (
    "fmt"
    "log"
    "time"

    binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
    // Initialize WebSocket client
    wsClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")

    // Define event handler
    depthHandler := func(event *binance_connector.WsDepthEvent) {
        fmt.Printf("Symbol: %s, Bids: %d, Asks: %d\n", 
            event.Symbol, len(event.Bids), len(event.Asks))
    }

    // Define error handler
    errHandler := func(err error) {
        log.Printf("WebSocket error: %v", err)
    }

    // Subscribe to depth stream
    doneCh, stopCh, err := wsClient.WsDepthServe("BTCUSDT", depthHandler, errHandler)
    if err != nil {
        log.Fatal(err)
    }

    // Stop after 30 seconds
    go func() {
        time.Sleep(30 * time.Second)
        stopCh <- struct{}{}
    }()

    // Wait for completion
    <-doneCh
}
```

#### Kline/Candlestick Stream
```go
// Kline stream handler
klineHandler := func(event *binance_connector.WsKlineEvent) {
    k := event.Kline
    fmt.Printf("Symbol: %s, Open: %s, High: %s, Low: %s, Close: %s, Volume: %s\n",
        k.Symbol, k.Open, k.High, k.Low, k.Close, k.Volume)
}

doneCh, stopCh, err := wsClient.WsKlineServe("BTCUSDT", "1m", klineHandler, errHandler)
```

#### Trade Stream
```go
// Trade stream handler
tradeHandler := func(event *binance_connector.WsTradeEvent) {
    fmt.Printf("Trade: %s %s@%s (Qty: %s)\n", 
        event.Symbol, event.Side, event.Price, event.Quantity)
}

doneCh, stopCh, err := wsClient.WsTradeServe("BTCUSDT", tradeHandler, errHandler)
```

#### 24hr Ticker Stream
```go
// Ticker stream handler
tickerHandler := func(event *binance_connector.WsTicker24HrEvent) {
    fmt.Printf("24hr Ticker: %s - Price: %s, Change: %s%%\n",
        event.Symbol, event.LastPrice, event.PriceChangePercent)
}

doneCh, stopCh, err := wsClient.WsTicker24HrServe("BTCUSDT", tickerHandler, errHandler)
```

### Combined Streams

Subscribe to multiple streams in a single connection:

```go
package main

import (
    "fmt"
    "log"
    "time"

    binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
    // Initialize combined stream client
    wsClient := binance_connector.NewWebsocketStreamClient(true)

    // Combined depth handler
    combinedDepthHandler := func(event *binance_connector.WsDepthEvent) {
        fmt.Printf("[%s] Depth Update - Bids: %d, Asks: %d\n", 
            event.Symbol, len(event.Bids), len(event.Asks))
    }

    errHandler := func(err error) {
        log.Printf("WebSocket error: %v", err)
    }

    // Subscribe to multiple symbols
    symbols := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT"}
    doneCh, stopCh, err := wsClient.WsCombinedDepthServe(symbols, combinedDepthHandler, errHandler)
    if err != nil {
        log.Fatal(err)
    }

    // Stop after 60 seconds
    go func() {
        time.Sleep(60 * time.Second)
        stopCh <- struct{}{}
    }()

    <-doneCh
}
```

### User Data Streams

For private account and order updates:

```go
// Create listen key for user data stream
listenKeyResp, err := client.NewCreateListenKeyService().Do(context.Background())
if err != nil {
    log.Fatal(err)
}

// User data stream handler
userDataHandler := func(event *binance_connector.WsUserDataEvent) {
    switch event.EventType {
    case "outboundAccountPosition":
        fmt.Println("Account update:", binance_connector.PrettyPrint(event))
    case "executionReport":
        fmt.Println("Order update:", binance_connector.PrettyPrint(event))
    }
}

doneCh, stopCh, err := wsClient.WsUserDataServe(listenKeyResp.ListenKey, userDataHandler, errHandler)
```

### Stream Management

```go
// Graceful shutdown example
func gracefulShutdown() {
    wsClient := binance_connector.NewWebsocketStreamClient(false)
    
    depthHandler := func(event *binance_connector.WsDepthEvent) {
        // Handle depth updates
    }
    
    errHandler := func(err error) {
        log.Printf("WebSocket error: %v", err)
    }

    doneCh, stopCh, err := wsClient.WsDepthServe("BTCUSDT", depthHandler, errHandler)
    if err != nil {
        log.Fatal(err)
    }

    // Handle shutdown signals
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-c
        log.Println("Shutting down WebSocket...")
        stopCh <- struct{}{}
    }()

    <-doneCh
    log.Println("WebSocket closed")
}
```

## ⚡ WebSocket API

The WebSocket API enables real-time trading operations with lower latency than REST API.

### Client Initialization

```go
// Initialize WebSocket API client
wsAPIClient := binance_connector.NewWebsocketAPIClient("your-api-key", "your-secret-key")

// Connect to the API
err := wsAPIClient.Connect()
if err != nil {
    log.Fatal("Connection failed:", err)
}
defer wsAPIClient.Close()
```

### Trading via WebSocket API

#### Place Order
```go
package main

import (
    "context"
    "fmt"
    "log"

    binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
    // Initialize WebSocket API client
    client := binance_connector.NewWebsocketAPIClient("your-api-key", "your-secret-key")
    
    // Connect
    err := client.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Place new order via WebSocket API
    orderResp, err := client.NewPlaceNewOrderService().
        Symbol("BTCUSDT").
        Side("BUY").
        Type("LIMIT").
        TimeInForce("GTC").
        Quantity(0.001).
        Price(40000.00).
        Do(context.Background())
    
    if err != nil {
        log.Printf("Order failed: %v", err)
        return
    }

    fmt.Println("Order placed:", binance_connector.PrettyPrint(orderResp))

    // Wait for close signal
    client.WaitForCloseSignal()
}
```

#### Cancel Order
```go
// Cancel order via WebSocket API
cancelResp, err := client.NewCancelOrderService().
    Symbol("BTCUSDT").
    OrderId(12345).
    Do(context.Background())
```

#### Query Account Information
```go
// Get account information via WebSocket API
accountResp, err := client.NewAccountInformationService().
    Do(context.Background())
```

#### Query Order History
```go
// Query OCO history
ocoHistory, err := client.NewAccountOCOHistoryService().
    Do(context.Background())

// Query order history
orderHistory, err := client.NewOrderHistoryService().
    Symbol("BTCUSDT").
    Do(context.Background())
```

### Market Data via WebSocket API

```go
// Get order book via WebSocket API
depthResp, err := client.NewDepthService().
    Symbol("BTCUSDT").
    Limit(100).
    Do(context.Background())

// Get recent trades
tradesResp, err := client.NewRecentTradesService().
    Symbol("BTCUSDT").
    Limit(500).
    Do(context.Background())

// Get klines
klinesResp, err := client.NewKlinesService().
    Symbol("BTCUSDT").
    Interval("1h").
    Limit(100).
    Do(context.Background())
```

### Connection Management

```go
func managedWebSocketAPI() {
    client := binance_connector.NewWebsocketAPIClient("api-key", "secret-key")
    
    // Connect with error handling
    if err := client.Connect(); err != nil {
        log.Fatal("Failed to connect:", err)
    }
    
    // Ensure proper cleanup
    defer func() {
        client.Close()
        log.Println("WebSocket API client closed")
    }()
    
    // Your trading logic here...
    
    // Wait for external close signal or handle gracefully
    client.WaitForCloseSignal()
}
```

## 🌍 Environment Configuration

### Production URLs

Binance provides multiple production endpoints for load balancing and redundancy:

- **Primary**: `https://api.binance.com`
- **Alternatives**: 
  - `https://api1.binance.com`
  - `https://api2.binance.com`  
  - `https://api3.binance.com`

### WebSocket URLs

- **Production Stream**: `wss://stream.binance.com:9443`
- **Production API**: `wss://ws-api.binance.com:443`

### Testnet Support

#### REST API Testnet
```go
client := binance_connector.NewClient("testnet-api-key", "testnet-secret-key", "https://testnet.binance.vision")
```

#### WebSocket Stream Testnet  
```go
wsClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
```

#### WebSocket API Testnet
```go
wsAPIClient := binance_connector.NewWebsocketAPIClient("testnet-api-key", "testnet-secret-key")
// Testnet WebSocket API URL is configured automatically when using testnet credentials
```

### Getting Testnet Credentials

1. Visit [Binance Testnet](https://testnet.binance.vision/)
2. Create an account or log in
3. Generate API key and secret
4. Fund your testnet account with test assets

**📖 Testnet Guide**: [Complete testnet setup instructions](https://dev.binance.vision/t/binance-testnet-environments/99)

## 🛠️ Development & Testing

### Prerequisites
- **Go**: Version 1.23 or higher
- **Binance Account**: For production API access
- **Testnet Account**: For development and testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detection  
go test -race ./...

# Run specific test file
go test -v ./account_test.go
```

### Code Quality
```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Tidy dependencies
go mod tidy
```

### Running Examples
```bash
# Navigate to specific example
cd examples/account/CreateOrder

# Run the example
go run CreateOrder.go
```

## 🎯 Best Practices

### 1. Error Handling
Always handle errors appropriately and check for API-specific error codes:

```go
if err != nil {
    if apiErr, ok := err.(*handlers.APIError); ok {
        switch apiErr.Code {
        case -1013:
            log.Println("Invalid quantity")
        case -2010:
            log.Println("Insufficient balance")
        default:
            log.Printf("API Error: %d - %s", apiErr.Code, apiErr.Message)
        }
    }
    return
}
```

### 2. Context Usage
Always use context for request timeout and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

response, err := client.NewCreateOrderService().
    Symbol("BTCUSDT").
    Do(ctx)
```

### 3. Rate Limiting
Respect Binance rate limits by implementing proper delays:

```go
import "time"

// Add delay between requests
time.Sleep(100 * time.Millisecond)
```

### 4. WebSocket Reconnection
Implement reconnection logic for WebSocket streams:

```go
func reconnectableStream() {
    for {
        wsClient := binance_connector.NewWebsocketStreamClient(false)
        doneCh, stopCh, err := wsClient.WsDepthServe("BTCUSDT", handler, errHandler)
        
        if err != nil {
            log.Printf("Connection failed, retrying in 5s: %v", err)
            time.Sleep(5 * time.Second)
            continue
        }
        
        <-doneCh
        log.Println("Connection lost, reconnecting...")
        time.Sleep(1 * time.Second)
    }
}
```

## 📊 Output Formatting

### PrettyPrint vs Standard Output

The library provides a `PrettyPrint` function for readable JSON output:

```go
// Standard output - single line, hard to read
fmt.Println(response)

// Pretty print - formatted JSON, easy to read
fmt.Println(binance_connector.PrettyPrint(response))
```

**Standard Output Example:**
```
&{depthUpdate 1680092520368 LTCBTC 1989614201 1989614210 [...]}
```

**PrettyPrint Output Example:**
```json
{
  "e": "depthUpdate",
  "E": 1680092041346,
  "s": "LTCBTC",
  "U": 1989606566,
  "u": 1989606596,
  "b": [
    {
      "Price": "0.00322800",
      "Quantity": "83.05100000"
    }
  ],
  "a": [
    {
      "Price": "0.00322900", 
      "Quantity": "79.52900000"
    }
  ]
}
```

## ⚠️ Limitations

This library focuses on **Spot Trading APIs only**. The following are **NOT supported**:

- **Futures Trading**: `/fapi/*` endpoints
- **Delivery Futures**: `/dapi/*` endpoints  
- **Options Trading**: `/vapi/*` endpoints
- **Associated WebSocket streams** for futures/options

For futures trading, use the official Binance Futures Go connector.

## 🧪 Troubleshooting

### Common Issues

#### 1. Timestamp Errors
```
Error: Timestamp for this request is outside of the recvWindow
```
**Solution**: Adjust time offset
```go
client.TimeOffset = -1000 // Adjust based on your system
```

#### 2. Signature Errors
```
Error: Signature for this request is not valid
```
**Solution**: Verify API credentials and ensure proper parameter ordering

#### 3. Rate Limit Exceeded
```
Error: Too many requests
```
**Solution**: Implement exponential backoff and respect rate limits

#### 4. WebSocket Connection Issues
```
Error: WebSocket connection failed
```
**Solution**: Check network connectivity and implement reconnection logic

### Debug Mode
Enable debug logging to troubleshoot issues:

```go
client.Debug = true
// This will log all HTTP requests and responses
```

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **🐛 Bug Reports**: Open an issue with detailed reproduction steps
2. **💡 Feature Requests**: Suggest new features or improvements  
3. **🔧 Code Contributions**: Submit pull requests with tests
4. **📚 Documentation**: Improve examples and documentation

### Before Contributing
- Ensure all tests pass: `go test ./...`
- Format your code: `go fmt ./...`
- Follow existing code patterns and conventions
- Add tests for new functionality

### API Issues
For Binance API-related issues (not library bugs), please visit the [Binance Developer Community](https://dev.binance.vision).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## 🔗 Resources

- **[Binance API Documentation](https://github.com/binance/binance-spot-api-docs)** - Official API docs
- **[Binance Developer Community](https://dev.binance.vision)** - Community support
- **[Binance Testnet](https://testnet.binance.vision/)** - Testing environment
- **[Go Documentation](https://pkg.go.dev/github.com/luciano-personal-org/binance-connector)** - Package documentation

---

⭐ If this library helps you, please consider giving it a star!

For questions or support, feel free to open an issue or visit the Binance Developer Community.
