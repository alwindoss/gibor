import styles from './App.module.css'
import { NavLink, Outlet } from 'react-router'

function App() {

  return (
    <>
      <div>
        <div className={styles.header}>Welcome to Gibor</div>
        <nav>
          <ul>
            <li>
              <NavLink to="/">Home</NavLink>
            </li>
          </ul>
          <ul>
            <li>
              <NavLink to="/about">About</NavLink>
            </li>
          </ul>
        </nav>
        <Outlet />
      </div>
    </>
  )
}

export default App
