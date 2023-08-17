defmodule RestBench.Users do
  import Ecto.Query, warn: false

  alias RestBench.Repo
  alias RestBench.Users.User

  def register_user(params) do
    %User{}
    |> User.registration_changeset(params)
    |> Repo.insert()
  end
end
