import {useParams} from "react-router";
import {useEffect, useState} from "react";
import axios from "axios";

const ExpressionPage = () => {
  const {id} = useParams();
  const [error, setError] = useState("");
  const [data, setData] = useState({});

  useEffect(() => {
    axios.get(`http://localhost:8080/api/v1/expressions/${id}`)
      .then(response => setData(response.data))
      .catch(error => setError(error.response.data));
  }, []);

  if (error) {
    return (
      <h1 className="text-center text-2xl mt-[10vh]">
        Не удалось загрузить статус арифметического выражения :(
      </h1>
    );
  }

  return (
    <div className="flex flex-col items-center mt-[10vh]">
      <h1 className="text-2xl">Арифметическое выражение № {data.id}</h1>
      <p className="text-xl my-2">
        Статус: <i>{data.status === "in-progress" ? "выполнение" : "завершено"}</i>
      </p>
      <p className="text-xl my-2">
        Результат: <i>{data.status === "finished" ? data.result : "дождитесь выполнения расчетов"}</i>
      </p>
    </div>
  );
};

export default ExpressionPage;