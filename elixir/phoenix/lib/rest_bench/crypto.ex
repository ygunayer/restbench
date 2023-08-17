defmodule RestBench.Crypto do
  alias Argon2, as: PasswordHasher

  def hash_password(password), do: PasswordHasher.hash_pwd_salt(password)

  def verify_password?(clear_password, hashed_password),
    do: PasswordHasher.verify_pass(clear_password, hashed_password)
end
