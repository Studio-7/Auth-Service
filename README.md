## APIDOC
**Endpoint** - '/user/signup'  
Required  
- *username* string
- *password* string
- *fname* string: First name
- *lname* string: Last name
- *email* string  
Response: Returns a token if successful else returns error.
```json
{"error" : "ERROR"}
```
  
**Endpoint** - '/user/login'  
Required  
- *username* string
- *password* string  
Response: Returns a token is successful else returns error. 
```json
{"error" : "ERROR"}
``` 
  
**Endpoint** - '/user/followuser'  
Required  
- *token* string: A valid JSON Web Token. A JWT is only valid for one API call.  
- *username* string
- *followee* string: Username of the user to follow  
Response: Returns  on success
```json
{ "result": "success", "token" : "<NEWTOKEN>" }
```
else returns error.
```json
{"error" : "ERROR"}
```