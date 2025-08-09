# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is the **Binance Spot Go Connector** - a lightweight Go library that serves as a connector to the Binance public API. It's a personal fork of the official Binance connector, focusing on spot trading, websockets, and account management.

## Commands

### Development and Testing
- `go test ./...` - Run all tests
- `go test -v ./...` - Run tests with verbose output  
- `go test ./... -race` - Run tests with race detection
- `go mod tidy` - Clean up module dependencies
- `go build` - Build the module
- `go run examples/[category]/[endpoint]/[endpoint].go` - Run specific examples

### Code Quality
- `go fmt ./...` - Format all Go code
- `go vet ./...` - Run Go static analysis
- `goimports -w .` - Organize imports (if installed)

## Architecture

### Core Components

**Client (`client.go`)**
- Central client struct with API credentials, base URL, HTTP client, debug mode, and time offset
- Factory methods for all service endpoints (Account, Market, Wallet, Margin, Sub-Account, etc.)
- Request parsing with HMAC-SHA256 signature generation for authenticated endpoints
- Built-in error handling and response parsing

**Request System (`request.go`)**
- Generic request struct handling HTTP method, endpoint, query params, form data, and security type
- Three security levels: None, API Key required, Signed (timestamp + signature required)
- Price formatting utilities with precision control for trading operations
- Request option pattern for flexible configuration (e.g., `WithRecvWindow`)

**Service Pattern**
- Each endpoint category in separate files: `account.go`, `market.go`, `wallet.go`, `margin.go`, `subaccount.go`, `fiat.go`, `user_stream.go`
- Consistent service struct pattern with client reference and fluent builder methods
- All services implement `Do(context.Context)` method returning typed responses

**WebSocket Support (`websocket.go`, `websocket_api.go`)**
- Stream client for market data and user data streams
- WebSocket API client for real-time trading operations  
- Support for combined streams and individual symbol streams
- Separate handlers for different event types with proper error handling

### Module Structure

- **Root files**: Core client, request handling, constants, and main service categories
- **handlers/**: Error handling utilities
- **examples/**: Comprehensive examples organized by endpoint category
  - `account/`: Trading operations, order management, account info
  - `market/`: Market data, tickers, order books, historical data
  - `wallet/`: Deposits, withdrawals, asset management
  - `margin/`: Margin trading operations
  - `subaccount/`: Sub-account management
  - `websocket/`: WebSocket stream examples
  - `websocket_api/`: WebSocket API examples

### Key Dependencies

- `github.com/bitly/go-simplejson` - JSON parsing and manipulation
- `github.com/gorilla/websocket` - WebSocket implementation
- `github.com/goccy/go-json` - High-performance JSON library
- `github.com/stretchr/testify` - Testing framework

## Development Guidelines

### Price Handling
- Use `setPriceWithPrecision()` for price parameters requiring specific decimal places
- Use `setParamFloat()` for standard 8-decimal precision
- Use `setParamHighFloat()` for 4-decimal precision values

### Authentication
- Testnet: `https://testnet.binance.vision`
- Production alternatives: `https://api1.binance.com`, `https://api2.binance.com`, `https://api3.binance.com`
- API keys must have appropriate permissions for endpoint access
- Signed requests require proper HMAC-SHA256 signature with current timestamp

### Error Handling
- API errors are wrapped in `handlers.APIError` with code and message
- Network errors bubble up as standard Go errors
- WebSocket errors handled via dedicated error handler functions

### WebSocket Usage
- Initialize with `NewWebsocketStreamClient(isCombined, baseURL)`
- Use `isCombined=true` for multiple streams, `false` for individual streams
- Always implement proper cleanup with stop channels and done channels
- WebSocket API requires authentication for trading operations

### Testing
- Unit tests follow `*_test.go` naming convention
- Tests use testify framework for assertions
- Mock HTTP responses for testing without API calls
- Test files organized by service category matching main code structure

## Common Patterns

### Service Initialization and Usage
```go
client := binance_connector.NewClient("apiKey", "secretKey", "baseURL")
service := client.NewCreateOrderService()
response, err := service.Symbol("BTCUSDT").Side("BUY").Type("MARKET").Quantity(0.001).Do(context.Background())
```

### WebSocket Stream Setup
```go
wsClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
doneCh, stopCh, err := wsClient.WsDepthServe("BNBUSDT", depthHandler, errorHandler)
```

### Request Options
```go
response, err := service.Symbol("BTCUSDT").Do(context.Background(), WithRecvWindow(10000))
```

This architecture provides a clean separation of concerns with consistent patterns across all Binance API endpoints while maintaining type safety and proper error handling.