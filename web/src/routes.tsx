import { Route, Routes } from "react-router-dom"
import App from "./app";
import Welcome from "./welcome";


const NasuRoutes = () => {
    return (
        <Routes>
            <Route path='/' element={<App/>}/>
            <Route path='/welcome' element={<Welcome />}/>
        </Routes>
    )
}

export default NasuRoutes;