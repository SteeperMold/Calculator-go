import {useEffect, useState} from "react";
import {useNavigate} from "react-router";
import axios from "axios";

const ExpressionsPage = () => {
  const [error, setError] = useState("");
  const [expressions, setExpressions] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    axios.get("http://localhost:8080/api/v1/expressions")
      .then(response => setExpressions(response.data.expressions))
      .catch(error => setError(error.response.data));
  }, []);

  if (error) {
    return (
      <h1 className="text-center text-2xl mt-[10vh]">
        Не удалось загрузить список арифметических выражений :(
      </h1>
    );
  }

  return (
    <div className="flex flex-col items-center">
      <table className="w-1/2 border-collapse border border-gray-200">
        <thead>
        <tr>
          <th className="border border-gray-200 px-4 py-2">ID</th>
          <th className="border border-gray-200 px-4 py-2">Статус</th>
          <th className="border border-gray-200 px-4 py-2">Результат</th>
        </tr>
        </thead>
        <tbody>
        {expressions.map((expression) => (
          <tr
            key={expression.id}
            className="hover:bg-amber-50 hover:cursor-pointer"
            onClick={() => navigate(`/expressions/${expression.id}`)}
          >
            <td className="border border-gray-200 px-4 py-2 text-center">
              № {expression.id}
            </td>
            <td className="border border-gray-200 px-4 py-2 text-center">
              <i>
                {expression.status === 'in-progress'
                  ? 'выполнение'
                  : 'завершено'}
              </i>
            </td>
            <td className="border border-gray-200 px-4 py-2 text-center">
              <i>
                {expression.status === 'finished'
                  ? expression.result
                  : 'дождитесь выполнения расчетов'}
              </i>
            </td>
          </tr>
        ))}
        </tbody>
      </table>
    </div>
  );
};

export default ExpressionsPage;

