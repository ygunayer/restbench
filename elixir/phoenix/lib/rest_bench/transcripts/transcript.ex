defmodule RestBench.Transcripts.Transcript do
  defstruct [:id, :status, :enabled, :times_transcribed, :video_id, :provider, :data]

  defimpl Jason.Encoder, for: __MODULE__ do
    def encode(transcript, _options) do
      Jason.encode!(%{
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
