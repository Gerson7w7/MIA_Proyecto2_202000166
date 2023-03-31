import { Link } from "react-router-dom";

const BarraNav = (props) => {
  return (
    <div className="main">
      <nav className="navbar navbar-expand-lg navbar-dark bg-primary">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">
            <img
              alt=""
              className="logo-prin"
              src="https://freepngimg.com/thumb/symbol/63241-console-command-line-icons-terminal-computer-linux-interface.png"
            />
          </Link>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarColor01"
            aria-controls="navbarColor01"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarColor01">
            <ul className="navbar-nav me-auto">
              <li className="nav-item">
                <br/>
                  <h3 className="text-secondary titulo">Proyecto 2</h3>
              </li>
              <li className="nav-item">
                <Link className="nav-link active" to="/">
                  <img
                    alt=""
                    className="logo"
                    src="https://cdn.pixabay.com/photo/2013/07/13/13/41/bash-161382_960_720.png"
                  />
                  <h6 className="text-secondary">Inicio</h6>
                  <span className="visually-hidden">(current)</span>
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/login">
                  <img
                    alt=""
                    className="logo"
                    src="https://cdn-icons-png.flaticon.com/512/295/295128.png"
                  />
                  <h6 className="text-secondary">Login</h6>
                </Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/reportes">
                  <img
                    alt=""
                    className="logo"
                    src="https://upload.wikimedia.org/wikipedia/commons/9/94/Report.png"
                  />
                  <h6 className="text-secondary">Reportes</h6>
                </Link>
              </li>
            </ul>
            <br/>
            <h1 className="text-warning titulo">{props.name}</h1>
          </div>
        </div>
      </nav>
    </div>
  );
};

export default BarraNav;