package transcript

import (
	"fmt"
	"math"
	"strings"
)

type TranscriptV1 struct {
	UnderscoreId     string               `json:"_id"`
	Id               string               `json:"id"`
	Video            string               `json:"video"`
	Status           string               `json:"status"`
	ShowTranscript   bool                 `json:"showTranscript"`
	Provider         string               `json:"provider"`
	TimesTranscribed int32                `json:"timesTranscribed"`
	Data             [][]TranscriptDataV1 `json:"data"`
}

type TranscriptDataV1 struct {
	Speaker   string          `json:"speaker"`
	Word      string          `json:"word"`
	StartTime *TranscriptTime `json:"startTime"`
	EndTime   *TranscriptTime `json:"endTime"`
}

type TranscriptTime struct {
	Seconds string `json:"seconds"`
	Nanos   uint64 `json:"nanos"`
}

type Transcript struct {
	Id               string           `json:"id"`
	VideoId          string           `json:"videoId"`
	Status           string           `json:"status"`
	Enabled          bool             `json:"enabled"`
	Provider         string           `json:"provider"`
	TimesTranscribed int32            `json:"timesTranscribed"`
	Data             []TranscriptData `json:"data"`
}

type TranscriptData struct {
	Id                   int32   `json:"id"`
	Data                 string  `json:"data"`
	Speaker              string  `json:"speaker"`
	TranscriptId         string  `json:"transcriptId"`
	StartTimeNanoSeconds float64 `json:"startTimeNanoSeconds"`
	EndTimeNanoSeconds   float64 `json:"endTimeNanoSeconds"`
}

func TimestampToTranscriptTime(t float64) *TranscriptTime {
	seconds := uint64(math.Floor(t / 1_000_000_000))
	remaining := uint64(math.Round(math.Mod(t, 1_000_000_000)))
	return &TranscriptTime{
		Seconds: fmt.Sprintf("%d", seconds),
		Nanos:   remaining,
	}
}

func SplitIntoWords(t *Transcript) *TranscriptV1 {
	result := &TranscriptV1{
		UnderscoreId:     t.Id,
		Id:               t.Id,
		Video:            t.VideoId,
		Status:           t.Status,
		ShowTranscript:   t.Enabled,
		Provider:         t.Provider,
		TimesTranscribed: t.TimesTranscribed,
	}

	data := make([][]TranscriptDataV1, len(t.Data))

	for i, speech := range t.Data {
		words := strings.Split(speech.Data, " ")

		shift := float64(speech.EndTimeNanoSeconds-speech.StartTimeNanoSeconds) / float64(len(words))

		startTimeNanos := speech.StartTimeNanoSeconds
		endTimeNanos := startTimeNanos + shift

		v1Data := make([]TranscriptDataV1, len(words))

		for j, word := range words {
			v1Data[j] = TranscriptDataV1{
				Speaker:   speech.Speaker,
				StartTime: TimestampToTranscriptTime(startTimeNanos),
				EndTime:   TimestampToTranscriptTime(endTimeNanos),
				Word:      word,
			}

			startTimeNanos += shift
			endTimeNanos += shift
		}

		data[i] = v1Data
	}

	result.Data = data

	return result
}
