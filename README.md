# Mongo X

## Table of Content

- [Features](#features)
	- [Collection Methods](#collection-methods)
	- [Docker](#docker)
	- [Query Operators](#query-operators)
	- [Update Operators](#update-operators)
- [Examples](#examples)
	- [Collection](#collection)
	- [Filter](#query-operator)
- [License](#license)

## Features

### Collection Methods

|    | Method                       | Function                |
|----|------------------------------|-------------------------|
| ✔️ | `db.collection.count()`      | `collection.Count()`    |
| ✔️ | `db.collection.delete()`     | `collection.Delete()`   |
| ✔️ | `db.collection.find()`       | `collection.Find()`     |
| ✔️ | `db.collection.findMany()`   | `collection.FindMany()` |
| ✔️ | `db.collection.insert()`     | `collection.Insert()`   |
| ✔️ | `db.collection.replaceOne()` | `collection.Replace()`  |

### Docker

Mongo in Docker is useful for testing purposes. Please, check `*_test.go` files in the `/collection` directory for
examples.

```go
package main

import (
	"context"

	"entrlcom.dev/mongox/docker"
)

func main() {
	ctx := context.Background()

	mongo, err := docker.NewMongo(ctx)
	if err != nil {
		// TODO: Handle error.
		return
	}

	defer func() {
		// TODO: Handle error.
		_ = mongo.Shutdown(ctx) //nolint:errcheck // OK.
	}()

	client := mongo.GetClient()

	// ...
}

```

### Query Operators

#### Array Query Operators

|    | Operator     | Function            |
|----|--------------|---------------------|
| ❌  | `$all`       |                     |
| ✔️ | `$elemMatch` | `query.ElemMatch()` |
| ❌  | `$size`      |                     |

#### Bitwise Query Operators

|   | Operator        | Function |
|---|-----------------|----------|
| ❌ | `$bitsAllClear` |          |
| ❌ | `$bitsAllSet`   |          |
| ❌ | `$bitsAnyClear` |          |
| ❌ | `$bitsAnySet`   |          |

#### Comparison Query Operators

|    | Operator | Function      |
|----|----------|---------------|
| ✔️ | `$eq`    | `query.Eq()`  |
| ✔️ | `$gt`    | `query.Gt()`  |
| ✔️ | `$gte`   | `query.Gte()` |
| ✔️ | `$in`    | `query.In()`  |
| ✔️ | `$lt`    | `query.Lt()`  |
| ✔️ | `$lte`   | `query.Lte()` |
| ✔️ | `$ne`    | `query.Ne()`  |
| ✔️ | `$nin`   | `query.Nin()` |

#### Element Query Operators

|   | Operator  | Function |
|---|-----------|----------|
| ❌ | `$exists` |          |
| ❌ | `$type`   |          |

#### Evaluation Query Operators

|   | Operator      | Function |
|---|---------------|----------|
| ❌ | `$expr`       |          |
| ❌ | `$jsonSchema` |          |
| ❌ | `$mod`        |          |
| ❌ | `$regex`      |          |
| ❌ | `$text`       |          |
| ❌ | `$where`      |          |

#### Geospatial Query Operators

|   | Operator        | Function |
|---|-----------------|----------|
| ❌ | `box`           |          |
| ❌ | `center`        |          |
| ❌ | `centerSphere`  |          |
| ❌ | `geoIntersects` |          |
| ❌ | `geoWithin`     |          |
| ❌ | `geometry`      |          |
| ❌ | `maxDistance`   |          |
| ❌ | `minDistance`   |          |
| ❌ | `near`          |          |
| ❌ | `nearSphere`    |          |
| ❌ | `polygon`       |          |

#### Logical Query Operators

|    | Operator | Function      |
|----|----------|---------------|
| ✔️ | `$and`   | `query.And()` |
| ✔️ | `$or`    | `query.Or()`  |
| ✔️ | `$nor`   | `query.Nor()` |
| ✔️ | `$not`   | `query.Not()` |

### Update Operators

#### Array Update Operators

|   | Operator    | Function |
|---|-------------|----------|
| ❌ | `$addToSet` |          |
| ❌ | `$each`     |          |
| ❌ | `$pop`      |          |
| ❌ | `$position` |          |
| ❌ | `$pull`     |          |
| ❌ | `$pullAll`  |          |
| ❌ | `$push`     |          |
| ❌ | `$slice`    |          |
| ❌ | `$sort`     |          |

#### Bitwise Update Operators

|   | Operator | Function |
|---|----------|----------|
| ❌ | `$bit`   |          |

#### Field Update Operators

|   | Operator       | Function |
|---|----------------|----------|
| ❌ | `$currentDate` |          |
| ❌ | `$inc`         |          |
| ❌ | `$max`         |          |
| ❌ | `$min`         |          |
| ❌ | `$mul`         |          |
| ❌ | `$rename`      |          |
| ❌ | `$set`         |          |
| ❌ | `$setOnInsert` |          |
| ❌ | `$unset`       |          |

## Examples

### Collection

```go
package main

import (
	"context"
	"time"

	"entrlcom.dev/mongox/collection"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	DataOfBirth time.Time `bson:"data_of_birth,omitempty"`
	Name        string    `bson:"name,omitempty"`
	TimeCreated time.Time `bson:"time_created,omitempty"`
}

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (x Repository) Create(ctx context.Context, person Person) error {
	return collection.Insert[Person](x.collection).Insert(ctx, person)
}

func NewRepository(client *mongo.Client) Repository {
	return Repository{
		client:     client,
		collection: client.Database("example").Collection("person"),
	}
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		// TODO: Handle error.
		return
	}

	repository := NewRepository(client)

	if err := repository.Create(ctx, Person{
		DataOfBirth: time.Date(2020, time.May, 15, 0, 0, 0, 0, time.UTC),
		Name:        "John Doe",
		TimeCreated: time.Now().UTC(),
	}); err != nil {
		// TODO: Handle error.
		return
	}

	// ...
}

```

### Query Operator

```go
package main

import (
	"entrlcom.dev/mongox/operator/query"
)

func main() {
	filter := query.And(
		query.Or(
			query.Eq().Field("name").String("Jane"),
			query.Eq().Field("name").String("John"),
		),
		query.Eq().Field("id").String("1"),
	)

	// {
	//   "$and": [
	//     {
	//       "$or": [
	//         {
	//           "name": {
	//             "$eq": "Jane"
	//           }
	//         },
	//         {
	//           "name": {
	//             "$eq": "John"
	//           }
	//         }
	//       ]
	//     },
	//     {
	//       "id": {
	//         "$eq": "1"
	//       }
	//     }
	//   ]
	// }
}

```

## License

[MIT](https://choosealicense.com/licenses/mit/)
