defmodule HelloWeb.DefaultController do
  use HelloWeb, :controller

  def index(conn,_params) do
   text conn,"The real Deal API is live -#{Mix.env()}"
  end

  def gettext(conn, _params) do
    text conn, "The real Deal API is live2 -#{Mix.env()}"
  end
  end
