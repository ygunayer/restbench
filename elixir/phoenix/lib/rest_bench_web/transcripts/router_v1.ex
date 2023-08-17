defmodule RestBenchWeb.Transcripts.RouterV1 do
  use RestBenchWeb, :router

  scope "/", RestBenchWeb do
    get "/:id", Transcripts.Controller, :get_v1
  end
end
