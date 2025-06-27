import React, { useState } from "react";
import logo from './logo.svg';
import './App.css';
import { AuthmasterClient } from "./proto/AuthmasterServiceClientPb"; 
import { LoginRequest, LoginResponse } from "./proto/authmaster_pb";

const ENVOY_URL = process.env.REACT_APP_ENVOY_URL ?? "http://localhost:8530"

const login = async (username: string, password: string): Promise<string> => {
  const client = new AuthmasterClient(ENVOY_URL);
  const request = new LoginRequest();
  request.setUsername(username);
  request.setPassword(password);
  return client.login(request, {})
    .then(resp => resp.getToken());
};

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");

  const onClickLogin = () => {
    login(username, password)
      .then(res => setToken(res))
      .catch(error => {
        console.log("failed to login: " + JSON.stringify(error))
      })
  };

  return (
    <div>
      <table >
        <thead>
          <tr>
            <th></th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>
              username
            </td>
            <td>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </td>
          </tr>
          <tr>
            <td>
              password
            </td>
            <td>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </td>
          </tr>
          <tr>
            <td>
            </td>
            <td>
              <button onClick={onClickLogin}>login</button>
            </td>
          </tr>
          <tr>
            <td>
              token
            </td>
            <td>
              <input
                type="text"
                value={token}
                readOnly
              />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}

export default App;
