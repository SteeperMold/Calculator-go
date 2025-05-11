import {Link} from "react-router";
import {useUser} from "../../shared/hooks/useUser";

const MainPage = () => {
  const {user} = useUser();

  return (
    <div className="flex flex-col items-center">
      <h1 className="text-3xl mt-6 mb-12">
        Calculator-go
      </h1>
      <div>
        <p className="text-xl">
          На этом сервисе вы можете вычислить и посмотреть результат арифметического выражения!
        </p>
      </div>
      <div className="max-w-3xl p-4">
        <h2 className="text-center font-semibold mt-6 mb-10 text-3xl">Для этого вы можете:</h2>
        <div className="flex flex-row items-center">
          {user ? (
            <Link
              className="hover:underline w-1/3 text-center font-bold"
              to="/calculate"
            >
              Добавить арифметическое выражение
            </Link>
          ) : (
            <Link
              className="hover:underline w-1/3 text-center font-bold"
              to="/login"
            >
              Войти, чтобы добавить арифметическое выражение
            </Link>
          )}
          <p className="w-1/3 text-center">
            или
          </p>
          <Link
            className="hover:underline w-1/3 text-center font-bold"
            to="/expressions"
          >
            Получить выражение
          </Link>
        </div>
      </div>
    </div>
  );
};

export default MainPage;
