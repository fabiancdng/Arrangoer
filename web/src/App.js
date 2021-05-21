import { useState, useEffect } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import { ChakraProvider } from '@chakra-ui/react'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Header from './components/Header'

const App = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState({});

  useEffect(() => {
    fetch("/api/auth/get/user")
        .then(async (res) => {
            if(res.ok) {
                res = await res.json();
                console.log(res);
                setLoggedIn(true);
                setUser(res);
                return
            }
        });
  }, [])

  return (
    <BrowserRouter>
      {/* Chakra Provider für globale Design-States */}
      <ChakraProvider>

        <Header />

        <Switch>
          <Route exact path="/">
            {loggedIn ? <Redirect to="/dashboard" /> : <Login />}
          </Route>

          <Route path="/dashboard">
            {loggedIn ? <Dashboard user={user} /> : <Redirect to="/" />}
          </Route>

          <Route path="/login">
            {loggedIn ? <Redirect to="/dashboard" /> : <Login />}
          </Route>
        </Switch>

      </ChakraProvider>
    </BrowserRouter>
  );
}

export default App;
