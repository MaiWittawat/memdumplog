package adapter

import (
	"time"

	"github.com/MaiWittawat/memdumplog/store"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	zapcore.Core
	store store.Store
}

func NewZap(core zapcore.Core, store store.Store) zapcore.Core {
	return ZapCore{Core: core, store: store}
}

func UseZap(store store.Store) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(zapcore.Lock(zapcore.AddSync(zapcore.AddSync(nil)))),
		cfg.Level,
	)

	logger := zap.New(ZapCore{
		Core:  core,
		store: store,
	})

	zap.ReplaceGlobals(logger)
}

func (z ZapCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	z.store.Add(store.Entry{
		Level:   ent.Level.String(),
		Message: ent.Message,
		Time:    ent.Time.Format(time.RFC3339),
	})
	return z.Core.Write(ent, fields)
}
