defmodule RestBench.Transcripts.TranscriptDataV1 do
  defstruct [:speaker, :word, :start_time, :end_time]

  alias RestBench.Transcripts.TranscriptData
  alias RestBench.Transcripts.TranscriptTime

  def from_v2(%TranscriptData{
        data: data,
        speaker: speaker,
        start_time: start_time,
        end_time: end_time
      }) do
    words = String.split(data, " ")

    shift = (end_time - start_time) / length(words)

    {_, _, v1} =
      Enum.reduce(words, {start_time, end_time, []}, fn word, {start_time, end_time, data} ->
        this_data = %__MODULE__{
          speaker: speaker,
          word: word,
          start_time: TranscriptTime.new(start_time),
          end_time: TranscriptTime.new(end_time)
        }

        {start_time + shift, end_time + shift, [this_data | data]}
      end)

    Enum.reverse(v1)
  end

  defimpl Jason.Encoder, for: __MODULE__ do
    def encode(data, _options) do
      Jason.encode!(%{
        "speaker" => data.speaker,
        "word" => data.word,
        "startTimeNanoSeconds" => data.start_time,
        "endTimeNanoSeconds" => data.end_time
      })
    end
  end
end
