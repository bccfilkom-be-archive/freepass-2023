# BCC University

### ⚠️⚠️⚠️
```
Submissions from 2022 students will have much higher priority than submissions from 2021, SAP, or higher students.
Please take note of this before planning to attempt this freepass challenge.
```
## :love_letter: Invitation Letter

In this digital age, transparency and management of information is greatly needed, especially in the field of education. At some point, there is an investor who establishes a university, while the university requires a system for managing information from its students.

This application is intended for university students to be able to access information provided by the university. To accomplish this quickly, BCC will need help. By this letter, we humbly invite you to join us on this journey to figure out the best solution. We cannot wait to see your ideas to overcome this problem.

## :star: Minimum Viable Product (MVP)

As we have mentioned earlier, we need technology that can support BCC University in the future. Please consider these features below:

* A new user can register account to the system ✔
* A new user can login to the system ✔
* User can edit their profile account ✔
* User can view a list of classes ✔
* User can add class and has maximum of 24 sks. ✔
* User can drop class ✔
* User can view participants of the classes they have taken ✔
* Admin can remove user from class ✔
* Admin can add user to class ✔
* Admin can get all class which has a difference course ✔
* Admin can create new class with specific course ✔
* Admin can edit the name of the class and the type of course that is owned ✔
* Admin can delete the class ✔

## :earth_americas: Service Implementation

```text
GIVEN => I am a new user
WHEN => I register to the system
THEN => System will record and return the visitor's username

GIVEN => I am a user
WHEN => I took an action to edit my account
THEN => System will show a "successfully edited" notification

GIVEN => I am a user
WHEN => I took an action to see my account
THEN => System will show the user's profile

GIVEN => I am a user
WHEN => I took an action to view all class
THEN => System will show all classes

GIVEN => I am a user
WHEN => I took an action to add new class
THEN => System will show a "successfully added new class" notification

GIVEN => I am a user
WHEN => I took an action to drop class
THEN => System will show a "successfully dropped a class" notification

GIVEN => I am a user
WHEN => I took an action to view participants of the classes they have taken
THEN => System will show all participants from the class

GIVEN => I am an admin
WHEN => I took an action to delete an user from class
THEN => System will show a "successfully deleted an user" notification

GIVEN => I am an admin
WHEN => I took an action to add an user to class
THEN => System will show a "successfully added new user" notification

GIVEN => I am an admin
WHEN => I took an action to delete a class
THEN => System will show a "successfully deleted" notification

GIVEN => I am an admin
WHEN => I took an action to create a class
THEN => System will record and return the class identity number

GIVEN => I am an admin
WHEN => I took an action to edit a class
THEN => System will show a "successfully edited" notification

GIVEN => I am an admin
WHEN => I took an action to view all classes
THEN => System will show all classes
```

## :family: Entities and Actors

We want to see your perspective about these problems. You can define various types of entities or actors. One thing for sure, there is no true or false statement to define the entities. As long as the results are understandable, then go for it! :rocket:

## :blue_book: References

You might be overwhelmed by these requirements. Don't worry, here's a list of some tools that you could use (it's not required to use all of them nor any of them):

1. [Example Project](https://github.com/meong1234/fintech)
2. [Git](https://try.github.io/)
3. [Cheatsheets](https://devhints.io/)
4. [REST API](https://restfulapi.net/)
5. [Insomnia REST Client](https://insomnia.rest/)
6. [Test-Driven Development](https://www.freecodecamp.org/news/test-driven-development-what-it-is-and-what-it-is-not-41fa6bca02a2/)
7. [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
8. [GraphQL](https://graphql.org/)
9. [gRPC](https://grpc.io/)
10. [Docker Compose](https://docs.docker.com/compose/install/)

## :hocho: Accepted Weapons

> **BEFORE CHOOSING YOUR LANGUAGE, PLEASE VISIT OUR [CONVENTION](CONVENTION.md) ON THIS PROJECT**
>
> **Any code that did not follow the convention will be rejected!**

1. Golang (preferred)
2. NodeJS
3. PHP
4. Java

You are welcome to use any libraries or frameworks, but we appreciate if you use the popular ones.

## :school_satchel: Tasks
```
The implementation of this project MUST be in the form of a REST, gRPC, or GraphQL API (choose AT LEAST one type).
```
1. Fork this repository
2. Follow the project convention
3. Finish all service implementations
4. Write the installation guide of your backend service in the section below

## :test_tube: API Installation
1. Use Go version 1.19 or above
2. Make sure you have created an empty database
3. Rename .env.example to .env and fill the necessary fields
4. go run main.go

https://documenter.getpostman.com/view/22317100/2s935iuSTa

## :gift: Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https://media.giphy.com/media/kv5fbxHVAEOjrHeCLk/giphy.gif)

> **This is *not* the only way to join us.**
>
> **But, this is the *one and only way* to instantly pass.**
