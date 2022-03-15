# RESTful API with Go and Gin

Follow along with [this](https://go.dev/doc/tutorial/web-service-gin) go.dev tutorial, a simple web service with 3 endpoints.

An overview of all the tutorials can be found [here](../../..).

## API

### Get all available albums

* **Endpoint:** `/albums`

* **Method:** `GET`

* **URL Params**

  **Required:** -

  **Optional:** -

* **Data Params:** -

* **Headers** - 

* **Success Response:**

  **Code:** 200 <br />
  **Content:**

    ```
      [
        {
          "id": "1",
          "title": "Title 1",
          "artist": "Artist 1",
          "price": 10.12
        },
        {
          "id": "2",
          "title": "Title 2",
          "artist": "Artist 2",
          "price": 14.48
        },
        {
          "id": "3",
          "title": "Title 3",
          "artist": "Artist 1",
          "price": 9.56
        },
        {
          "id": "4",
          "title": "The Modern Sound of Betty Carter",
          "artist": "Betty Carter",
          "price": 49.99
        }
      ]
    ```

* **Error Response:**

  No defined errors at the moment.

### Get album by id

* **Endpoint:** `/albums/:id`

* **Method:** `GET`

* **URL Params**

  **Required:** `id=[string]`

  **Optional:** -

* **Data Params:** -

* **Headers:** -

* **Success Response:**

    **Code:** 200 <br />
    **Content:**

      ```
        {
          "id": "1",
          "title": "Title 1",
          "artist": "Artist 1",
          "price": 10.12
        }
      ```

* **Error Response:**

  **Code:** 404 <br />
  **Content:**

    ```
      {
        "message": "Album not found."
      }
    ```

### Add a new album (no validation added)

* **Endpoint:** `/albums`

* **Method:** `POST`

* **URL Params** -

  **Required:** -

  **Optional:** -

* **Data Params** -

  ```
    {
      "id": "4",
      "title": "Title 4",
      "artist": "Artist 4",
      "price": 10.12
    }
  ```
  
* **Headers:** `Content-Type: application/json`

* **Call:**

  ```
  curl http://localhost:8080/albums \
    --include \ 
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter", "artist": "Betty Carter","price": 49.99}'
  ```

* **Success Response:**

  **Code:** 200 <br />
  **Content:**

    ```
      {
        "id": "1",
        "title": "Title 1",
        "artist": "Artist 1",
        "price": 10.12
      }
    ```

* **Error Response:**

  No defined errors at the moment.