# Steps to replicate

1. Initialize a .env file
Create a .env file in the root of your folder. Add the following key, values

```
GoogleClientID="YOUR_GOOGLE_CLIENT_ID"
GoogleClientSecret="YOUR GOOGLE_CLIENT_SECRET"
MONGO_USERNAME="MONGO_USERNAME"
MONGO_PASSWORD="MONGO_PASSWORD"
SIGNING_KEY="SIGNING_KEY_FOR_JWT"
```



### Tags
1. **1.0-google-login-mongo-integration**
- Google login
- Mongo integration
- Find if a document with the email of the user is present in the mongo collection
- If user is not present in the collection, create a document 

2. **1.1-skills-crud**
- Allow CRUD on Skills model

3. **1.2-jwt-authentication**
- Introduce middleware to create jwt token upon signing in
- Create Admin model
- Wrap skills path within the newly introduced middleware
- Check if the current user is an admin when performing CRUD on the skills path

4. **1.3-LecturerSkills-crud**
- Allow CRUD on LecturerSkill model
- Delete LecturerSkill if a particular skill is deleted