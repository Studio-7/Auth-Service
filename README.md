## APIDOC
**Endpoint** - '/user/signup'  
Required  
- *username* string
- *password* string
- *fname* string: First name
- *lname* string: Last name
- *email* string  
- *googleauth* string: set to true (string) is google auth is used  
Response: Returns a token if successful else returns error.
```json
Success
{
    "result": "success",
    "token": "eWFrb21ha2F6YTE=.dENzTEFoREZrd09QbHdOTE1nYldQWWF2a21UQ3BiS2o=.N3JrYzVvZHlZNTBrMzJoU1ROWnpBdUxldmptLzFESytoUURqVmJZcWMvTT0="
}
Error
{
    "error" : "ERROR",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
```
  
**Endpoint** - '/user/login'  
Required  
- *username* string
- *password* string  
- *googleauth* string: set to true (string) is google auth is used  
Response: Returns a token is successful else returns error. 
```json
Success
{
    "result": "success",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
Error
{
    "error" : "ERROR",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
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
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
Error
{
    "error" : "ERROR",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
```

**Endpoint** - '/user/unfollowuser'  
Required  
- *token* string: A valid JSON Web Token. A JWT is only valid for one API call.  
- *username* string
- *followee* string: Username of the user to unfollow  
Response: Returns  on success
```json
Success
{
    "result": "success",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
Error
{
    "error" : "ERROR",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
```

**Endpoint** - '/user/getprofile'  
Required  
- *token* string: A valid JSON Web Token. A JWT is only valid for one API call.  
- *username* string
- *profilename* string: Username of the user whose profile is needed
Response: Returns  on success
```json
Success
{
    "result": {
        "FName": "thor",
        "LName": "odinson",
        "UName": "thor",
        "Email": "test@test.com",
        "ProfilePic": {
            "Link": "https://i.pinimg.com/originals/bd/93/3a/bd933a1b384c808326201c104c24ee6d.png",
            "CreatedOn": "",
            "CreatedBy": "",
            "UploadedOn": "0001-01-01T00:00:00Z",
            "Manufacturer": "",
            "Model": ""
        },
        "Followers": [
            "thor"
        ],
        "Following": [
            "ac491",
            "divya21raj"
        ],
        "FollowersCount": 1,
        "FollowingCount": 2,
        "TravelCapsules": [
            {
                "Id": "5fdda452-57d5-42c1-9bf4-b22e05b115af",
                "Title": "Postcard",
                "Posts": [],
                "CreatedOn": "2019-04-25T17:17:22.371Z",
                "CreatedBy": "thor",
                "UpdatedOn": "2019-04-25T17:17:22.371Z",
                "Hashtags": [],
                "Likes": 0,
                "ProfilePic": "https://i.pinimg.com/originals/bd/93/3a/bd933a1b384c808326201c104c24ee6d.png"
            },
            {
                "Id": "6118e1b5-e091-48a3-acc0-66a35ab9d515",
                "Title": "Postcard 2",
                "Posts": [
                    "f0afc238-ff98-4936-8001-fb7222e1376c",
                    "ad275790-43bd-44ab-9942-8a7c2edfa507",
                    "e3d01660-2980-46fc-ae23-5c2c6546161f",
                    "48700a85-cde9-494c-ba96-7d2ef3a1733b",
                    "1aab8b4e-f2ee-46ed-8a25-e160615c81d4"
                ],
                "CreatedOn": "2019-04-23T17:16:55.177Z",
                "CreatedBy": "thor",
                "UpdatedOn": "2019-04-25T17:10:53.18Z",
                "Hashtags": [],
                "Likes": 0,
                "ProfilePic": ""
            }
        ],
        "Images": [
            "https://ipfs.io/ipfs/QmcteUnUzoPQ52NaCfjEpEQk5CSBsHdW7bJJNi8Yt1KY8y"
        ]
    },
    "token": "dGhvcg==.Y0VrWEJBa2pRWkxDdFRNdFRDb2FOYXR5eWlOS0FSZUs=.UGJ5anFqWDExRDYydm9NamZxazBxbjF6ZkdseFdGTlN2MzU1SEJDQkEwUT0="
}
Error
{
    "error" : "ERROR",
    "token": "eWFrb21ha2F6YTE=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.ekltc1hQa2JTdWdQc0luYSs2cUpIaTlJZ2tOQ2FpbFJJdE5SQ3ovQnR0UT0="
}
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