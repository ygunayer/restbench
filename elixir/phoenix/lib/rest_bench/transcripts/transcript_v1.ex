defmodule RestBench.Transcripts.TranscriptV1 do
  defstruct [:id, :status, :enabled, :times_transcribed, :video_id, :provider, :data]

  alias RestBench.Transcripts.Transcript
  alias RestBench.Transcripts.TranscriptDataV1

  def from_v2(%Transcript{} = v2) do
    %__MODULE__{
      id: v2.id,
      status: v2.status,
      enabled: v2.enabled,
      times_transcribed: v2.times_transcribed,
      video_id: v2.video_id,
      provider: v2.provider,
      data: Enum.map(v2.data, &TranscriptDataV1.from_v2/1)
    }
  end

  defimpl Jason.Encoder, for: __MODULE__ do
    def encode(transcript, _options) do
      Jason.encode!(%{
        "_id" => transcript.id,
        "id" => transcript.id,
        "status" => transcript.status,
        "enabled" => transcript.enabled,
        "timesTranscribed" => transcript.times_transcribed,
        "videoId" => transcript.video_id,
        "provider" => transcript.provider,
        "data" => transcript.data
      })
    end
  end
end
