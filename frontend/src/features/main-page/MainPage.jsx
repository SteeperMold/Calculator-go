import {Link} from "react-router";

const MainPage = () => {
  return (
    <div className="flex flex-col items-center mt-[10vh]">
      <h1 className="text-3xl mt-6 mb-3">
        Calculator-go
      </h1>
      <div>
        <p className="text-xl">
          На этом сервисе вы можете вычислить и посмотреть результат арифметического выражения!
        </p>
      </div>
      <div className="max-w-3xl p-4">
        <h2 className="text-center font-semibold mb-6 text-3xl">Для этого вы можете:</h2>
        <div className="flex flex-row items-center">
          <Link
            className="hover:underline w-1/3 text-center font-bold"
            to="/calculate">
            Добавить арифметическое выражение
          </Link>
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
