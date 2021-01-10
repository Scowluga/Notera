# Notera Server

## AWS Setup 
- Setup a Postgres DB using [Amazon RDS for PostgreSQL](https://aws.amazon.com/rds/postgresql/). Ensure that the "Public accessibility" flag is set to true. 

- Setup a Redis cluster using [Amazon ElastiCache for Redis](https://aws.amazon.com/elasticache/redis/). 

- Setup an [AWS Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/) application. Upload the `Dockerrun.aws.json` file, which will pull the notera image directly from [Docker Hub](https://hub.docker.com/repository/docker/scowluga/notera/general). 

- Configure a security group for your application and setup the inbound rules to connect your application. 
  - Add rules with type PostgreSQL, and for source select the security group(s) of your elastic beanstalk server. 
  - Add a rule for Custom TCP and enter your Redis port for port range (default 6379). Then select `0.0.0.0/0`. Note that this seems unsafe, but does not actually publish your cluster to the internet. See [documentation](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/accessing-elasticache.html) on Amazon ElastiCache for more information. 
- Add any additional configuration to Elastic Beanstalk, and start hitting endpoints. 

## Local Setup
- Download and install 
  1. [Golang](https://golang.org/doc/install)
  2. [PostgreSQL](https://www.postgresql.org/download/)
  3. [Redis](https://redis.io/)
  4. [Docker](https://www.docker.com/)

- Clone the repository, and configure `.env` variables accordingly. 

- Start the container with `docker-compose up`

## API Documentation 
The REST API currently consists of a single CRUD module for notes. 

### `POST /notes`
This endpoint will create new notes and return the created note with ID.  

#### Request Body

```
{
  "MediaID": "media_",
  "UserID": "user_id",
  "Timestamp": 0,
  "Text": "hello world"
}
```

#### Response
```
{
    "note": {
        "ID": 10,
        "MediaID": "media_id",
        "UserID": "user_id",
        "Timestamp": 0,
        "Text": "hello world"
    }
}
```

Returns the following status codes: 
| Status Code | Description             |
| :---------- | :---------------------- |
| 200         | `CREATED`               |
| 400         | `BAD REQUEST`           | 
| 500         | `INTERNAL SERVER ERROR` |

### `GET /notes/media/{mediaID}`
This endpoint will query notes by media ID. 

#### Response
```
{
    "notes": [
        {
            "ID": 1,
            "MediaID": "spotify:track:4t4UX8LzCJ0K1yDlCjteTM",
            "UserID": "David",
            "Timestamp": 340,
            "Text": "The final build up to the end! "
        },
        {
            "ID": 11,
            "MediaID": "spotify:track:4t4UX8LzCJ0K1yDlCjteTM",
            "UserID": "David",
            "Timestamp": 375,
            "Text": "Here comes the oboe :D"
        }
    ]
}
```

Returns the following status codes: 
| Status Code | Description             |
| :---------- | :---------------------- |
| 200         | `OK`                    |
| 500         | `INTERNAL SERVER ERROR` |

### `GET /notes/user/{userID}`
This endpoint will query notes by user ID. Response is the same as querying by media ID. 

Returns the following status codes: 
| Status Code | Description             |
| :---------- | :---------------------- |
| 200         | `OK`                    |
| 500         | `INTERNAL SERVER ERROR` |

### `PUT /notes`
This endpoint will update notes and return the updated version. 

#### Request Body
```
{
    "ID": 1,
    "MediaID": "spotify:track:6fa4A3spdg6hQIYdHfX0vz",
    "UserID": "David",
    "Timestamp": 100,
    "Text": "This piece is really good!"
}
```

#### Response
```
{
    "note": {
        "ID": 1,
        "MediaID": "spotify:track:6fa4A3spdg6hQIYdHfX0vz",
        "UserID": "David",
        "Timestamp": 100,
        "Text": "This piece is really good!"
    }
}
```

Returns the following status codes: 
| Status Code | Description             |
| :---------- | :---------------------- |
| 200         | `UPDATED`               |
| 400         | `BAD REQUEST`           | 
| 500         | `INTERNAL SERVER ERROR` |

### `DELETE /notes/{noteID}`
This endpoint will delete notes by ID. 

Returns the following status codes: 
| Status Code | Description             |
| :---------- | :---------------------- |
| 204         | `DELETED`               |
| 500         | `INTERNAL SERVER ERROR` |