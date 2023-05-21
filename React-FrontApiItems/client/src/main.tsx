import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import { Provider } from "react-redux";
import './index.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'react-redux-toastr/lib/css/react-redux-toastr.min.css'
import { BrowserRouter } from 'react-router-dom'
import { store } from './redux/store.ts';
import ReduxToastr from 'react-redux-toastr'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <BrowserRouter>
      <Provider store={store}>
        <>
          <App />
          <ReduxToastr
            timeOut={4000}
            newestOnTop={false}
            preventDuplicates
            transitionIn="fadeIn"
            transitionOut="fadeOut"
            progressBar
            closeOnToastrClick/>
        </>
      </Provider>
    </BrowserRouter>
  </React.StrictMode>,
)
