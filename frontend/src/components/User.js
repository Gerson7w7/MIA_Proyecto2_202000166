import BarraNav from "./BarraNav";
import { useState, useEffect } from "react";

window.user = "";
window.login = false;

const Inicio = () => {
  const [idPart, setIdPart] = useState("");
  const [user, setUser] = useState("");
  const [pwd, setPwd] = useState("");
  const [salida, setSalida] = useState("");
  const [logeado, setLogeado] = useState("");

  useEffect(() => {
    console.log("login: ", window.login);
    console.log("user: ", window.user);
    if (window.login) {
      setLogeado(`Actualmente esta logueado como ${window.user}.`);
    } else {
      setLogeado("Actualmente no hay ninguna sesion iniciada.");
    }
  }, []);

  const handleSubmit = (event) => {
    event.preventDefault();
    const url = "http://localhost:5000/login";
    const data = { usuario: user, clave: pwd, id: idPart };
    console.log(data);
    fetch(url, {
      method: "POST", // or 'PUT'
      body: JSON.stringify(data), // data can be `string` or {object}!
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .catch((error) => console.error("Error:", error))
      .then((res) => {
        const user = res.name 
        const pwd = res.pwd 
        if (!window.login) {
          if (user !== "" && pwd !== "") {
            setSalida(`Bienvenido ${user}! :D`);
            window.user = user;
            window.login = true;
            setLogeado(`Actualmente esta logueado como ${window.user}.`);
          } else {
            setSalida(
              "Usuario no encontrado, para crear usuarios primero ingrese como superusuario."
            );
          }
        } else {
          setSalida('Error. Primero cierra la sesion actual.')
        }
      });
  };

  const cerrarSesion = (event) => {
    event.preventDefault();
    if (window.login) {
      setLogeado("Sesion cerrada correctamente! :D");
      window.login = false;
      window.user = "";
      setSalida("");
    } else {
      setLogeado("Error. No hay sesion iniciada.");
    }
  };

  return (
    <div className="main">
      <BarraNav name="Login" />
      <br />
      <div className="d-flex justify-content-end">
        <h6 className="text-info log">{logeado}</h6>
        <form onSubmit={cerrarSesion}>
          <fieldset>
            <button type="submit" className="btn btn-warning">
              Cerrar Sesion
            </button>
          </fieldset>
        </form>
      </div>
      <div className="d-flex justify-content-evenly">
        <img
          alt=""
          className="logo-cont"
          src="https://cdn-icons-png.flaticon.com/512/272/272354.png"
        />
      </div>
      <div>
        <form onSubmit={handleSubmit}>
          <fieldset>
            <div className="d-flex justify-content-evenly">
              <div className="form-group">
                <label htmlFor="particion" className="form-label mt-4">
                  <h2>ID Particion</h2>
                </label>
                <input
                  type="text"
                  className="form-control"
                  id="particion"
                  aria-describedby="emailHelp"
                  placeholder="ejemplo: 661a"
                  onChange={(event) => setIdPart(event.target.value)}
                />
              </div>
            </div>
            <div className="d-flex justify-content-evenly">
              <div className="form-group">
                <label htmlFor="usuario" className="form-label mt-4">
                  <h2>Usuario</h2>
                </label>
                <input
                  type="text"
                  className="form-control"
                  id="usuario"
                  aria-describedby="emailHelp"
                  placeholder="root"
                  onChange={(event) => setUser(event.target.value)}
                />
              </div>
            </div>
            <div className="d-flex justify-content-evenly">
              <div className="form-group">
                <label htmlFor="pwd" className="form-label mt-4">
                  <h2>Password</h2>
                </label>
                <input
                  type="password"
                  className="form-control"
                  id="pwd"
                  placeholder="ingrese su contrasenia"
                  onChange={(event) => setPwd(event.target.value)}
                />
                <small id="emailHelp" className="form-text text-muted">
                  No comparta su contraseña con nadie.
                </small>
              </div>
            </div>
            <br />
            <div className="d-flex justify-content-evenly">
              <button type="submit" className="btn btn-success">
                Ingresar
              </button>
            </div>
          </fieldset>
        </form>
      </div>
      <br />
      <div className="d-flex justify-content-evenly">
        <h4 className="text-success">{salida}</h4>
      </div>
    </div>
  );
};

export default Inicio;
