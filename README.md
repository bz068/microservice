# Go-Node Microservices
---

This project consists of two microservices, first `Posts` microservice, which is implemented in Go and the the second microservice `Comments` which is implemented in Node JS.

`Posts` microservice communicates with the `Comments` microservice through HTTP requests, allowing client to use data from both services by making a HTTP request to just `Posts` microservice.

------------------------
## `Posts` Microservice
Implemented in Go.
> Runs on PORT `8080`
-----

### Routes

`/` ---> GET
> index handler - just welcomes user to API.

`/api/v1/allposts` ----> GET
> retrieves all posts.

Response JSON:
```
{
    "data": [
        {
            "id": 1,
            "title": "title",
            "body": "body",
            "date": "2022-04-06T17:09:54.871369+01:00",
            "comments": [
                {
                    "id": 23,
                    "postId": 1,
                    "comment": "first comment"
                }
            ]
        }
    ],
    "message": "Posts Fetched Successfully",
    "success": true
}
```

`/api/v1/addpost` -----> POST
> send post to service.

Request Body:
```
{
    "title": "title",
    "body": "body"
}
```
Date Types:
> title = string

> body = string
------------------------
## `Comments` Microservice
Implemented in Node JS.
> Runs on PORT `9000`
-----
### Routes

`/api/v1/comments` ----> POST
> send comment to `Comments` service

```
{
    "postId": 2432,
    "comment": "comment"
}
```
Date Types:
> postId = int

> comment = string

`/api/v1/posts/:id/comments` ----> GET
> get all comments for post by `postId`

```
{
    "data": [
        {
            "id": 23,
            "postId": 1,
            "comment": "first comment"
        }
    ],
    "msg": "Comment Fetched Successfully",
    "success": true
}
```
---
