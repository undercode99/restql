# restql


## Instalation

1. Build binnary go
   - `git clone https://github.com/undercode99/restql.git`
   - `cd cmd/restql`
   - `go build`

2. Create an .env file for application and database configuration containing
```
DB_USERNAME=<DBUSERNAME>
DB_PASSWORD=<DBPAWSSWORD>
DB_HOST=<DBHOST>
DB_PORT=<DBPORT>
DB_NAME=<DBNAME>
APP_PORT=8863
```
Note: only support postgresql

3. Running file binary `restql` for Running app in the same folder location as the .env file with command 
 ```./restql```



## Usage

#### Example Request API :

Endpoint :  `http://localhost:8863/api/v1/query`

Method   :  `GET`

Header   : `'Content-Type: application/json'`

Body     : `
    {
    "query": "SELECT id, name FROM users;"
    }
    `

#### Example Response :
```
{
    "columns": {
        "id": "INT8",
        "name": "VARCHAR"
    },
    "data": [
        {
            "id": 1,
            "name": "Admin"
        },
        {
            "id": 3,
            "name": "User"
        }
    ],
    "sql": "SELECT id, name FROM users;"
}
