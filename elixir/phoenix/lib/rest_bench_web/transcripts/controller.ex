defmodule RestBenchWeb.Transcripts.Controller do
  use RestBenchWeb.BaseController

  alias RestBench.Transcripts.TranscriptV1
  alias RestBench.Transcripts.Transcripts

  def get_v2(conn, %{"id" => id}) do
    case Transcripts.find_by_id(id) do
      {:ok, transcript} ->
        json(conn, transcript)

      other ->
        respond(other, conn)
    end
  end

  def get_v1(conn, %{"id" => id}) do
    case Transcripts.find_by_id(id) do
      {:ok, transcript} ->
        json(conn, TranscriptV1.from_v2(transcript))

      other ->
        respond(other, conn)
    end
  end
end
