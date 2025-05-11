import React, {createContext, useContext, useEffect, useState} from "react";
import Api from "../../api";

const UserContext = createContext(null);

export const useUser = () => {
  const context = useContext(UserContext);

  if (!context) {
    throw new Error("useUser must be used within an UserProvider")
  }

  return context;
}

export const UserProvider = ({children}) => {
  const [user, setUser] = useState(null);

  const updateUser = () => {
    Api.get("/profile")
      .then(response => setUser(response.data))
      .catch(() => setUser(null));
  };

  useEffect(() => {
    updateUser();
  }, []);

  return (
    <UserContext.Provider value={{user, updateUser}}>
      {children}
    </UserContext.Provider>
  );
}