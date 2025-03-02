import {useState} from "react";
import axios from "axios";
import {useNavigate} from "react-router";
import {ReactComponent as LensSvg} from "./lens.svg";

const CalculatePage = () => {
  const navigate = useNavigate();
  const [expression, setExpression] = useState("");
  const [error, setError] = useState("");

  const onClick = () => axios.post("http://localhost:8080/api/v1/calculate", {expression})
    .then(response => navigate(`/expressions/${response.data.id}`))
    .catch(error => setError(error.response.data));

  return (
    <div className="flex flex-col items-center mt-[10vh]">
      <h1 className="text-3xl mt-6 mb-3">
        Что вам нужно посчитать?
      </h1>
      {error && <p className="text-red-500">{error}</p>}
      <div className="w-1/3 mt-20 flex flex-row items-center justify-center">
        <input
          type="text"
          className="w-2/3 outline-none bg-inherit border-black border-b-2"
          placeholder="Введите арифметическое выражение"
          value={expression}
          onChange={event => setExpression(event.target.value)}
        />
        <LensSvg
          className="w-1/6 h-8 hover:cursor-pointer hover:bg-amber-50 rounded"
          onClick={onClick}
        />
      </div>
    </div>
  );
};

export default CalculatePage;
