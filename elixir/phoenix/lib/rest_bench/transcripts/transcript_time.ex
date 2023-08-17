defmodule RestBench.Transcripts.TranscriptTime do
  @derive Jason.Encoder
  defstruct [:seconds, :nanos]

  def new(timestamp_nanos) do
    seconds = floor(timestamp_nanos / 1_000_000_000) |> to_string()
    nanos = timestamp_nanos |> round() |> rem(1_000_000_000)
    %__MODULE__{seconds: seconds, nanos: nanos}
  end
end
