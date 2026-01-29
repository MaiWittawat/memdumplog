# memdump-log

`memdump-log` is an **in-memory log collector** for Go applications.

It captures logs from common Go loggers (Logrus, Zap, Zerolog, slog),
stores them in memory (ring buffer), and allows you to **retrieve logs
programmatically**, typically to expose them via an HTTP API or admin UI.

---

## Installation
```powershell
go get github.com/MaiWittawat/memdumplog
```

## Why memdump-log?

This library is designed for scenarios where you want to:

- View logs through an HTTP API
- Debug production systems without SSH access
- Build an admin log viewer
- Temporarily inspect recent logs (ring buffer)
- Keep logging flow unchanged

`memdump-log` does **not** replace centralized logging systems (ELK, Loki, etc.).
It complements them.

---

## High-Level Concept

Your application continues to log normally.
`memdump-log` hooks into the logger and mirrors logs into memory.

---

## Project Structure

```text
memdump-log/
├── adapter/
│   ├── logrus.go
│   ├── zap.go
│   ├── zerolog.go
│   └── slog.go
├── store/
│   ├── store.go
│   └── memory.go
├── config.go
└── logger.go
```


## Log Entry Structure

Each captured log is stored as a simple struct:
```go
type Entry struct {
	Level   string
	Message string
	Time    string
}
```

## Basic Usage
```go
logs, err := memdumplog.New(memdumplog.Config{
	Driver:     memdumplog.Logrus,
	BufferSize: 200,
})
if err != nil {
	panic(err)
}
```

## Configuration
| Field        | Description                                                     |
| ------------ | --------------------------------------------------------------- |
| `Driver`     | Logger driver to hook into (`Logrus`, `Zap`, `Zerolog`, `Slog`) |
| `BufferSize` | Maximum number of log entries kept in memory (ring buffer)      |

## Example: Expose Logs via HTTP API (Gin)

### JSON Endpoint
```go
router.GET("/logs", func(c *gin.Context) {
	c.JSON(http.StatusOK, logs.Logs())
})
```

### Example request
```curl
curl http://localhost:8080/logs
```

## Plain Text (tail-style) Endpoint
```go
router.GET("/logs/raw", func(c *gin.Context) {
	var out strings.Builder

	for _, e := range logs.Logs() {
		out.WriteString(fmt.Sprintf(
			"[%s] %s\n",
			e.Level,
			e.Message,
		))
	}

	c.Data(
		http.StatusOK,
		"text/plain; charset=utf-8",
		[]byte(out.String()),
	)
})
```