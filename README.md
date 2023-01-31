 ##  setting up hasura and graphql 
 check if wheather hasura and graphql instance alerady running in your docker if not 

 1.head to **\hasuraconfig**   folder and open **docker-compose.yaml** file make sure to change all **postgrespassword**
 field   with Your postgres database password
 
 2.make sure to set **HASURA_GRAPHQL_JWT_SECRET key** with any 32 character that you want
 
 3.make sure to set **HASURA_GRAPHQL_ADMIN_SECRET**  to your own hasura console password
 
 4.open cmd and change directory to **\hasuraconfig**
 
 5.type   **docker-compose up -d**    and run the command
 
 6.open hasura console in browser and head over to setting
 
 7.Import hasura metadata name **hasura_metadata_2023_01_29_13_30_51_667**  found in **\hasuraconfig** 
 
 8.now hasura is ready
 ## setting up golang and backend
 1.head to **\hahu_backend** folder
 
 2.open cmd and change directory to  **\hahu_backend**
 
 3.type  **go get**   and hit enter to install all required dependencies
 
 4.head to **\hahu_backend\helper** and open **jwt.go** file and make sure to change the string passed to this method 
 is your  **HASURA_GRAPHQL_JWT_SECRET key** that we set in **docker-compose.yaml** file
 
 **token.SignedString([]byte("HASURA_GRAPHQL_JWT_SECRET"))**
 
 5.type  **go run main.go**  in cmd and run command 
 
 6.your golang server start in port 8000

 ## setting up vue js and frontend
  1.open cmd and change directory to the root directory of clone github repository
  
  2.type **npm install** to  install all required dependencies
  
  3.type **npm run dev** to start your vue js and vite server





