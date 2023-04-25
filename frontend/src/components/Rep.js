import BarraNav from "./BarraNav";
import { useState, useEffect } from "react";

const Rep = () => {
  const [fileContent, setFileContent] = useState("");
  const [salida, setSalida] = useState("");
  const [logeado, setLogeado] = useState("");
  const [imagen, setImagen] = useState("");
  const host = "52.14.255.35";
  // const host = "localhost";

  useEffect(() => {
    console.log("login: ", window.login);
    console.log("user: ", window.user);
    if (window.login) {
      setLogeado(`Actualmente esta logueado como ${window.user}.`);
    }
  }, [logeado]);

  const handleFileInputChange = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (event) => {
      const fileContent = event.target.result;
      setFileContent(fileContent);
    };

    reader.readAsText(file);
  };

  const analizador = (event) => {
    event.preventDefault();
    const url = `http://${host}:80/reportes`;
    const data = { exp: fileContent };
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
        console.log(res);
        setSalida(res.contenido);
        setImagen(res.datob64);
      });
  };

  return (
    <div className="main">
      <BarraNav name="Reportes" />
      <div className="d-flex justify-content-end">
        <h6 className="text-info log">{logeado}</h6>
      </div>
      <div className="d-flex justify-content-evenly">
        <form>
          <div className="form-group">
            <label htmlFor="formFile" className="form-label mt-4">
              Seleccione un script
            </label>
            <input
              className="form-control"
              type="file"
              id="formFile"
              onChange={handleFileInputChange}
            />
          </div>
        </form>
      </div>

      <div className="container text-center">
        <div className="row">
          <div className="col-6">
            <form onSubmit={analizador}>
              <fieldset>
                <div className="form-group">
                  <label htmlFor="exampleTextarea" className="form-label mt-4">
                    <h2>Consola</h2>
                  </label>
                  <textarea
                    className="form-control"
                    id="exampleTextarea"
                    rows="5"
                    defaultValue={fileContent}
                    onChange={(event) => setFileContent(event.target.value)}
                  ></textarea>
                </div>
                <br />
                <div className="container text-center">
                  <button type="submit" className="btn btn-success">
                    Analizar
                  </button>
                </div>
              </fieldset>
            </form>
          </div>
          <div className="col-6">
            <form>
              <fieldset>
                <div className="form-group">
                  <label htmlFor="exampleTextarea" className="form-label mt-4">
                    <h2>Salida</h2>
                  </label>
                  <textarea
                    disabled
                    className="form-control"
                    id="exampleTextarea"
                    rows="5"
                    defaultValue={salida}
                  ></textarea>
                </div>
              </fieldset>
            </form>
          </div>
        </div>
      </div>
      <br />
      <div className="text-center">
        <img src={imagen} alt="reporte"/>
      </div>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
    </div>
  );
};

export default Rep;
