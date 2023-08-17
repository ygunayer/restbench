defmodule RestBenchWeb.BaseController do
  require Logger

  use RestBenchWeb, :controller

  defmacro __using__(_) do
    quote do
      use RestBenchWeb, :controller

      import RestBenchWeb.BaseController, only: [respond: 2]
    end
  end

  def respond(nil, conn), do: json(conn, nil)
  def respond(:ok, conn), do: json(conn, %{})
  def respond({:ok, result}, conn), do: json(conn, %{"data" => result})

  def respond({:error, %Ecto.Changeset{valid?: false} = changeset}, conn) do
    error_messages =
      Ecto.Changeset.traverse_errors(changeset, fn {msg, opts} ->
        Regex.replace(~r"%{(\w+)}", msg, fn _, key ->
          opts |> Keyword.get(String.to_existing_atom(key), key) |> to_string()
        end)
      end)
      |> Enum.map(fn {key, msg} -> "#{key} #{msg}" end)
      |> Enum.join(", ")

    conn
    |> put_status(400)
    |> json(%{
      "error" => %{"message" => "Validation(s) failed", "description" => error_messages}
    })
  end

  def respond({:error, err}, conn) do
    Logger.error(err)

    conn
    |> put_status(500)
    |> json(%{"error" => "An unknown error has occurred"})
  end
end
