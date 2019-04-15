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
Success
{
    "result": "success",
    "jwt": "eWFrb21ha2F6YTE=.dENzTEFoREZrd09QbHdOTE1nYldQWWF2a21UQ3BiS2o=.N3JrYzVvZHlZNTBrMzJoU1ROWnpBdUxldmptLzFESytoUURqVmJZcWMvTT0="
}
Error
{"error" : "ERROR"}
```
  
**Endpoint** - '/user/login'  
Required  
- *username* string
- *password* string  
Response: Returns a token is successful else returns error. 
```json
Success
{
    "result": "success",
    "jwt": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
Error
{"error" : "ERROR"}
``` 
  
**Endpoint** - '/user/followuser'  
Required  
- *token* string: A valid JSON Web Token. A JWT is only valid for one API call.  
- *username* string
- *followee* string: Username of the user to follow  
Response: Returns  on success
```json
Success
{
    "result": "success",
    "jwt": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
Error
{"error" : "ERROR"}
```
  
**Endpoint** - '/user/updatedetails'  
Required 
- *username* string: Username of the user for which details are to be updated
- *password* string: Password of the user for which details are to be updated
- *newUsername* string (optional): New username for the user
- *newPassword* string (optional): New password for the user
- *fname* string (optional): New First Name of the user
- *lname* string (optional): New Last Name of the user
- *email* string (optional): New E-mail of the user

Response: Returns  on success
```json
{ "result": "success" }
```
else returns error.
```json
{"error" : "ERROR"}
```