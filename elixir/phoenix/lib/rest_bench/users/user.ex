defmodule RestBench.Users.User do
  use Ecto.Schema
  import Ecto.Changeset

  @email_regex ~r/@/

  schema "users" do
    field(:email, :string)
    field(:status, Ecto.Enum, values: [:deleted, :passive, :pending_activation, :active])
    field(:password_hash, :string)
    field(:password, :string, virtual: true)
    field(:password_confirmation, :string, virtual: true)
    field(:name, :string)
    field(:activation_code, :string)
    field(:activation_code_expires_at, :utc_datetime)

    timestamps()
  end

  @doc false
  def changeset(user, attrs) do
    user
    |> cast(attrs, [:email, :name, :activation_code])
    |> validate_required([:email, :name])
    |> validate_format(:email, @email_regex)
    |> unique_constraint(:email)
  end

  def registration_changeset(user, attrs) do
    user
    |> changeset(attrs)
    |> cast(attrs, [:password, :password_confirmation])
    |> validate_required([:password, :password_confirmation])
    |> validate_length(:password, min: 8)
    |> validate_confirmation(:password)
    |> put_change(:status, :pending_activation)
    |> hash_password()
    |> add_activation_code()
  end

  def add_activation_code(%{valid?: true} = changeset),
    do: put_change(changeset, :activation_code, Ecto.UUID.generate())

  def add_activation_code(other), do: other

  def hash_password(%{valid?: true, changes: %{password: password}} = changeset) do
    hashed_password = RestBench.Crypto.hash_password(password)
    put_change(changeset, :password_hash, hashed_password)
  end

  def hash_password(other), do: other
end
