import React from 'react'
import { BrowserRouter as Router, Route, } from 'react-router-dom'
import Register from './pages/Register'
import Login from './pages/Login'
import Session from './pages/Session'
import Dashboard from './pages/Dashboard'
import './App.css'

function App() {
  return (
    <Router>
      <Route exact path="/" component={Session} />
      <Route path="/register" component={Register} />
      <Route path="/login" component={Login} />
      <Route path="/session" component={Session} />
      <Route path="/dashboard" component={Dashboard} />
    </Router>
  )
}

export default App
