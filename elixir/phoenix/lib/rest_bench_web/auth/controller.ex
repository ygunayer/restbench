defmodule RestBenchWeb.Auth.Controller do
  use RestBenchWeb.BaseController

  alias RestBench.Users

  def register(conn, params) do
    case Users.register_user(params) do
      {:ok, _} ->
        resp(conn, 201, "")

      other ->
        respond(other, conn)
    end
  end
end
