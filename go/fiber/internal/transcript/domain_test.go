package transcript

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ReadFixture(filename string, v any) error {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("Failed to get caller information")
	}

	path := path.Join(path.Dir(currentFile), "..", "..", "test", "fixtures", filename)

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}

	return nil
}

func doTestTimestampToTranscriptTime(timestamp float64, expected TranscriptTime, t *testing.T) {
	actual := *TimestampToTranscriptTime(timestamp)
	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestTimestampToTranscriptTime(t *testing.T) {
	doTestTimestampToTranscriptTime(2382000000, TranscriptTime{Seconds: "2", Nanos: 382000000}, t)
	doTestTimestampToTranscriptTime(3382000000, TranscriptTime{Seconds: "3", Nanos: 382000000}, t)
	doTestTimestampToTranscriptTime(3615547000000, TranscriptTime{Seconds: "3615", Nanos: 547000000}, t)
	doTestTimestampToTranscriptTime(3635996999756, TranscriptTime{Seconds: "3635", Nanos: 996999756}, t)
}

func TestSplitIntoWords(t *testing.T) {
	v1 := &TranscriptV1{}
	if err := ReadFixture("transcript.v1.json", v1); err != nil {
		t.Fatal(err)
	}

	v2 := &Transcript{}
	if err := ReadFixture("transcript.v2.json", v2); err != nil {
		t.Fatal(err)
	}

	actual := SplitIntoWords(v2)
	diff := cmp.Diff(v1, actual, cmp.AllowUnexported(TranscriptV1{}))
	if diff != "" {
		t.Error(diff)
	}
}
