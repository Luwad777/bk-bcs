/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package logging

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/config"
)

var loggerInitOnce sync.Once

// use zap for better performance
var logger *zap.Logger
var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func InitLogger(logConf *config.LogConf) {
	loggerInitOnce.Do(func() {
		// json logger
		logger = newZapJSONLogger(logConf)
	})
}

// 修改时间编码器并设置日志级别为大写，比如DEBUG/INFO
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "call",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
}

func newZapJSONLogger(cfg *config.LogConf) *zap.Logger {
	writer, err := getWriter(cfg.WriterType, cfg.Settings)
	if err != nil {
		panic(err)
	}
	w := &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(writer),
		FlushInterval: time.Duration(cfg.FlushInterval) * time.Second,
	}

	// 设置日志级别
	l, ok := levelMap[cfg.Level]
	if !ok {
		l = zap.InfoLevel
	}

	core := zapcore.NewCore(getEncoder(), w, l)
	return zap.New(core)
}

// GetLogger ...
// TODO: 是否分为不同的类型，比如请求第三方、API等，根据不同的配置，设置不同的日志
func GetLogger() *zap.Logger {
	// 如果要进一步性能，可以使用SugaredLogger
	if logger == nil {
		logger, _ = zap.NewProduction()
	}

	return logger
}
