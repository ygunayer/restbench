defmodule RestBench.Transcripts.TranscriptData do
  defstruct [:id, :data, :speaker, :start_time, :end_time]

  defimpl Jason.Encoder, for: __MODULE__ do
    def encode(data, _options) do
      Jason.encode!(%{
        "id" => data.id,
        "data" => data.data,
        "speaker" => data.speaker,
        "startTimeNanoSeconds" => data.start_time,
        "endTimeNanoSeconds" => data.end_time
      })
    end
  end
end
