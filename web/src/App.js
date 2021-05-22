import { useContext, useEffect } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import Login from './pages/Login'
import AdminDashboard from './components/AdminDashboard'
import MemberDashboard from './components/MemberDashboard'
import Header from './components/Header'
import { UserContext } from './context/UserContext'

const App = () => {

  const { guild, loggedIn, setLoggedIn, setUser, setGuild } = useContext(UserContext)
    
  useEffect(() => {
    fetch("/api/auth/get/user")
      .then(async res => {
        if(res.ok) {
          res = await res.json()
          console.log(res);
          setLoggedIn(true);
          setUser(res);
        }
      })

      fetch("/api/auth/get/guild")
        .then(async res => {
          if(res.ok) {
            res = await res.json();
            console.log(res);
            setGuild(res)
          } else if(res.status === 401) {
            setLoggedIn(false);
            setGuild(false);
          }
        });
    }, [setLoggedIn, setUser, setGuild])

  if(loggedIn === "pending" || guild === "pending") {
    return null
  }

  return (
    <BrowserRouter>
      <Header />
      <Switch>
        <Route exact path="/">
          {
            loggedIn
              ? <Redirect to="/dashboard" />
              : <Redirect to="/login" />
          }
        </Route>

        <Route exact path="/dashboard">
          {
            !loggedIn
            ? <Redirect to="/login" /> 
            : guild.user_is_admin
            ? <Redirect to="/dashboard/admin" />
            : <MemberDashboard />
          }
        </Route>

        <Route path="/dashboard/admin">
          {
            !loggedIn
            ? <Redirect to="/login" /> 
            : !guild.user_is_admin
            ? <Redirect to="/dashboard" />
            : <AdminDashboard />
          }
        </Route>

        <Route path="/login">
          {
            loggedIn 
            ? <Redirect to="/dashboard" />
            : <Login />
          }
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
