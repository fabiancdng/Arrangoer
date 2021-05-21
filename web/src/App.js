import { useContext, useEffect } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Header from './components/Header'
import { UserContext } from './context/UserContext'

const App = () => {

  const { loggedIn, user, setLoggedIn, setUser } = useContext(UserContext)
    
  useEffect(() => {
    fetch("/api/auth/get/user")
      .then(res => {
        if(res.ok) {
          res = res.json()
          .then(res => {
              console.log(res);
              setLoggedIn(true);
              setUser(res);
          });
        }
      })
    }, [setLoggedIn, setUser])

  return (
    <BrowserRouter>
      <Header />
      <Switch>
        <Route exact path="/">
          {loggedIn ? <Redirect to="/dashboard" /> : <Redirect to="/login" /> }
        </Route>

        <Route path="/dashboard">
          {loggedIn ? <Dashboard /> : <Redirect to="/" />}
        </Route>

        <Route path="/login">
          {loggedIn ? <Redirect to="/dashboard" /> : <Login />}
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
