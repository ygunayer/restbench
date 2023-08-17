defmodule RestBenchWeb.Transcripts.RouterV2 do
  use RestBenchWeb, :router

  scope "/", RestBenchWeb do
    get "/:id/full", Transcripts.Controller, :get_v2
  end
end
