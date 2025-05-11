import React, {useState} from "react";
import {useNavigate} from "react-router";
import {useUser} from "../../shared/hooks/useUser";
import Api from "../../api";

const LoginPage = () => {
  const [error, setError] = useState(null);
  const {updateUser} = useUser();
  const navigate = useNavigate();

  const onSubmit = (event) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const login = formData.get("login");
    const password = formData.get("password");

    Api.post("/login", {login, password})
      .then(response => {
        localStorage.setItem("accessToken", response.data.accessToken);
        localStorage.setItem("refreshToken", response.data.refreshToken);
        updateUser();
        navigate("/login");
      })
      .catch(error => {
        if (error.status === 404) {
          setError("Пользователь с указанным адресом электронной почты не существует");
          return;
        }
        if (error.status === 401) {
          setError("Неверный пароль или адрес электронной почты");
          return;
        }
        setError("Не удалось войти в аккаунт");
      });
  };

  return <>
    <h1 className="text-3xl text-center mt-10">Вход в аккаунт</h1>

    <form onSubmit={onSubmit} className="mx-auto w-1/4 mt-16 flex flex-col items-center">
      {error && <h2 className="self-start mb-8 text-xl text-red-600">{error}</h2>}

      <input
        type="text"
        required={true}
        name="login"
        placeholder="Логин"
        className="mb-10 w-full outline-none bg-inherit border-b-2 py-2 dark:placeholder:text-dark-text-primary"
      />

      <input
        type="password"
        autoComplete="current-password"
        required={true}
        name="password"
        placeholder="Пароль"
        className="w-full outline-none bg-inherit border-b-2 py-2 dark:placeholder:text-dark-text-primary"
      />

      <button
        type="submit"
        className="mt-14 text-xl px-7 py-2 rounded dark:bg-dark-secondary"
      >
        Войти в аккаунт →
      </button>
    </form>
  </>;
}

export default LoginPage;
