# About

**What it is?**

A project inspired by W2G and Rabb.it streaming websites, that allows users to watch movies, series & videos together remotely.

**How to run**

*First step*

- Clone the repository

*Second step*

- Acess 'front' folder, open a termninal and type `npm i`

*Third step* 

- Type `npm run dev` in 'front' terminal directory to run front, the post will be listening in `http://localhost:3000`

*Fourth step*

- Acess the back folder and create a file called `config.toml` in root's project and fill out with your informations:

```
[api]
port=":4000"

[db]
host="localhost"
port="5432"
user=""
pass=""
name=""

[jwt]
secretkey=""
issure=""
```

*Fifth step* 

- Type `go run main.go` in root directory and acess the API in `http://localhost:4000`



