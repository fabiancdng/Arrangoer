import { useContext, useEffect } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Header from './components/Header'
import { UserContext } from './context/UserContext'

const App = () => {

  const userContext = useContext(UserContext)

  useEffect(() => {
    fetch("/api/auth/get/user")
        .then(async (res) => {
            if(res.ok) {
                res = await res.json();
                console.log(res);
                userContext.setLoggedIn(true);
                userContext.setUser(res);
                return
            }
        });
  }, [userContext])

  return (
    <BrowserRouter>
      <Header />
      <Switch>
        <Route exact path="/">
          {userContext.loggedIn ? <Redirect to="/dashboard" /> : <Login />}
        </Route>

        <Route path="/dashboard">
          {userContext.loggedIn ? <Dashboard /> : <Redirect to="/" />}
        </Route>

        <Route path="/login">
          {userContext.loggedIn ? <Redirect to="/dashboard" /> : <Login />}
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
