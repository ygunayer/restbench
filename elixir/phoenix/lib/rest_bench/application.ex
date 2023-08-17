defmodule RestBench.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false
  require Logger

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Telemetry supervisor
      RestBenchWeb.Telemetry,
      # Start the Ecto repository
      RestBench.Repo,
      # Start the PubSub system
      {Phoenix.PubSub, name: RestBench.PubSub},
      # Start Finch
      {Finch, name: RestBench.Finch},
      # Start the Endpoint (http/https)
      RestBenchWeb.Endpoint,
      {Xandra, get_xandra_args()}
      # Start a worker by calling: RestBench.Worker.start_link(arg)
      # {RestBench.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: RestBench.Supervisor]
    Supervisor.start_link(children, opts)
  end

  defp get_xandra_args() do
    xandra_options = Application.get_env(:rest_bench, :xandra, [])

    nodes = Keyword.get(xandra_options, :nodes, ["localhost:9042"])
    default_keyspace = Keyword.get(xandra_options, :keyspace, "data")

    [
      name: Xandra,
      nodes: nodes,
      after_connect: fn conn ->
        Xandra.execute(conn, "USE #{default_keyspace}")

        Logger.info(
          "Connected to Cassandra at #{inspect(nodes)} using default keyspace #{default_keyspace}"
        )
      end
    ]
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    RestBenchWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
