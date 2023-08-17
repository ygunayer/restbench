defmodule RestBench.Repo do
  use Ecto.Repo,
    otp_app: :rest_bench,
    adapter: Ecto.Adapters.Postgres
end
