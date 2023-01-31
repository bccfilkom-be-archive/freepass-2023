## DISCLAIMER

Because of my lack of skill and inexperience, this project unfortunatelly wont make it in time with the deadline `January 31th, 2023 23:59:59`. Therefore i would like to apologize to anyone reading this and especially to bcc_university commitee 🙏. Even thought it will be overdue the deadline, it does not mean i wont finish it anytime soon. I will update this repository when its done, not just because i want it to be finished, but personally i do have high interest in this particular part of Software Engineering, therefore thank you all for the oppurtunity. Best Regards, Mirza `January 31th, 2023 22:40:36`

# Problem Statement

In this digital age, transparency and management of information is greatly needed, especially in the field of education. At some point, there is an investor who establishes a university, while the university requires a system for managing information from its students.

This application is intended for university students to be able to access information provided by the university. Designed and developed from scratch, this API was built using Web Framework [GIN](https://github.com/gin-gonic/gin) and ORM library called [GORM](https://gorm.io/) which all of it was written in [Go](https://github.com/golang/go) Programming Language.

## Prerequisites

To run this repository on your local machine, you will need to install:

- [Docker](https://www.docker.com/)

## API User

### `POST /user/register`

```
{
    "student_id": "225150707111067",
    "fullname": "Sandy Cheekz"
    "email": "sandycheekz@student.ub.ac.id",
    "password": "texas"
}
```

##### Response

```
{
  "msg": "user successfuly registered"
}
```

---

### `POST /user/login`

```
{
    "fullname": "Sandy Cheekz",
    "email": "sandycheekz@student.ub.ac.id"
}
```

##### Response

```
{
  "msg": "login succes, here's your token 12673xv4536257"
}
```

---

### `PUT /user/profile/edit`

```
{
	"student_id": "2251507071110670",
    "email": "cheekzsandy@ub.ac.id",
    "enrolled_class": {
        "274863583n7c4y7": "Sistem Basis Data",
        "712bc354627it1c": "Bahasa Indonesia"
    }
}
```

---

Oops, looks like the page already ended. Please wait for the Author to update the repository!
