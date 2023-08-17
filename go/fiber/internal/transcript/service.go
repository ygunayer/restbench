package transcript

import (
	"strconv"

	"github.com/ygunayer/restbench/internal/cassandra"
)

func GetTranscript(transcriptId string) (*Transcript, error) {
	cass := cassandra.Get()

	var id string
	var enabled bool
	var provider string
	var status string
	var timesTranscribed int32
	var videoId string

	if err := cass.Query("SELECT id, enabled, provider, status, times_transcribed, video_id FROM transcripts WHERE id = ?", transcriptId).Scan(
		&id,
		&enabled,
		&provider,
		&status,
		&timesTranscribed,
		&videoId,
	); err != nil {
		return nil, err
	}

	transcriptData, err := GetTranscriptDataByTranscriptId(transcriptId)

	if err != nil {
		return nil, err
	}

	transcript := Transcript{
		Id:               id,
		Enabled:          enabled,
		Provider:         provider,
		Status:           status,
		TimesTranscribed: timesTranscribed,
		VideoId:          videoId,
		Data:             *transcriptData,
	}

	return &transcript, nil
}

func GetTranscriptDataByTranscriptId(transcriptId string) (*[]TranscriptData, error) {
	cass := cassandra.Get()

	var allTranscriptData = make([]TranscriptData, 0)

	iter := cass.Query("SELECT id, data, speaker, start_time, end_time FROM transcript_data WHERE transcript_id = ?", transcriptId).Iter()

	var id int32
	var data string
	var speaker string
	var startTimeStr string
	var endTimeStr string

	for iter.Scan(&id, &data, &speaker, &startTimeStr, &endTimeStr) {
		startTime, err := strconv.ParseFloat(startTimeStr, 64)
		if err != nil {
			return nil, err
		}

		endTime, err := strconv.ParseFloat(endTimeStr, 64)
		if err != nil {
			return nil, err
		}

		allTranscriptData = append(allTranscriptData, TranscriptData{
			Id:                   id,
			Data:                 data,
			Speaker:              speaker,
			TranscriptId:         transcriptId,
			StartTimeNanoSeconds: startTime,
			EndTimeNanoSeconds:   endTime,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &allTranscriptData, nil
}

func GetTranscriptV1(id string) (*TranscriptV1, error) {
	if transcript, err := GetTranscript(id); err != nil {
		return nil, err
	} else {
		return SplitIntoWords(transcript), nil
	}
}
