# EmergingTechnologiesProject
  We have created a single page, blogging application in Golang, Angular JS and MongoDB.
  The website allows for the basic CRUD operations in terms of blog posts, posting comments and user registration and login authentication
  
  Member of the group included:
  
  Pauric Boyle: [Pauric's Github](https://github.com/Pboyle5h)
  
  Eamon McNicholas: [Eamon's Github](https://github.com/DevEMCN)
  
  Alan Niemiec: [Alan's Github](https://github.com/sinderpl)
  
![solarized dualmode](https://github.com/Pboyle5h/EmergingTechnologiesProject/blob/master/ArchitectureDesign.png)
##Introduction 
The Project scope was to design a single page web application with a go server. For our front end we used AngularJS and styled it using bootstrap. We used MongoDB as our database and hosted the whole thing on heroku. The idea of our project was to design a blog that let the user log in, create a post, view a post and delete a post. WE did hit limitations within this project which can be read in our wiki and our read me file.

##Online Usage:
Our website is hosted on herohu at the following URL [Hosted Blog Page](https://goproject.herokuapp.com/)
Read the User Guide for instruction on how to interact with the web application. 

Heroku is designed to deploy from the master branch on git so every push to master made heroku spin up a new instance of our web application.
<br />
<br />
##Local installation
####[Install GIT on your computer] (https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
####[Install Golang on your computer] (https://golang.org/)
####[Clone the git repository] (https://github.com/Pboyle5h/EmergingTechnologiesProject)
Navigate to the EmergingTechnologiesProject repository folder.
####Install the gorilla toolkit with the following command:
```
go get github.com/gorilla/mux
```
####Run command:
```
go build App.go
```
####Make sure you have a working internet connection and run command:
```
go run App.go
```
####Open a web browser and go to:
Make sure port 4000 is not used by any other application.
```
localhost:4000
```
Fefer to the User Guide for further instructions.
<br />
<br />

##User Guide:
####Registration:
1. Click on the Register button in the top right corner.<br />
2. Input the data into the correct fields, in the correct format. (eg. email has to be in the "emailName"@"gmail.com" format)<br />
3. Click submit<br />
4. You will be transported to the Login page automatically.<br />
  **Limitation** The code does not currently validate whether a user with the same name already exists.

####Log in:
1. Click on the login button in the top right corner.<br />
2. Input your login details.<br />
3. Click submit and await response from the database.<br />

###Tabs:
####Blogs:
1. Click on the Blogs tab in the left corner.<br />
2. All the blogs that are currently in the database will be displayed.<br />
3. Click on the blog that you are interested in reading.<br />
  **Limitation** A comments function has been implemented in GO but no fully implemented on the HTML and Angular side.

####User Page:
1. Follow the Login Guide in order to make this tab available.<br />
2. Click on your username in the top right corner.<br />
3. A list of your personal posts will appear.<br />

####Reviewing your blogs:
1. Click on the blog you wish to read.<br />
2. Read your blog.<br />
  **Limitation** A comments function has been implemented in GO but no fully implemented on the HTML and Angular side.

####Add a new blog:
1. Click on the blue plus button in the top right corner.<br />
2. Insert a title.<br />
3. Insert the blog post. (Must be over 70 characters)<br />
4. Insert the author name.<br />
5. Review the blog details. The frame surrounding the text fields will turn green if the input is valid and red if it is not/ or field is empty<br />
6. Click the submit button at the bottom of the page.<br />
7. You will be redirected to the home page. Navigate back to the User Page in order to view your new post.<br />


####Table of Calls being made to the server:
API URL | HTTP Method | Request Data | Response Data 
------------ | ----------- | ----------- | -----------
"/" | GET | None. | index.html and all required resources.
"/blogs" | POST | Post Object and Session Cookie Object. | ThreadPost object that includes all posts.
"/blogs" | GET | Post Object and Session Cookie Object. | ThreadPost object that includes all posts.
"/register" | POST | Store Inputted Data to the databas | ???
"/login" | POST | Compares the data entered with data in the database| ???
"/logout" | POST | ???| ???
"/user" | POST | ???| ???
"/user" | PUT | ???| ???
"/user" | DELETE | Deletes session cookies realted to user| ???
"/user" | GET | Gets the user specific blogs from datatbse | user specific blog returned


