import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Calendar from './component/Calendar'
// import 'react-calendar/dist/Calendar.css';
import Navbar from './component/Navbar';

function App() {
  

  return (
    <>
    <Navbar/>
    <Calendar />

    </>
  )
}

export default App
