#!/usr/bin/env sh
SECRET_KEY_BASE=E+HxBSYTS4Oi/MhVQ5zaEaTwdeVNBG8k74sf3kHmXwiNq0ZyARE7S7O49wUQGjGg DATABASE_URL=postgres://restbench:restbench@localhost:5432/restbench?sslmode=disable PORT=4000 MIX_ENV=prod mix phx.server
