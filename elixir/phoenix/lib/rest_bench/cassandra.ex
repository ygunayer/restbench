defmodule RestBench.Cassandra do
  @name Xandra

  def find_all(query, params \\ []) do
    with {:ok, stmt} <- Xandra.prepare(@name, query),
         {:ok, %Xandra.Page{} = page} <- Xandra.execute(Xandra, stmt, params) do
      {:ok, Enum.to_list(page)}
    else
      other -> other
    end
  end

  def find_one(query, params \\ []) do
    case find_all(query, params) do
      {:ok, [row | _]} -> {:ok, row}
      {:ok, []} -> {:ok, nil}
      other -> other
    end
  end
end
