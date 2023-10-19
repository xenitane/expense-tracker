# Expense Tracker Apllication

This is an Expense Tracking Application written in Golang and React with MongoDB as it's datastore. Use it to track your earnings and expenses with a detailed dashboard showing the pattern of these events and take insights on how to manage this.

## Setup
- Server:
   1. Create the `.env` file in the server folder with the following keys:
      - `PORT` : Port number for the backend server
      - `MONGOURI` : Address for the MongoDB datastore to be used
      - `DBNAME` : The name of the database to work with
- Client:
  1. Create the `.env.json` file in the client folder with the key `server` with these server options:
     - `cors` : allowing cors
     - `proxy` : proxy options for `api`, [refer](https://vitejs.dev/config/server-options.html#server-proxy).

### **🛠 &nbsp;Tech Stack**
![Visual Studio Code](https://img.shields.io/badge/visual_studio_code-%23007ACC.svg?style=for-the-badge&logo=visualstudiocode&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23F7DF1E.svg?style=for-the-badge&logo=javascript&logoColor=white)
![NodeJS](https://img.shields.io/badge/node.js-%23339933.svg?style=for-the-badge&logo=nodedotjs&logoColor=white)
![React](https://img.shields.io/badge/react-%2361DAFB.svg?style=for-the-badge&logo=react&logoColor=white)
![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![GoFiber](https://img.shields.io/badge/GoFiber-%2300ADD8.svg?style=for-the-badge&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CCjxzdmcgdmlld0JveD0iMCAwIDgwMCAzMDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI%2BPGcgZmlsbD0iI2ZmZiI%2BPHBhdGggZD0ibTIzOS4yNSA2OS41OGgxMTcuODRsLTUuNzQgMzMuODloLTc2LjExbC01LjczIDM1LjRoNjguNjFsLTUuNzMgMzRoLTY4LjYxbC0xMS40NyA2OS4zN2gtNDEuNzJ6Ii8%2BPHBhdGggZD0ibTM1OSAxMTIuNzRoNDEuMjJsLTIxLjU0IDEyOS40N2gtNDEuMjJ6bTUtMzQuODljMS43LTEwLjg1IDEyLjQxLTE5LjY0IDI0LTE5LjY0czE5Ljg1IDguNzkgMTguNDEgMTkuNjQtMTIuMjIgMTkuNzItMjMuODUgMTkuNzItMjAuMTQtOC45NC0xOC41Ni0xOS43MnoiLz48cGF0aCBkPSJtNDIxLjYxIDY5LjU4aDQxLjIxbC0xMC43IDY1LjMzaC44OGM2Ljc0LTExLjg5IDIwLjE1LTIzLjg1IDQwLjYzLTIzLjg1IDI3LjA2IDAgNDggMjAuODIgNDAuMzcgNjYuNS03LjI1IDQ0LjA4LTM0LjIyIDY2LjUtNjIuODggNjYuNS0xOS41NSAwLTI5LjMzLTEwLjg3LTMyLjM2LTIyLjg0aC0xLjYxbC0zLjM3IDIxaC00MC43OHptNDAuNzEgMTQyLjQyYzE1LjE3IDAgMjYtMTMuNTcgMjkuMzMtMzQuNDdzLTIuNy0zNC4yMy0xOC0zNC4yM2MtMTUuMDkgMC0yNi4xMyAxMy4wNy0yOS41IDM0LjIzLTMuNDEgMjAuOTMgMy4wOCAzNC40NyAxOC4xNyAzNC40N3oiLz48cGF0aCBkPSJtNTM4Ljc5IDE3OGM2LjY1LTQwLjEyIDM2LjMyLTY2LjkyIDc1LjM1LTY2LjkyIDM2Ljc1IDAgNTguNjcgMjMuNTEgNTEuNTkgNjYuMDhsLTEuNjEgMTAuMjhoLTg2LjIybC0uMjYgMS41MmMtMi42MSAxNi40NCA1LjE0IDI2IDIxLjI0IDI2IDEwLjcxIDAgMjAtNC41NSAyNC43LTEzLjIzbDM3LjU5IDEuMDljLTkuMTggMjUuNDYtMzMuODggNDEuODktNjggNDEuODktNDAuNjUtLjA2LTYxLjM4LTI1LjM0LTU0LjM4LTY2Ljcxem05MC44Ni0xNC41YzIuMTEtMTMuMjMtNS44Mi0yMi42Ny0xOS4zLTIyLjY3LTEzLjMyIDAtMjUuMTIgOS44Ni0yOC4wNyAyMi42N3oiLz48cGF0aCBkPSJtNjg1LjggMTEyLjc0aDQwbC00IDIzLjZoMS4zNWM3LjY3LTE3LjExIDIwLjM5LTI1LjM0IDM0LjU1LTI1LjM0YTQ5IDQ5IDAgMCAxIDExLjMgMS40bC02LjA3IDM1LjkxYy0zLjc5LTEuNDMtMTAuNzktMi4xOS0xNS45My0yLjE5LTE0LjU4IDAtMjcuMjMgMTAuMzctMjkuOTIgMjUuNzlsLTExLjY0IDcwLjNoLTQxLjIxeiIvPjwvZz48cGF0aCBkPSJtMjEzLjc3IDExNS40MWgtMTM0LjUxbC0xMC43NiAxNS4yNyAxNDIuNzgtLjA4eiIgZmlsbD0iI2ZmZiIvPjxwYXRoIGQ9Im0yMDkuMDIgMTQ3LjA2aC0xNzQuMDdsLTExLjA3IDE1LjIxIDE4Mi42NS0uMTN6IiBmaWxsPSIjZmZmIi8%2BPHBhdGggZD0ibTIwNC4yNyAxNzguN2gtNzQuMzdsLTEwIDE1LjI3IDgxLjg4LS4wNnoiIGZpbGw9IiNmZmYiLz48L3N2Zz4K&logoColor=white)
![MongoDB](https://img.shields.io/badge/mongodb-%2347A248.svg?style=for-the-badge&logo=mongodb&logoColor=white)