defmodule RestBench.Transcripts.Transcripts do
  alias RestBench.Cassandra

  alias RestBench.Transcripts.Transcript
  alias RestBench.Transcripts.TranscriptData

  def find_by_id(id) do
    with {:ok, row} when not is_nil(row) <-
           Cassandra.find_one(
             "SELECT id, enabled, provider, status, times_transcribed, video_id FROM transcripts WHERE id = ?",
             [id]
           ),
         {:ok, data} <- find_data_by_transcript_id(id) do
      {:ok,
       %Transcript{
         id: row["id"],
         enabled: row["enabled"],
         provider: row["provider"],
         status: row["status"],
         times_transcribed: row["times_transcribed"],
         video_id: row["video_id"],
         data: data
       }}
    else
      other -> other
    end
  end

  def find_data_by_transcript_id(transcript_id) do
    case Cassandra.find_all(
           "SELECT id, data, speaker, start_time, end_time FROM transcript_data WHERE transcript_id = ?",
           [transcript_id]
         ) do
      {:ok, rows} ->
        data =
          Enum.map(rows, fn row ->
            %TranscriptData{
              id: row["id"],
              data: row["data"],
              speaker: row["speaker"],
              start_time: row["start_time"] |> String.to_integer(),
              end_time: row["end_time"] |> String.to_integer()
            }
          end)

        {:ok, data}

      other ->
        other
    end
  end
end
