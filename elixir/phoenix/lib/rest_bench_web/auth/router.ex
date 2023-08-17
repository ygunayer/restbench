defmodule RestBenchWeb.Auth.Router do
  use RestBenchWeb, :router

  scope "/", RestBenchWeb do
    post "/register", Auth.Controller, :register
  end
end
