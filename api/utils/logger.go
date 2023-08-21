package utils

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

var logger zerolog.Logger

func InitLogger() {
	logFile := &lumberjack.Logger{
		Filename:   "server.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
	localLogger := zerolog.New(logFile).With().Timestamp().Logger()
	influxdbLogger := zerolog.Logger{}
	logger = localLogger
	isUsingInfluxDB := false

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if influxClient := NewInfluxDBClient(); influxClient.Client != nil {
					if !isUsingInfluxDB {
						logger = influxdbLogger
						logger = logger.Hook(InfluxDBHook{Client: influxClient})
						isUsingInfluxDB = true
					}
					//uploadRotatedLogsToInfluxDB(influxClient.Client)
				} else {
					if isUsingInfluxDB {
						logger = localLogger
						isUsingInfluxDB = false
					}
					logData := LogData{Event: "InfluxDB connection failed, using local logs", User: "system"}
					LogError(logData)
				}
				time.Sleep(time.Minute)
			}
		}
	}()

	time.Sleep(time.Second)
}

func uploadRotatedLogsToInfluxDB(influxClient influxdb2.Client) {
	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()

	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	files, err := filepath.Glob("app.log*")
	if err != nil {
		data := LogData{User: "system", Event: "Fail to find log file", Details: "error: " + err.Error()}
		LogError(data)
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			data := LogData{User: "system", Event: "Fail to open log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			point := influxdb2.NewPointWithMeasurement("logs").
				SetTime(time.Now()).
				AddField("message", line)

			err := writeAPI.WritePoint(context.Background(), point)
			if err != nil {
				data := LogData{User: "system", Event: "Fail to upload log file", Details: "File name: " + filename + " error: " + err.Error()}
				LogError(data)
			}
		}
		err = file.Close()
		if err != nil {
			data := LogData{User: "system", Event: "Fail to close log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
			continue
		}
		err = os.Remove(filename)
		if err != nil {
			data := LogData{User: "system", Event: "Fail to delete log file", Details: "File name: " + filename + " error: " + err.Error()}
			LogError(data)
		}
	}
}

type LogData struct {
	Event   string      `json:"event"`
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Header  interface{} `json:"header"`
	User    string      `json:"user"`
	Details string      `json:"details"`
}

func LogDebug(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Debug().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in debug log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogInfo(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Info().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in info log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogWarn(obj LogData) {

	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Warn().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in warn log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogError(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Error().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in error log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogFatal(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Fatal().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in fatal log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}

func LogPanic(obj LogData) {
	jsonData, err := json.Marshal(obj)
	if err == nil {
		logger.Panic().Msg(string(jsonData))
	} else {
		logger.Error().Msg("\"event\":\"error in panic log\",\"details\":\"" + err.Error() + "\",\"user\":\"system\"")
	}
}