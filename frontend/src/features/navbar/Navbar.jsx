import {Link} from "react-router";
import {useUser} from "../../shared/hooks/useUser";

const Navbar = () => {
  const {user} = useUser();

  return (
    <nav className="flex justify-between py-[2.5vh]">
      <div>
        <Link to="/" className="text-[3vh] ml-20">Главная</Link>
      </div>

      <div className="mr-10">
        {user ? (
          <Link to="/logout">{user?.login}</Link>
        ) : <>
          <Link to="/login" className="mr-10">Войти</Link>
          <Link to="/signup">Зарегистрироваться</Link>
        </>}
      </div>
    </nav>
  );
};

export default Navbar;
